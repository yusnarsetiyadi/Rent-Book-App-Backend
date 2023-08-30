package delivery

import (
	"net/http"
	"rentbook/internal/features/auth"
	"rentbook/internal/middleware"
	"rentbook/utils/helper"

	"github.com/labstack/echo/v4"
)

type delivery struct {
	service auth.ServiceInterface
}

func New(service auth.ServiceInterface, e *echo.Echo) {
	handler := &delivery{
		service: service,
	}
	g := e.Group("/auth")
	g.POST("/login", handler.Login)
	g.POST("/logout", handler.Logout, middleware.JWTMiddleware())
}

func (d *delivery) Login(c echo.Context) error {
	var inputData = new(auth.LoginRequest)
	if err := c.Bind(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Bind inputData", http.StatusBadRequest, err.Error()))
	}
	if err := c.Validate(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Validate inputData", http.StatusBadRequest, err.Error()))
	}
	results, msg, code, err := d.service.Login(*inputData, c)
	if err != nil {
		return c.JSON(code, helper.FailedResponse("Error Handler Service / message="+msg, code, err.Error()))
	}

	return c.JSON(code, helper.SuccessWithDataResponse("Success Login / message="+msg, code, results))
}

func (d *delivery) Logout(c echo.Context) error {
	results, msg, code, err := d.service.Logout(c)
	if err != nil {
		return c.JSON(code, helper.FailedResponse("Error Handler Service / message="+msg, code, err.Error()))
	}

	return c.JSON(code, helper.SuccessWithDataResponse("Success Logout / message="+msg, code, results))
}
