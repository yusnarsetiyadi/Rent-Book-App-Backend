package user

import "time"

type Users struct {
	UserId       string    `json:"user_id" gorm:"primaryKey"`
	UserName     string    `json:"user_name"`
	UserEmail    string    `json:"user_email"`
	UserPassword string    `json:"user_password"`
	IsDelete     bool      `json:"is_delete"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
