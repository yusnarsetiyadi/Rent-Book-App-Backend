package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"rentbook/internal/config"
	"rentbook/internal/features/rent"
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
	rentRepo rent.RepositoryInterface
	validate *validator.Validate
	redis    *redis.Client
}

func New(repoRent rent.RepositoryInterface, Redis *redis.Client) rent.ServiceInterface {
	return &service{
		rentRepo: repoRent,
		validate: vald.NewValidator(),
		redis:    Redis,
	}
}

func (s *service) Create(inputData *rent.CreateRequest, e echo.Context) (*rent.CreateResponse, string, int, error) {
	var result rent.CreateResponse
	var data = new(rent.Rents)

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	userId := middleware.ExtractTokenMapClaim(e, "userId")
	if userId == "" {
		return nil, "Error Claim user id from token, please login.", http.StatusForbidden, errors.New("error claim user id")
	}

	rentId := general.GenerateUUID()
	startDate := general.ParseStringToTime(*inputData.RentStartDate)
	endDate := general.ParseStringToTime(*inputData.RentEndDate)
	data.RentId = rentId
	data.BookId = *inputData.BookId
	data.UserId = userId.(string)
	data.RentQty = *inputData.RentQty
	data.RentStartDate = startDate
	data.RentEndDate = endDate
	if general.DateTodayLocal().Before(startDate) {
		data.RentStatus = config.RENT_PENDING
	} else if general.DateTodayLocal().After(startDate) && general.DateTodayLocal().Before(endDate) {
		data.RentStatus = config.RENT_ACTIVE
	} else if general.DateTodayLocal().After(endDate) {
		data.RentStatus = config.RENT_EXPIRED
	}
	data.CreatedAt = *general.DateTodayLocal()
	data.UpdatedAt = *general.DateTodayLocal()

	_, errCreateRent := s.rentRepo.Create(data)
	if errCreateRent != nil {
		return nil, "Error query CreateRent", http.StatusInternalServerError, errCreateRent
	}

	result = rent.CreateResponse{
		Message: "Success Create Rent!",
	}

	return &result, "Success Create Rent", http.StatusOK, nil
}

func (s *service) GetById(inputData *rent.GetByIdRequest, e echo.Context) (*rent.GetByIdResponse, string, int, error) {
	var result rent.GetByIdResponse

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	rentData, errGetById := s.rentRepo.GetById(inputData.Id)
	if errGetById != nil {
		return nil, "Error query GetById Rent", http.StatusInternalServerError, errGetById
	}

	result = rent.GetByIdResponse{
		RentId:        rentData.RentId,
		BookId:        rentData.BookId,
		UserId:        rentData.UserId,
		RentStartDate: rentData.RentStartDate,
		RentEndDate:   rentData.RentEndDate,
		RentStatus:    rentData.RentStatus,
		RentQty:       rentData.RentQty,
		CreatedAt:     rentData.CreatedAt,
		UpdatedAt:     rentData.UpdatedAt,
	}

	return &result, "Success GetById Rent", http.StatusOK, nil
}

func (s *service) GetAllByIdUser(inputData *rent.GetAllByIdUserRequest, e echo.Context) (*rent.GetAllByIdUserResponse, string, int, error) {
	var result *rent.GetAllByIdUserResponse
	var data *[]rent.RentsGetAllByIdUserResponse
	var count *int

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	queryParam := map[string]string{
		"book_name":   "%" + e.QueryParam("book_name") + "%",
		"user_name":   "%" + e.QueryParam("user_name") + "%",
		"rent_status": "%" + e.QueryParam("rent_status") + "%",
	}

	data, errGetAll := s.rentRepo.GetAllByIdUser(inputData.Id, queryParam)
	if errGetAll != nil {
		return nil, "Error query GetAllByIdUser", http.StatusInternalServerError, errGetAll
	}

	count, errGetCount := s.rentRepo.GetCountByIdUser(inputData.Id, queryParam)
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
	dataRedisSave := fmt.Sprintf("Query Get All By Id User & Get Count By Id User with data= %s \ncount= %d \nparam user_name= %s \nparam book_name= %s \nparam rent_status= %s ", string(dataResult), *count, queryParam["user_name"], queryParam["book_name"], queryParam["rent_status"])
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

	result = &rent.GetAllByIdUserResponse{
		Data:  *data,
		Count: count,
	}

	return result, "Success GetAllById & GetCountById Rent", http.StatusOK, nil
}

func (s *service) GetAllByIdBook(inputData *rent.GetAllByIdBookRequest, e echo.Context) (*rent.GetAllByIdBookResponse, string, int, error) {
	var result *rent.GetAllByIdBookResponse
	var data *[]rent.RentsGetAllByIdBookResponse
	var count *int

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	queryParam := map[string]string{
		"book_name":   "%" + e.QueryParam("book_name") + "%",
		"user_name":   "%" + e.QueryParam("user_name") + "%",
		"rent_status": "%" + e.QueryParam("rent_status") + "%",
	}

	data, errGetAll := s.rentRepo.GetAllByIdBook(inputData.Id, queryParam)
	if errGetAll != nil {
		return nil, "Error query GetAllByIdBook", http.StatusInternalServerError, errGetAll
	}

	count, errGetCount := s.rentRepo.GetCountByIdBook(inputData.Id, queryParam)
	if errGetCount != nil {
		return nil, "Error query GetCountByIdBook", http.StatusInternalServerError, errGetCount
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
	dataRedisSave := fmt.Sprintf("Query Get All By Id Book & Get Count By Id Book with data= %s \ncount= %d \nparam user_name= %s \nparam book_name= %s \nparam rent_status= %s ", string(dataResult), *count, queryParam["user_name"], queryParam["book_name"], queryParam["rent_status"])
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

	result = &rent.GetAllByIdBookResponse{
		Data:  *data,
		Count: count,
	}

	return result, "Success GetAllById & GetCountById Rent", http.StatusOK, nil
}

func (s *service) Update(inputData *rent.UpdateRequest, e echo.Context) (*rent.UpdateResponse, string, int, error) {
	var result rent.UpdateResponse
	var data = new(rent.Rents)

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	userId := middleware.ExtractTokenMapClaim(e, "userId")
	if userId == "" {
		return nil, "Error Claim user id from token, please login.", http.StatusForbidden, errors.New("error claim user id")
	}

	rentData, errGetRentData := s.rentRepo.GetById(inputData.Id)
	if errGetRentData != nil {
		return nil, "Error query GetById", http.StatusInternalServerError, errGetRentData
	}

	_, errGetDataBookUserRent := s.rentRepo.FindByRentIdBookIdUserId(inputData.Id, rentData.BookId, userId.(string))
	if errGetDataBookUserRent != nil && errGetDataBookUserRent.Error() != "record not found" {
		return nil, "Error query FinfByBookIdAndUserIdAndRentId", http.StatusInternalServerError, errGetDataBookUserRent
	} else if errGetDataBookUserRent != nil && errGetDataBookUserRent.Error() == "record not found" {
		return nil, "Error this rent not allowed update by user", http.StatusNotAcceptable, errGetDataBookUserRent
	}

	startDate := general.ParseStringToTime(*inputData.RentStartDate)
	endDate := general.ParseStringToTime(*inputData.RentEndDate)
	data.RentStartDate = startDate
	data.RentEndDate = endDate
	data.RentQty = *inputData.RentQty
	if general.DateTodayLocal().Before(startDate) {
		data.RentStatus = config.RENT_PENDING
	} else if general.DateTodayLocal().After(startDate) && general.DateTodayLocal().Before(endDate) {
		data.RentStatus = config.RENT_ACTIVE
	} else if general.DateTodayLocal().After(endDate) {
		data.RentStatus = config.RENT_EXPIRED
	}
	data.UpdatedAt = *general.DateTodayLocal()

	_, errUpdateRent := s.rentRepo.Update(data, inputData.Id)
	if errUpdateRent != nil {
		return nil, "Error query UpdateRent", http.StatusInternalServerError, errUpdateRent
	}

	result = rent.UpdateResponse{
		Message: "Success Update Rent!",
	}

	return &result, "Success Update Rent", http.StatusOK, nil
}

func (s *service) Delete(inputData *rent.DeleteRequest, e echo.Context) (*rent.DeleteResponse, string, int, error) {
	var result rent.DeleteResponse

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	userId := middleware.ExtractTokenMapClaim(e, "userId")
	if userId == "" {
		return nil, "Error Claim user id from token, please login.", http.StatusForbidden, errors.New("error claim user id")
	}

	_, errGetDataRentUser := s.rentRepo.FindByRentIdUserId(inputData.Id, userId.(string))
	if errGetDataRentUser != nil && errGetDataRentUser.Error() != "record not found" {
		return nil, "Error query FinfByRentIdAndUserId", http.StatusInternalServerError, errGetDataRentUser
	} else if errGetDataRentUser != nil && errGetDataRentUser.Error() == "record not found" {
		return nil, "Error this rent not allowed delete by user", http.StatusNotAcceptable, errGetDataRentUser
	}

	_, errDeleteRent := s.rentRepo.Delete(inputData.Id)
	if errDeleteRent != nil {
		return nil, "Error query DeleteRent", http.StatusInternalServerError, errDeleteRent
	}

	result = rent.DeleteResponse{
		Message: "Success Delete Rent!",
	}
	return &result, "Success Delete Rent", http.StatusOK, nil
}
