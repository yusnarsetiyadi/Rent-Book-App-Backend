package middleware

import (
	"fmt"
	"net/http"
	"os"
	"rentbook/internal/config"
	"rentbook/utils/pkg/validator"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	APP := config.GetConfig().APP
	e.Use(
		echoMiddleware.Recover(),
		echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
		}),
		echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
			Format:           fmt.Sprintf("\n| %s | Host: ${host} | Time: ${time_custom} | Status: ${status} | LatencyHuman: ${latency_human} | UserAgent: ${user_agent} | RemoteIp: ${remote_ip} | Method: ${method} | Uri: ${uri} |\n", APP),
			CustomTimeFormat: "2006/01/02 15:04:05",
			Output:           os.Stdout,
		}),
		echoMiddleware.SecureWithConfig(echoMiddleware.SecureConfig{
			XFrameOptions: "DENY",
		}),
	)
	e.Validator = &validator.CustomValidator{Validator: validator.NewValidator()}
}
