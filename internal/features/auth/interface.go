package auth

import "github.com/labstack/echo/v4"

type ServiceInterface interface {
	Login(inputData LoginRequest, c echo.Context) (*LoginResponse, string, int, error)
	Logout(c echo.Context) (*LogoutResponse, string, int, error)
}
