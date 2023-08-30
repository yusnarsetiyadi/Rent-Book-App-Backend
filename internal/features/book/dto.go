package book

import "time"

type CreateRequest struct {
	BookName      *string `json:"book_name" form:"book_name" validate:"required"`
	BookPublisher *string `json:"book_publisher" form:"book_publisher"`
	BookAuthor    *string `json:"book_author" form:"book_author"`
}

type CreateResponse struct {
	Message string `json:"message"`
}

type GetByIdBookRequest struct {
	Id string `param:"id" validate:"required"`
}

type GetByIdBookResponse struct {
	BookId        string    `json:"book_id"`
	UserId        string    `json:"user_id"`
	BookName      string    `json:"book_name"`
	BookPublisher string    `json:"book_publisher"`
	BookAuthor    string    `json:"book_author"`
	IsDelete      bool      `json:"is_delete"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type GetAllByIdUserRequest struct {
	Id string `param:"id" validate:"required"`
}

type GetAllByIdUserResponse struct {
	Data  []BooksGetAllByIdUserResponse `json:"item"`
	Count *int                          `json:"count"`
}

type BooksGetAllByIdUserResponse struct {
	BookId        string    `json:"book_id"`
	UserId        string    `json:"user_id"`
	UserName      string    `json:"user_name"`
	BookName      string    `json:"book_name"`
	BookPublisher string    `json:"book_publisher"`
	BookAuthor    string    `json:"book_author"`
	IsDelete      bool      `json:"is_delete"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type BooksCountByIdUserResponse struct {
	Count int `json:"count"`
}

type GetAllResponse struct {
	Data  []BooksGetAllResponse `json:"item"`
	Count *int                  `json:"count"`
}

type BooksGetAllResponse struct {
	BookId        string    `json:"book_id"`
	UserId        string    `json:"user_id"`
	BookName      string    `json:"book_name"`
	BookPublisher string    `json:"book_publisher"`
	BookAuthor    string    `json:"book_author"`
	IsDelete      bool      `json:"is_delete"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type BooksCountResponse struct {
	Count int `json:"count"`
}

type UpdateRequest struct {
	Id            string  `param:"id" validate:"required"`
	BookName      *string `json:"book_name" form:"book_name" validate:"required"`
	BookPublisher *string `json:"book_publisher" form:"book_publisher"`
	BookAuthor    *string `json:"book_author" form:"book_author"`
}

type UpdateResponse struct {
	Message string `json:"message"`
}

type DeleteRequest struct {
	Id string `param:"id" validate:"required"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}
