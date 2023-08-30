package book

import "github.com/labstack/echo/v4"

type ServiceInterface interface {
	Create(inputData *CreateRequest, e echo.Context) (*CreateResponse, string, int, error)
	GetByIdBook(inputData *GetByIdBookRequest, e echo.Context) (*GetByIdBookResponse, string, int, error)
	GetAllByIdUser(inputData *GetAllByIdUserRequest, e echo.Context) (*GetAllByIdUserResponse, string, int, error)
	GetAll(e echo.Context) (*GetAllResponse, string, int, error)
	Update(inputData *UpdateRequest, e echo.Context) (*UpdateResponse, string, int, error)
	Delete(inputData *DeleteRequest, e echo.Context) (*DeleteResponse, string, int, error)
}

type RepositoryInterface interface {
	Create(data *Books) (*Books, error)
	FindByBookNameAndUserId(bookName, userId string) (*Books, error)
	FindByBookIdAndUserId(bookId, userId string) (*Books, error)
	GetByIdBook(bookId string) (*Books, error)
	GetAllByIdUser(userId string, queryParam map[string]string) (*[]BooksGetAllByIdUserResponse, error)
	GetCountByIdUser(userId string, queryParam map[string]string) (*int, error)
	GetAll(queryParam map[string]string) (*[]BooksGetAllResponse, error)
	GetCount(queryParam map[string]string) (*int, error)
	Update(inputData *Books, bookId string) (*Books, error)
	Delete(bookId string) (*Books, error)
}
