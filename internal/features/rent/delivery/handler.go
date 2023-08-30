package delivery

import (
	"net/http"
	"rentbook/internal/features/rent"
	"rentbook/internal/middleware"
	"rentbook/utils/helper"

	"github.com/labstack/echo/v4"
)

type delivery struct {
	service rent.ServiceInterface
}

func New(service rent.ServiceInterface, e *echo.Echo) {
	handler := &delivery{
		service: service,
	}
	g := e.Group("/rents")
	g.POST("", handler.Create, middleware.JWTMiddleware())
	g.GET("/:id", handler.GetById, middleware.JWTMiddleware())
	g.GET("/user_id/:id", handler.GetAllByIdUser, middleware.JWTMiddleware())
	g.GET("/book_id/:id", handler.GetAllByIdBook, middleware.JWTMiddleware())
	g.PUT("/:id", handler.Update, middleware.JWTMiddleware())
	g.DELETE("/:id", handler.Delete, middleware.JWTMiddleware())
}

func (d *delivery) Create(c echo.Context) error {
	var inputData = new(rent.CreateRequest)
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
	var inputData = new(rent.GetByIdRequest)
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

func (d *delivery) GetAllByIdUser(c echo.Context) error {
	var inputData = new(rent.GetAllByIdUserRequest)
	if err := c.Bind(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Bind inputData", http.StatusBadRequest, err.Error()))
	}
	if err := c.Validate(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Validate inputData", http.StatusBadRequest, err.Error()))
	}
	results, msg, code, err := d.service.GetAllByIdUser(inputData, c)
	if err != nil {
		return c.JSON(code, helper.FailedResponse("Error Handler Service / message="+msg, code, err.Error()))
	}

	return c.JSON(code, helper.SuccessWithDataResponse("Success GetAllByIdUser & GetCountByIdUser / message="+msg, code, results))
}

func (d *delivery) GetAllByIdBook(c echo.Context) error {
	var inputData = new(rent.GetAllByIdBookRequest)
	if err := c.Bind(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Bind inputData", http.StatusBadRequest, err.Error()))
	}
	if err := c.Validate(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Validate inputData", http.StatusBadRequest, err.Error()))
	}
	results, msg, code, err := d.service.GetAllByIdBook(inputData, c)
	if err != nil {
		return c.JSON(code, helper.FailedResponse("Error Handler Service / message="+msg, code, err.Error()))
	}

	return c.JSON(code, helper.SuccessWithDataResponse("Success GetAllByIdUser & GetCountByIdUser / message="+msg, code, results))
}

func (d *delivery) Update(c echo.Context) error {
	var inputData = new(rent.UpdateRequest)
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
	var inputData = new(rent.DeleteRequest)
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
