package delivery

import (
	"net/http"
	"rentbook/internal/features/book"
	"rentbook/internal/middleware"
	"rentbook/utils/helper"

	"github.com/labstack/echo/v4"
)

type delivery struct {
	service book.ServiceInterface
}

func New(service book.ServiceInterface, e *echo.Echo) {
	handler := &delivery{
		service: service,
	}
	g := e.Group("/books")
	g.POST("", handler.Create, middleware.JWTMiddleware())
	g.GET("/book_id/:id", handler.GetByIdBook, middleware.JWTMiddleware())
	g.GET("/user_id/:id", handler.GetAllByIdUser, middleware.JWTMiddleware())
	g.GET("", handler.GetAll, middleware.JWTMiddleware())
	g.PUT("/:id", handler.Update, middleware.JWTMiddleware())
	g.PATCH("/:id", handler.Delete, middleware.JWTMiddleware())
}

func (d *delivery) Create(c echo.Context) error {
	var inputData = new(book.CreateRequest)
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

func (d *delivery) GetByIdBook(c echo.Context) error {
	var inputData = new(book.GetByIdBookRequest)
	if err := c.Bind(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Bind inputData", http.StatusBadRequest, err.Error()))
	}
	if err := c.Validate(inputData); err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Validate inputData", http.StatusBadRequest, err.Error()))
	}
	results, msg, code, err := d.service.GetByIdBook(inputData, c)
	if err != nil {
		return c.JSON(code, helper.FailedResponse("Error Handler Service / message="+msg, code, err.Error()))
	}

	return c.JSON(code, helper.SuccessWithDataResponse("Success GetByIdBook / message="+msg, code, results))
}

func (d *delivery) GetAllByIdUser(c echo.Context) error {
	var inputData = new(book.GetAllByIdUserRequest)
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

func (d *delivery) GetAll(c echo.Context) error {
	results, msg, code, err := d.service.GetAll(c)
	if err != nil {
		return c.JSON(code, helper.FailedResponse("Error Handler Service / message="+msg, code, err.Error()))
	}

	return c.JSON(code, helper.SuccessWithDataResponse("Success GetAll & GetCount / message="+msg, code, results))
}

func (d *delivery) Update(c echo.Context) error {
	var inputData = new(book.UpdateRequest)
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
	var inputData = new(book.DeleteRequest)
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
