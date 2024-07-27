package user

import "time"

type CreateRequest struct {
	UserName  *string `json:"user_name" form:"user_name" validate:"required"`
	UserEmail *string `json:"user_email" form:"user_email" validate:"required"`
}

type CreateResponse struct {
	Message string `json:"message"`
}

type GetByIdRequest struct {
	Id string `param:"id" validate:"required"`
}

type GetByIdResponse struct {
	UserId    string    `json:"user_id"`
	UserName  string    `json:"user_name"`
	UserEmail string    `json:"user_email"`
	IsDelete  bool      `json:"is_delete"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetAllResponse struct {
	Data  []UsersGetAllResponse `json:"item"`
	Count *int                  `json:"count"`
}

type UsersGetAllResponse struct {
	UserId    string    `json:"user_id"`
	UserName  string    `json:"user_name"`
	UserEmail string    `json:"user_email"`
	IsDelete  bool      `json:"is_delete"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UsersCountResponse struct {
	Count int `json:"count"`
}

type UpdateRequest struct {
	Id        string  `param:"id" validate:"required"`
	UserName  *string `json:"user_name" form:"user_name" validate:"required"`
	UserEmail *string `json:"user_email" form:"user_email" validate:"required"`
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

type ChangePasswordRequest struct {
	Id          string `param:"id" validate:"required"`
	OldPassword string `json:"old_password" form:"old_password" validate:"required"`
	NewPassword string `json:"new_password" form:"new_password" validate:"required"`
}

type ChangePasswordResponse struct {
	Message string `json:"message"`
}
