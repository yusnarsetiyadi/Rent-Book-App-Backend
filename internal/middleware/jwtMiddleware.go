package middleware

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"rentbook/internal/config"
	"rentbook/internal/features/auth"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

var key string

func InitJWT(c *config.AppConfig) {
	key = c.JWT_SECRET
}

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(key),
		ErrorHandler: func(err error) error {
			if err == middleware.ErrJWTMissing {
				return echo.NewHTTPError(http.StatusBadRequest, "missing or malformed jwt")
			}
			if _, ok := err.(*jwt.ValidationError); ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired jwt")
			}
			return err
		},
	})
}

func GenerateAccessToken(user *auth.Auth) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"userId":     user.UserId,
		"userName":   user.UserName,
		"userEmail":  user.UserEmail,
		"isDelete":   user.IsDelete,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

func GenerateRefreshToken(token auth.Token) (auth.Token, error) {
	sha1 := sha1.New()
	io.WriteString(sha1, config.GetConfig().JWT_SECRET)

	salt := string(sha1.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		return token, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return token, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return token, err
	}

	token.RefreshToken = base64.URLEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(token.AccessToken), nil))

	return token, nil
}

func ExtractTokenMapClaim(e echo.Context, field string) any {
	tokenString := e.Request().Header.Get("Authorization")
	if tokenString == "" {
		logrus.Error("Unauthorized JWT, because token not found.")
	}
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		logrus.Error("Error parsing token: ", err)
		return ""
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		result := claims[field]
		return result
	}
	return ""
}

func ChangeTokenForLogout(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		return errors.New("unauthorized jwt, because token not found")
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, errParseToken := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if errParseToken != nil || !token.Valid {
		return errParseToken
	}

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Unix()
	token.Claims = claims
	token.Valid = false

	newTokenString, errSignedString := token.SignedString([]byte(key))
	if errSignedString != nil {
		return errSignedString
	}

	c.Response().Header().Set("Authorization", "Bearer "+newTokenString)

	return nil
}
