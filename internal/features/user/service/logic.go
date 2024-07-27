package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"rentbook/internal/features/user"
	"rentbook/utils/pkg/general"
	vald "rentbook/utils/pkg/validator"
	"rentbook/utils/thirdparty/gomail"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepo user.RepositoryInterface
	validate *validator.Validate
	redis    *redis.Client
}

func New(repoUser user.RepositoryInterface, Redis *redis.Client) user.ServiceInterface {
	return &service{
		userRepo: repoUser,
		validate: vald.NewValidator(),
		redis:    Redis,
	}
}

func (s *service) Create(inputData *user.CreateRequest, e echo.Context) (*user.CreateResponse, string, int, error) {
	var result user.CreateResponse
	var data = new(user.Users)

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	userData, errFindByEmail := s.userRepo.FindByEmail(*inputData.UserEmail)
	if errFindByEmail != nil && errFindByEmail.Error() != "record not found" {
		return nil, "Error query FindByEmailUser.", http.StatusInternalServerError, errFindByEmail
	}

	if userData != nil {
		if userData.UserEmail == *inputData.UserEmail {
			return nil, "Error, email already exist.", http.StatusNotAcceptable, errors.New("email already exist")
		}
	}

	userId := general.GenerateUUID()
	passwordString := general.GeneratePassword(10, 1, 1, 1, 1)
	password := []byte(passwordString)
	hashedPassword, errGeneratePassword := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if errGeneratePassword != nil {
		return nil, "Error generate password", http.StatusInternalServerError, errGeneratePassword
	}

	data.UserId = userId
	data.UserName = *inputData.UserName
	data.UserEmail = *inputData.UserEmail
	data.UserPassword = string(hashedPassword)
	data.IsDelete = false
	data.CreatedAt = *general.DateTodayLocal()
	data.UpdatedAt = *general.DateTodayLocal()

	_, errCreateUser := s.userRepo.Create(data)
	if errCreateUser != nil {
		return nil, "Error query CreateUser", http.StatusInternalServerError, errCreateUser
	}

	msg, errSendEmailLoginInfo := gomail.SendEmailLoginInfo(data.UserEmail, "Rentbook - Login Info", data.UserEmail, passwordString, data.UserName)
	if errSendEmailLoginInfo != nil {
		return nil, msg, http.StatusInternalServerError, errSendEmailLoginInfo
	}

	result = user.CreateResponse{
		Message: "Success Create User!",
	}

	return &result, "Success Create User", http.StatusOK, nil
}

func (s *service) GetById(inputData *user.GetByIdRequest, e echo.Context) (*user.GetByIdResponse, string, int, error) {
	var result user.GetByIdResponse

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	userData, errGetByIdUser := s.userRepo.GetById(inputData.Id)
	if errGetByIdUser != nil {
		return nil, "Error query GetByIdUser", http.StatusInternalServerError, errGetByIdUser
	}

	result = user.GetByIdResponse{
		UserId:    userData.UserId,
		UserName:  userData.UserName,
		UserEmail: userData.UserEmail,
		IsDelete:  userData.IsDelete,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}

	return &result, "Success GetById User", http.StatusOK, nil
}

func (s *service) GetAll(e echo.Context) (*user.GetAllResponse, string, int, error) {
	var result *user.GetAllResponse
	var data *[]user.UsersGetAllResponse
	var count *int

	queryParam := map[string]string{
		"user_name": "%" + e.QueryParam("user_name") + "%",
	}

	data, errGetAll := s.userRepo.GetAll(queryParam)
	if errGetAll != nil {
		return nil, "Error query GetAll", http.StatusInternalServerError, errGetAll
	}

	count, errGetCount := s.userRepo.GetCount(queryParam)
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
	dataRedisSave := fmt.Sprintf("Query Get All & Get Count with data=%s \ncount= %d \nquery param=%s", string(dataResult), *count, queryParam["user_name"])
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

	result = &user.GetAllResponse{
		Data:  *data,
		Count: count,
	}

	return result, "Success GetAll & GetCount User", http.StatusOK, nil
}

func (s *service) Update(inputData *user.UpdateRequest, e echo.Context) (*user.UpdateResponse, string, int, error) {
	var result user.UpdateResponse
	var data = new(user.Users)

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	userData, errFindByEmail := s.userRepo.FindByEmail(*inputData.UserEmail)
	if errFindByEmail != nil && errFindByEmail.Error() != "record not found" {
		return nil, "Error query FindByEmailUser.", http.StatusInternalServerError, errFindByEmail
	}

	if userData != nil {
		if userData.UserEmail == *inputData.UserEmail {
			return nil, "Error, email already exist.", http.StatusNotAcceptable, errors.New("email already exist")
		}
	}

	data.UserName = *inputData.UserName
	data.UserEmail = *inputData.UserEmail
	data.UpdatedAt = *general.DateTodayLocal()

	_, errUpdateUser := s.userRepo.Update(data, inputData.Id)
	if errUpdateUser != nil {
		return nil, "Error query UpdateUser", http.StatusInternalServerError, errUpdateUser
	}

	result = user.UpdateResponse{
		Message: "Success Update User!",
	}

	return &result, "Success Update User", http.StatusOK, nil
}

func (s *service) Delete(inputData *user.DeleteRequest, e echo.Context) (*user.DeleteResponse, string, int, error) {
	var result user.DeleteResponse

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	userData, errGetUser := s.userRepo.GetById(inputData.Id)
	if errGetUser != nil {
		return nil, "Error query GetUser", http.StatusInternalServerError, errGetUser
	}

	_, errDeleteUser := s.userRepo.Delete(inputData.Id)
	if errDeleteUser != nil {
		return nil, "Error query DeleteUser", http.StatusInternalServerError, errDeleteUser
	}

	msg, errSendEmail := gomail.SendEmailDeactiveUser(userData.UserEmail, "ProExSys - Deactive User", userData.UserName)
	if errSendEmail != nil {
		return nil, msg, http.StatusInternalServerError, errSendEmail
	}

	result = user.DeleteResponse{
		Message: "Success Deactive User!",
	}
	return &result, "Success Deactive User", http.StatusOK, nil
}

func (s *service) ChangePassword(inputData *user.ChangePasswordRequest, e echo.Context) (*user.ChangePasswordResponse, string, int, error) {
	var result user.ChangePasswordResponse
	var data = new(user.Users)

	errValidate := s.validate.Struct(inputData)
	if errValidate != nil {
		return nil, "Error Validate input data, check required field", http.StatusBadRequest, errValidate
	}

	userData, errGetUser := s.userRepo.GetByIdOnly(inputData.Id)
	if errGetUser != nil {
		return nil, "Error query GetUser", http.StatusInternalServerError, errGetUser
	}

	errHashing := bcrypt.CompareHashAndPassword([]byte(userData.UserPassword), []byte(inputData.OldPassword))
	if errHashing != nil {
		return nil, "Error, your password is wrong!", http.StatusBadRequest, errHashing
	}

	password := []byte(inputData.NewPassword)
	hashedPassword, errHashedPassword := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if errHashedPassword != nil {
		return nil, "Error Hashed Password", http.StatusInternalServerError, errHashedPassword
	}

	data.UserPassword = string(hashedPassword)

	_, errChangePassword := s.userRepo.Update(data, inputData.Id)
	if errChangePassword != nil {
		return nil, "Error query UpdateUser", http.StatusInternalServerError, errChangePassword
	}

	result = user.ChangePasswordResponse{
		Message: "Success Change Password!",
	}

	return &result, "Success Change Password", http.StatusOK, nil
}
