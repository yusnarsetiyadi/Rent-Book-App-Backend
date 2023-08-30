package repository

import (
	"rentbook/internal/features/rent"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) rent.RepositoryInterface {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(data *rent.Rents) (*rent.Rents, error) {
	err := r.db.Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *repository) FindByRentIdBookIdUserId(rentId, bookId, userId string) (*rent.Rents, error) {
	var data rent.Rents

	err := r.db.Where("rent_id = ? AND user_id = ? AND book_id = ?", rentId, userId, bookId).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) FindByRentIdUserId(rentId, userId string) (*rent.Rents, error) {
	var data rent.Rents

	err := r.db.Where("rent_id = ? AND user_id = ?", rentId, userId).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) GetById(rentId string) (*rent.GetByIdResponse, error) {
	var data rent.GetByIdResponse
	err := r.db.Table("rents").Select("rents.*").Where("rents.rent_id = ?", rentId).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) GetAllByIdUser(userId string, queryParam map[string]string) (*[]rent.RentsGetAllByIdUserResponse, error) {
	var data []rent.RentsGetAllByIdUserResponse
	err := r.db.Table("rents").
		Select("rents.*,users.user_name,books.book_name").
		Joins("JOIN users ON users.user_id = rents.user_id").
		Joins("JOIN books ON books.book_id = rents.book_id").
		Where("rents.user_id = ? AND users.user_name LIKE ? AND books.book_name LIKE ? AND rents.rent_status LIKE ?", userId, queryParam["user_name"], queryParam["book_name"], queryParam["rent_status"]).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *repository) GetCountByIdUser(userId string, queryParam map[string]string) (*int, error) {
	var data rent.RentsCountByIdUserResponse
	err := r.db.Table("rents").
		Select("count(rents.rent_id) as count").
		Joins("JOIN users ON users.user_id = rents.user_id").
		Joins("JOIN books ON books.book_id = rents.book_id").
		Where("rents.user_id = ? AND users.user_name LIKE ? AND books.book_name LIKE ? AND rents.rent_status LIKE ?", userId, queryParam["user_name"], queryParam["book_name"], queryParam["rent_status"]).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data.Count, nil
}

func (r *repository) GetAllByIdBook(bookId string, queryParam map[string]string) (*[]rent.RentsGetAllByIdBookResponse, error) {
	var data []rent.RentsGetAllByIdBookResponse
	err := r.db.Table("rents").
		Select("rents.*,users.user_name,books.book_name").
		Joins("JOIN users ON users.user_id = rents.user_id").
		Joins("JOIN books ON books.book_id = rents.book_id").
		Where("rents.book_id = ? AND users.user_name LIKE ? AND books.book_name LIKE ? AND rents.rent_status LIKE ?", bookId, queryParam["user_name"], queryParam["book_name"], queryParam["rent_status"]).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *repository) GetCountByIdBook(bookId string, queryParam map[string]string) (*int, error) {
	var data rent.RentsCountByIdBookResponse
	err := r.db.Table("rents").
		Select("count(rents.rent_id) as count").
		Joins("JOIN users ON users.user_id = rents.user_id").
		Joins("JOIN books ON books.book_id = rents.book_id").
		Where("rents.book_id = ? AND users.user_name LIKE ? AND books.book_name LIKE ? AND rents.rent_status LIKE ?", bookId, queryParam["user_name"], queryParam["book_name"], queryParam["rent_status"]).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data.Count, nil
}

func (r *repository) Update(inputData *rent.Rents, rentId string) (*rent.Rents, error) {
	var data rent.Rents
	err := r.db.Model(&data).Where("rent_id = ?", rentId).Updates(inputData).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) Delete(rentId string) (*rent.Rents, error) {
	var data rent.Rents
	err := r.db.Where("rent_id = ?", rentId).Delete(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}
