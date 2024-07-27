package repository

import (
	"rentbook/internal/features/book"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) book.RepositoryInterface {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(data *book.Books) (*book.Books, error) {
	err := r.db.Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *repository) FindByBookNameAndUserId(bookName, userId string) (*book.Books, error) {
	var data book.Books

	err := r.db.Where("user_id = ? AND book_name = ? AND is_delete = ?", userId, bookName, false).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) FindByBookIdAndUserId(bookId, userId string) (*book.Books, error) {
	var data book.Books

	err := r.db.Where("user_id = ? AND book_id = ? AND is_delete = ?", userId, bookId, false).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) GetByIdBook(bookId string) (*book.Books, error) {
	var data book.Books
	err := r.db.Where("book_id = ? AND is_delete = ?", bookId, false).First(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) GetAllByIdUser(userId string, queryParam map[string]string) (*[]book.BooksGetAllByIdUserResponse, error) {
	var data []book.BooksGetAllByIdUserResponse
	err := r.db.Table("books").
		Select("books.*,users.user_name").
		Joins("JOIN users ON users.user_id = books.user_id").
		Where("books.user_id = ? AND books.is_delete = ? AND books.book_name LIKE ? AND users.user_name LIKE ? AND books.book_publisher LIKE ? AND books.book_author LIKE ?", userId, false, queryParam["book_name"], queryParam["user_name"], queryParam["book_publisher"], queryParam["book_author"]).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *repository) GetCountByIdUser(userId string, queryParam map[string]string) (*int, error) {
	var data book.BooksCountByIdUserResponse
	err := r.db.Table("books").
		Select("count(books.book_id) as count").
		Joins("JOIN users ON users.user_id = books.user_id").
		Where("books.user_id = ? AND books.is_delete = ? AND books.book_name LIKE ? AND users.user_name LIKE ? AND books.book_publisher LIKE ? AND books.book_author LIKE ?", userId, false, queryParam["book_name"], queryParam["user_name"], queryParam["book_publisher"], queryParam["book_author"]).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data.Count, nil
}

func (r *repository) GetAll(queryParam map[string]string) (*[]book.BooksGetAllResponse, error) {
	var data []book.BooksGetAllResponse
	err := r.db.Table("books").
		Select("books.*").
		Joins("JOIN users ON users.user_id = books.user_id").
		Where("books.is_delete = ? AND books.book_name LIKE ? AND users.user_name LIKE ? AND books.book_publisher LIKE ? AND books.book_author LIKE ?", false, queryParam["book_name"], queryParam["user_name"], queryParam["book_publisher"], queryParam["book_author"]).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *repository) GetCount(queryParam map[string]string) (*int, error) {
	var data book.BooksCountResponse
	err := r.db.Table("books").
		Select("count(books.book_id) as count").
		Joins("JOIN users ON users.user_id = books.user_id").
		Where("books.is_delete = ? AND books.book_name LIKE ? AND users.user_name LIKE ? AND books.book_publisher LIKE ? AND books.book_author LIKE ?", false, queryParam["book_name"], queryParam["user_name"], queryParam["book_publisher"], queryParam["book_author"]).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data.Count, nil
}

func (r *repository) Update(inputData *book.Books, bookId string) (*book.Books, error) {
	var data book.Books
	err := r.db.Model(&data).Where("book_id = ?", bookId).Updates(inputData).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) Delete(bookId string) (*book.Books, error) {
	var data book.Books
	var inputData book.Books
	inputData.IsDelete = true
	err := r.db.Model(&data).Where("book_id = ?", bookId).Updates(inputData).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}
