package delivery

import (
	"net/http"
	"rentbook/internal/features/user"
	"rentbook/internal/middleware"
	"rentbook/utils/helper"

	"github.com/labstack/echo/v4"
)

type delivery struct {
	service user.ServiceInterface
}

func New(service user.ServiceInterface, e *echo.Echo) {
	handler := &delivery{
		service: service,
	}
	g := e.Group("/users")
	g.POST("", handler.Create)
	g.GET("/:id", handler.GetById, middleware.JWTMiddleware())
	g.GET("", handler.GetAll, middleware.JWTMiddleware())
	g.PUT("/:id", handler.Update, middleware.JWTMiddleware())
	g.PATCH("/:id", handler.Delete, middleware.JWTMiddleware())
	g.PATCH("/change_password/:id", handler.ChangePassword, middleware.JWTMiddleware())
}

func (d *delivery) Create(c echo.Context) error {
	var inputData = new(user.CreateRequest)
	if err := c.Bind(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Bind inputData", http.StatusBadRequest, err.Error()))
	}
	if err := c.Validate(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Validate inputData", http.StatusBadRequest, err.Error()))
	}
	results, msg, code, err := d.service.Create(inputData, c)
	if err != nil {
		return c.JSON(code, helper.FailedResponse("Error Handler Service / message="+msg, code, err.Error()))
	}

	return c.JSON(code, helper.SuccessWithDataResponse("Success Create / message="+msg, code, results))
}

func (d *delivery) GetById(c echo.Context) error {
	var inputData = new(user.GetByIdRequest)
	// inputData.Id = middleware.ExtractTokenClaimString(c, "userId") // extract token user id from user login
	if err := c.Bind(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Bind inputData", http.StatusBadRequest, err.Error()))
	}
	if err := c.Validate(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Validate inputData", http.StatusBadRequest, err.Error()))
	}
	results, msg, code, err := d.service.GetById(inputData, c)
	if err != nil {
		return c.JSON(code, helper.FailedResponse("Error Handler Service / message="+msg, code, err.Error()))
	}

	return c.JSON(code, helper.SuccessWithDataResponse("Success GetById / message="+msg, code, results))
}

func (d *delivery) GetAll(c echo.Context) error {
	results, msg, code, err := d.service.GetAll(c)
	if err != nil {
		return c.JSON(code, helper.FailedResponse("Error Handler Service / message="+msg, code, err.Error()))
	}

	return c.JSON(code, helper.SuccessWithDataResponse("Success GetAll & GetCount / message="+msg, code, results))
}

func (d *delivery) Update(c echo.Context) error {
	var inputData = new(user.UpdateRequest)
	// inputData.Id = middleware.ExtractTokenClaimString(c, "userId") // extract token user id from user login
	if err := c.Bind(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Bind inputData", http.StatusBadRequest, err.Error()))
	}
	if err := c.Validate(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Validate inputData", http.StatusBadRequest, err.Error()))
	}
	results, msg, code, err := d.service.Update(inputData, c)
	if err != nil {
		return c.JSON(code, helper.FailedResponse("Error Handler Service / message="+msg, code, err.Error()))
	}

	return c.JSON(code, helper.SuccessWithDataResponse("Success Update / message="+msg, code, results))
}

func (d *delivery) Delete(c echo.Context) error {
	var inputData = new(user.DeleteRequest)
	// inputData.Id = middleware.ExtractTokenClaimString(c, "userId") // extract token user id from user login
	if err := c.Bind(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Bind inputData", http.StatusBadRequest, err.Error()))
	}
	if err := c.Validate(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Validate inputData", http.StatusBadRequest, err.Error()))
	}
	results, msg, code, err := d.service.Delete(inputData, c)
	if err != nil {
		return c.JSON(code, helper.FailedResponse("Error Handler Service / message="+msg, code, err.Error()))
	}

	return c.JSON(code, helper.SuccessWithDataResponse("Success Delete / message="+msg, code, results))
}

func (d *delivery) ChangePassword(c echo.Context) error {
	var inputData = new(user.ChangePasswordRequest)
	// inputData.Id = middleware.ExtractTokenClaimString(c, "userId") // extract token user id from user login
	if err := c.Bind(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Bind inputData", http.StatusBadRequest, err.Error()))
	}
	if err := c.Validate(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Validate inputData", http.StatusBadRequest, err.Error()))
	}
	results, msg, code, err := d.service.ChangePassword(inputData, c)
	if err != nil {
		return c.JSON(code, helper.FailedResponse("Error Handler Service / message="+msg, code, err.Error()))
	}

	return c.JSON(code, helper.SuccessWithDataResponse("Success Change Password / message="+msg, code, results))
}
