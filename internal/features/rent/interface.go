package rent

import "github.com/labstack/echo/v4"

type ServiceInterface interface {
	Create(inputData *CreateRequest, e echo.Context) (*CreateResponse, string, int, error)
	GetById(inputData *GetByIdRequest, e echo.Context) (*GetByIdResponse, string, int, error)
	GetAllByIdUser(inputData *GetAllByIdUserRequest, e echo.Context) (*GetAllByIdUserResponse, string, int, error)
	GetAllByIdBook(inputData *GetAllByIdBookRequest, e echo.Context) (*GetAllByIdBookResponse, string, int, error)
	Update(inputData *UpdateRequest, e echo.Context) (*UpdateResponse, string, int, error)
	Delete(inputData *DeleteRequest, e echo.Context) (*DeleteResponse, string, int, error)
}

type RepositoryInterface interface {
	Create(data *Rents) (*Rents, error)
	FindByRentIdBookIdUserId(rentId, bookId, userId string) (*Rents, error)
	FindByRentIdUserId(rentId, userId string) (*Rents, error)
	GetById(rentId string) (*GetByIdResponse, error)
	GetAllByIdUser(userId string, queryParam map[string]string) (*[]RentsGetAllByIdUserResponse, error)
	GetCountByIdUser(userId string, queryParam map[string]string) (*int, error)
	GetAllByIdBook(bookId string, queryParam map[string]string) (*[]RentsGetAllByIdBookResponse, error)
	GetCountByIdBook(bookId string, queryParam map[string]string) (*int, error)
	Update(inputData *Rents, rentId string) (*Rents, error)
	Delete(rentId string) (*Rents, error)
}
