package user

import "github.com/labstack/echo/v4"

type ServiceInterface interface {
	Create(inputData *CreateRequest, e echo.Context) (*CreateResponse, string, int, error)
	GetById(inputData *GetByIdRequest, e echo.Context) (*GetByIdResponse, string, int, error)
	GetAll(e echo.Context) (*GetAllResponse, string, int, error)
	Update(inputData *UpdateRequest, e echo.Context) (*UpdateResponse, string, int, error)
	Delete(inputData *DeleteRequest, e echo.Context) (*DeleteResponse, string, int, error)
	ChangePassword(inputData *ChangePasswordRequest, e echo.Context) (*ChangePasswordResponse, string, int, error)
}

type RepositoryInterface interface {
	FindByEmail(userEmail string) (*Users, error)
	Create(data *Users) (*Users, error)
	GetById(userId string) (*GetByIdResponse, error)
	GetByIdOnly(userId string) (*Users, error)
	GetAll(queryParam map[string]string) (*[]UsersGetAllResponse, error)
	GetCount(queryParam map[string]string) (*int, error)
	Update(inputData *Users, userId string) (*Users, error)
	Delete(userId string) (*Users, error)
	ChangePassword(inputData *Users, userId string) (*Users, error)
}
