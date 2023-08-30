package repository

import (
	"rentbook/internal/features/user"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.RepositoryInterface {
	return &repository{
		db: db,
	}
}

func (r *repository) FindByEmail(userEmail string) (*user.Users, error) {
	var data user.Users

	err := r.db.Where("user_email = ? AND is_delete = ?", userEmail, false).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) Create(data *user.Users) (*user.Users, error) {
	err := r.db.Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *repository) GetByIdOnly(userId string) (*user.Users, error) {
	var data user.Users
	err := r.db.Where("user_id = ? AND is_delete = ?", userId, false).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) GetById(userId string) (*user.GetByIdResponse, error) {
	var data user.GetByIdResponse
	err := r.db.Table("users").
		Select("users.user_id, users.user_name, users.user_email, users.is_delete, users.created_at, users.updated_at").Where("users.user_id = ? AND users.is_delete = ?", userId, false).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) GetAll(queryParam map[string]string) (*[]user.UsersGetAllResponse, error) {
	var data []user.UsersGetAllResponse
	err := r.db.Table("users").
		Select("users.user_id, users.user_name, users.user_email, users.is_delete, users.created_at, users.updated_at").
		Where("users.is_delete = ? AND users.user_name LIKE ?", false, queryParam["user_name"]).
		Order("users.user_name asc").
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *repository) GetCount(queryParam map[string]string) (*int, error) {
	var data user.UsersCountResponse
	err := r.db.Table("users").
		Select("count(users.user_id) as count").
		Where("users.is_delete = ? AND users.user_name LIKE ?", false, queryParam["user_name"]).
		Order("users.user_name asc").
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data.Count, nil
}

func (r *repository) Update(inputData *user.Users, userId string) (*user.Users, error) {
	var data user.Users
	err := r.db.Model(&data).Where("user_id = ?", userId).Updates(inputData).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) Delete(userId string) (*user.Users, error) {
	var data user.Users
	var inputData user.Users
	inputData.IsDelete = true
	err := r.db.Model(&data).Where("user_id = ?", userId).Updates(inputData).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) ChangePassword(inputData *user.Users, userId string) (*user.Users, error) {
	var data user.Users
	err := r.db.Model(&data).Where("user_id = ?", userId).Updates(inputData).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}
