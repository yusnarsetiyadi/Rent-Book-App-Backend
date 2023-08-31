package service

import (
	"errors"
	"fmt"
	"net/http"
	"rentbook/internal/config"
	"rentbook/internal/features/auth"
	"rentbook/internal/features/rent"
	"rentbook/internal/features/user"
	"rentbook/internal/middleware"
	"rentbook/utils/pkg/general"
	vald "rentbook/utils/pkg/validator"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepo user.RepositoryInterface
	rentRepo rent.RepositoryInterface
	validate *validator.Validate
	redis    *redis.Client
}

func New(repoUser user.RepositoryInterface, repoRent rent.RepositoryInterface, Redis *redis.Client) auth.ServiceInterface {
	return &service{
		userRepo: repoUser,
		validate: vald.NewValidator(),
		redis:    Redis,
		rentRepo: repoRent,
	}
}

func (s *service) Login(inputData auth.LoginRequest, c echo.Context) (*auth.LoginResponse, string, int, error) {
	var result auth.LoginResponse
	var dataUser = new(user.Users)

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	dataUser, errFindByEmail := s.userRepo.FindByEmail(*inputData.UserEmail)
	if errFindByEmail != nil && errFindByEmail.Error() != "record not found" {
		return nil, "Error query FindByEmailUser.", http.StatusInternalServerError, errFindByEmail
	} else if errFindByEmail != nil && errFindByEmail.Error() == "record not found" {
		return nil, "Error, your email entered didn't match or user deactivated", http.StatusBadRequest, errFindByEmail
	}

	if errCheckPassword := bcrypt.CompareHashAndPassword([]byte(dataUser.UserPassword), []byte(*inputData.UserPassword)); errCheckPassword != nil {
		return nil, "Error, your password entered didn't match", http.StatusBadRequest, errCheckPassword
	}

	var newToken = new(auth.Token)

	saveClaims := auth.Auth{
		UserId:    dataUser.UserId,
		UserName:  dataUser.UserName,
		UserEmail: dataUser.UserEmail,
		IsDelete:  dataUser.IsDelete,
	}
	accessToken, errAccessToken := middleware.GenerateAccessToken(&saveClaims)
	if errAccessToken != nil {
		return nil, "Error, failed generate access token.", http.StatusInternalServerError, errAccessToken
	}
	newToken.AccessToken = accessToken
	fullToken, errRefreshToken := middleware.GenerateRefreshToken(*newToken)
	if errRefreshToken != nil {
		return nil, "Error, failed generate refresh token.", http.StatusInternalServerError, errRefreshToken
	}

	timeCreated := general.DateTodayLocal().String()
	dataRedisSave := fmt.Sprintf("Event Created for Login user=%s(id=%s) at %s", dataUser.UserEmail, dataUser.UserId, timeCreated)
	_, errCreateEvent := s.redis.LPush(dataUser.UserId, dataRedisSave).Result()
	if errCreateEvent != nil {
		return nil, "Error save event data to redis", http.StatusInternalServerError, errCreateEvent
	}
	logrus.Info("Success Save Redis Event Data")

	dataEventRedis, errGetEventRedis := s.redis.LRange(dataUser.UserId, 0, -1).Result()
	if errGetEventRedis != nil {
		logrus.Error("Error Get Redis Event Data")
	} else {
		for _, event := range dataEventRedis {
			logrus.Print("Redis Event Data: " + event)
		}
	}

	queryParam := map[string]string{
		"book_name":   "%" + c.QueryParam("book_name") + "%",
		"user_name":   "%" + c.QueryParam("user_name") + "%",
		"rent_status": "%" + c.QueryParam("rent_status") + "%",
	}
	rowRentOnUser, errGetRentByIdUser := s.rentRepo.GetAllByIdUser(dataUser.UserId, queryParam)
	if errGetRentByIdUser != nil && errGetRentByIdUser.Error() != "record not found" {
		return nil, "Error get row rent data by user id", http.StatusInternalServerError, errGetRentByIdUser
	}
	if rowRentOnUser != nil {
		for _, rowRent := range *rowRentOnUser {
			var dataRent = new(rent.Rents)
			if general.DateTodayLocal().Before(rowRent.RentStartDate) {
				dataRent.RentStatus = config.RENT_PENDING
			} else if general.DateTodayLocal().After(rowRent.RentStartDate) && general.DateTodayLocal().Before(rowRent.RentEndDate) {
				dataRent.RentStatus = config.RENT_ACTIVE
			} else if general.DateTodayLocal().After(rowRent.RentEndDate) {
				dataRent.RentStatus = config.RENT_EXPIRED
			}
			dataRent.UpdatedAt = *general.DateTodayLocal()
			_, errUpdateRent := s.rentRepo.Update(dataRent, rowRent.RentId)
			if errUpdateRent != nil {
				return nil, "Error query UpdateRent", http.StatusInternalServerError, errUpdateRent
			}
		}
	}

	result = auth.LoginResponse{
		UserId:    dataUser.UserId,
		UserName:  dataUser.UserName,
		UserEmail: dataUser.UserEmail,
		IsDelete:  dataUser.IsDelete,
		Token:     fullToken,
	}

	return &result, "Success Login", http.StatusOK, nil
}

func (s *service) Logout(c echo.Context) (*auth.LogoutResponse, string, int, error) {
	var result auth.LogoutResponse

	errTokenLogout := middleware.ChangeTokenForLogout(c)
	if errTokenLogout != nil {
		return nil, "Error Change Logout Token.", http.StatusInternalServerError, errTokenLogout
	}

	keyRedis := middleware.ExtractTokenMapClaim(c, "userId")
	if keyRedis == "" {
		return nil, "Error Claim redis key (user id) from token, please login.", http.StatusForbidden, errors.New("error claim user id")
	}
	dataRedisEvent, errGetEventRedis := s.redis.LPop(keyRedis.(string)).Result()
	if errGetEventRedis == redis.Nil {
		logrus.Error("Data Redis Event not found. on user=" + keyRedis.(string))
	} else if errGetEventRedis != nil {
		logrus.Error("Error Get Redis Event Data not found")
	} else {
		logrus.Print("Redis Event Data Deleted: " + dataRedisEvent)
	}
	logrus.Info("Success Deleted Redis Event Data")

	result = auth.LogoutResponse{
		Message: "Success Logout and token deactivation",
	}

	return &result, "Success Logout", http.StatusOK, nil
}
