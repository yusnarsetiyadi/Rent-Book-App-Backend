package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"rentbook/internal/features/book"
	"rentbook/internal/middleware"
	"rentbook/utils/pkg/general"
	vald "rentbook/utils/pkg/validator"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type service struct {
	bookRepo book.RepositoryInterface
	validate *validator.Validate
	redis    *redis.Client
}

func New(repoBook book.RepositoryInterface, Redis *redis.Client) book.ServiceInterface {
	return &service{
		bookRepo: repoBook,
		validate: vald.NewValidator(),
		redis:    Redis,
	}
}

func (s *service) Create(inputData *book.CreateRequest, e echo.Context) (*book.CreateResponse, string, int, error) {
	var result book.CreateResponse
	var data = new(book.Books)

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	userId := middleware.ExtractTokenMapClaim(e, "userId")
	if userId == "" {
		return nil, "Error Claim user id from token, please login.", http.StatusForbidden, errors.New("error claim user id")
	}

	bookData, errFindByBookName := s.bookRepo.FindByBookNameAndUserId(*inputData.BookName, userId.(string))
	if errFindByBookName != nil && errFindByBookName.Error() != "record not found" {
		return nil, "Error query FindByBookName.", http.StatusInternalServerError, errFindByBookName
	}

	if bookData != nil {
		if bookData.BookName == *inputData.BookName {
			return nil, "Error, book name already used.", http.StatusNotAcceptable, errors.New("book name is already used")
		}
	}

	bookId := general.GenerateUUID()
	data.BookId = bookId
	data.BookName = *inputData.BookName
	data.BookAuthor = *inputData.BookAuthor
	data.BookPublisher = *inputData.BookPublisher
	data.UserId = userId.(string)
	data.IsDelete = false
	data.CreatedAt = *general.DateTodayLocal()
	data.UpdatedAt = *general.DateTodayLocal()

	_, errCreateBook := s.bookRepo.Create(data)
	if errCreateBook != nil {
		return nil, "Error query CreateBook", http.StatusInternalServerError, errCreateBook
	}

	result = book.CreateResponse{
		Message: "Success Create Book!",
	}

	return &result, "Success Create Book", http.StatusOK, nil
}

func (s *service) GetByIdBook(inputData *book.GetByIdBookRequest, e echo.Context) (*book.GetByIdBookResponse, string, int, error) {
	var result book.GetByIdBookResponse

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	bookData, errGetByIdBook := s.bookRepo.GetByIdBook(inputData.Id)
	if errGetByIdBook != nil {
		return nil, "Error query GetByIdBook", http.StatusInternalServerError, errGetByIdBook
	}

	result = book.GetByIdBookResponse{
		BookId:        bookData.BookId,
		UserId:        bookData.UserId,
		BookName:      bookData.BookName,
		BookPublisher: bookData.BookPublisher,
		BookAuthor:    bookData.BookAuthor,
		IsDelete:      bookData.IsDelete,
		CreatedAt:     bookData.CreatedAt,
		UpdatedAt:     bookData.UpdatedAt,
	}

	return &result, "Success GetByIdBook Book", http.StatusOK, nil
}

func (s *service) GetAllByIdUser(inputData *book.GetAllByIdUserRequest, e echo.Context) (*book.GetAllByIdUserResponse, string, int, error) {
	var result *book.GetAllByIdUserResponse
	var data *[]book.BooksGetAllByIdUserResponse
	var count *int

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	queryParam := map[string]string{
		"book_name":      "%" + e.QueryParam("book_name") + "%",
		"user_name":      "%" + e.QueryParam("user_name") + "%",
		"book_publisher": "%" + e.QueryParam("book_publisher") + "%",
		"book_author":    "%" + e.QueryParam("book_author") + "%",
	}

	data, errGetAll := s.bookRepo.GetAllByIdUser(inputData.Id, queryParam)
	if errGetAll != nil {
		return nil, "Error query GetAllByIdUser", http.StatusInternalServerError, errGetAll
	}

	count, errGetCount := s.bookRepo.GetCountByIdUser(inputData.Id, queryParam)
	if errGetCount != nil {
		return nil, "Error query GetCountByIdUser", http.StatusInternalServerError, errGetCount
	}

	keyRedis := "save_temporary_data"
	_, errDel := s.redis.Del(keyRedis).Result()
	if errDel != nil {
		logrus.Error("Error delete data redis")
	}
	logrus.Info("Success Delete Redis Temporary Data Previously")

	dataResult, errMarshalData := json.Marshal(data)
	if errMarshalData != nil {
		return nil, "Error marshal data for save temporary redis", http.StatusInternalServerError, errMarshalData
	}
	dataRedisSave := fmt.Sprintf("Query Get All By Id User & Get Count By Id User with data= %s \ncount= %d \nparam user_name= %s \nparam book_name= %s \nparam book_publisher= %s \nparam book_author= %s", string(dataResult), *count, queryParam["user_name"], queryParam["book_name"], queryParam["book_publisher"], queryParam["book_author"])
	errSaveDataRedis := s.redis.Set(keyRedis, dataRedisSave, 1*time.Hour).Err()
	if errSaveDataRedis != nil {
		return nil, "Error save temporary data to redis", http.StatusInternalServerError, errSaveDataRedis
	}
	logrus.Info("Success Save Redis Temporary Data")

	dataRedisTemporary, errGetDataRedis := s.redis.Get(keyRedis).Result()
	if errGetDataRedis == redis.Nil {
		logrus.Error("Data Redis Temporary Data not found.")
	} else if errGetDataRedis != nil {
		logrus.Error("Error Get Redis Temporary Data not found")
	} else {
		logrus.Print("Redis Temporary Data: " + dataRedisTemporary)
	}

	result = &book.GetAllByIdUserResponse{
		Data:  *data,
		Count: count,
	}

	return result, "Success GetAllById & GetCountById Book", http.StatusOK, nil
}

func (s *service) GetAll(e echo.Context) (*book.GetAllResponse, string, int, error) {
	var result *book.GetAllResponse
	var data *[]book.BooksGetAllResponse
	var count *int

	queryParam := map[string]string{
		"book_name":      "%" + e.QueryParam("book_name") + "%",
		"user_name":      "%" + e.QueryParam("user_name") + "%",
		"book_publisher": "%" + e.QueryParam("book_publisher") + "%",
		"book_author":    "%" + e.QueryParam("book_author") + "%",
	}

	data, errGetAll := s.bookRepo.GetAll(queryParam)
	if errGetAll != nil {
		return nil, "Error query GetAll", http.StatusInternalServerError, errGetAll
	}

	count, errGetCount := s.bookRepo.GetCount(queryParam)
	if errGetCount != nil {
		return nil, "Error query GetCount", http.StatusInternalServerError, errGetCount
	}

	keyRedis := "save_temporary_data"
	_, errDel := s.redis.Del(keyRedis).Result()
	if errDel != nil {
		logrus.Error("Error delete data redis")
	}
	logrus.Info("Success Delete Redis Temporary Data Previously")

	dataResult, errMarshalData := json.Marshal(data)
	if errMarshalData != nil {
		return nil, "Error marshal data for save temporary redis", http.StatusInternalServerError, errMarshalData
	}
	dataRedisSave := fmt.Sprintf("Query Get All & Get Count with with data= %s \ncount= %d \nparam user_name= %s \nparam book_name= %s \nparam book_publisher= %s \nparam book_author= %s", string(dataResult), *count, queryParam["user_name"], queryParam["book_name"], queryParam["book_publisher"], queryParam["book_author"])
	errSaveDataRedis := s.redis.Set(keyRedis, dataRedisSave, 1*time.Hour).Err()
	if errSaveDataRedis != nil {
		return nil, "Error save temporary data to redis", http.StatusInternalServerError, errSaveDataRedis
	}
	logrus.Info("Success Save Redis Temporary Data")

	dataRedisTemporary, errGetDataRedis := s.redis.Get(keyRedis).Result()
	if errGetDataRedis == redis.Nil {
		logrus.Error("Data Redis Temporary Data not found.")
	} else if errGetDataRedis != nil {
		logrus.Error("Error Get Redis Temporary Data not found")
	} else {
		logrus.Print("Redis Temporary Data: " + dataRedisTemporary)
	}

	result = &book.GetAllResponse{
		Data:  *data,
		Count: count,
	}

	return result, "Success GetAll & GetCount Book", http.StatusOK, nil
}

func (s *service) Update(inputData *book.UpdateRequest, e echo.Context) (*book.UpdateResponse, string, int, error) {
	var result book.UpdateResponse
	var data = new(book.Books)

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	userId := middleware.ExtractTokenMapClaim(e, "userId")
	if userId == "" {
		return nil, "Error Claim user id from token, please login.", http.StatusForbidden, errors.New("error claim user id")
	}

	bookData, errFindByBookName := s.bookRepo.FindByBookNameAndUserId(*inputData.BookName, userId.(string))
	if errFindByBookName != nil && errFindByBookName.Error() != "record not found" {
		return nil, "Error query FindByBookName.", http.StatusInternalServerError, errFindByBookName
	}

	if bookData != nil {
		if bookData.BookName == *inputData.BookName {
			return nil, "Error, book name already used.", http.StatusNotAcceptable, errors.New("book name is already used")
		}
	}

	_, errGetDataBookUser := s.bookRepo.FindByBookIdAndUserId(inputData.Id, userId.(string))
	if errGetDataBookUser != nil && errGetDataBookUser.Error() != "record not found" {
		return nil, "Error query FinfByBookIdAndUserId", http.StatusInternalServerError, errGetDataBookUser
	} else if errGetDataBookUser != nil && errGetDataBookUser.Error() == "record not found" {
		return nil, "Error this book not allowed update by user", http.StatusNotAcceptable, errGetDataBookUser
	}

	data.BookName = *inputData.BookName
	data.BookPublisher = *inputData.BookPublisher
	data.BookAuthor = *inputData.BookAuthor
	data.UpdatedAt = *general.DateTodayLocal()

	_, errUpdateBook := s.bookRepo.Update(data, inputData.Id)
	if errUpdateBook != nil {
		return nil, "Error query UpdateBook", http.StatusInternalServerError, errUpdateBook
	}

	result = book.UpdateResponse{
		Message: "Success Update Book!",
	}

	return &result, "Success Update Book", http.StatusOK, nil
}

func (s *service) Delete(inputData *book.DeleteRequest, e echo.Context) (*book.DeleteResponse, string, int, error) {
	var result book.DeleteResponse

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	userId := middleware.ExtractTokenMapClaim(e, "userId")
	if userId == "" {
		return nil, "Error Claim user id from token, please login.", http.StatusForbidden, errors.New("error claim user id")
	}

	_, errGetDataBookUser := s.bookRepo.FindByBookIdAndUserId(inputData.Id, userId.(string))
	if errGetDataBookUser != nil && errGetDataBookUser.Error() != "record not found" {
		return nil, "Error query FinfByBookIdAndUserId", http.StatusInternalServerError, errGetDataBookUser
	} else if errGetDataBookUser != nil && errGetDataBookUser.Error() == "record not found" {
		return nil, "Error this book not allowed delete by user", http.StatusNotAcceptable, errGetDataBookUser
	}

	_, errDeleteBook := s.bookRepo.Delete(inputData.Id)
	if errDeleteBook != nil {
		return nil, "Error query DeleteBook", http.StatusInternalServerError, errDeleteBook
	}

	result = book.DeleteResponse{
		Message: "Success Delete Book!",
	}
	return &result, "Success Delete Book", http.StatusOK, nil
}
