package book

import "time"

type Books struct {
	BookId        string    `json:"book_id" gorm:"primaryKey"`
	UserId        string    `json:"user_id"`
	BookName      string    `json:"book_name"`
	BookPublisher string    `json:"book_publisher"`
	BookAuthor    string    `json:"book_author"`
	IsDelete      bool      `json:"is_delete"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
