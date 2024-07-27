package rent

import "time"

type CreateRequest struct {
	BookId        *string `json:"book_id" form:"book_id" validate:"required"`
	RentStartDate *string `json:"rent_start_date" form:"rent_start_date" validate:"required"`
	RentEndDate   *string `json:"rent_end_date" form:"rent_end_date" validate:"required"`
	RentQty       *int    `json:"rent_qty" form:"rent_qty" validate:"required"`
}

type CreateResponse struct {
	Message string `json:"message"`
}

type GetByIdRequest struct {
	Id string `param:"id" validate:"required"`
}

type GetByIdResponse struct {
	RentId        string    `json:"rent_id" gorm:"primaryKey"`
	UserId        string    `json:"user_id"`
	BookId        string    `json:"book_id"`
	RentStartDate time.Time `json:"rent_start_date"`
	RentEndDate   time.Time `json:"rent_end_date"`
	RentStatus    string    `json:"rent_status"`
	RentQty       int       `json:"rent_qty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type GetAllByIdUserRequest struct {
	Id string `param:"id" validate:"required"`
}

type GetAllByIdUserResponse struct {
	Data  []RentsGetAllByIdUserResponse `json:"item"`
	Count *int                          `json:"count"`
}

type RentsGetAllByIdUserResponse struct {
	RentId        string    `json:"rent_id"`
	UserId        string    `json:"user_id"`
	BookId        string    `json:"book_id"`
	UserName      string    `json:"user_name"`
	BookName      string    `json:"book_name"`
	RentStartDate time.Time `json:"rent_start_date"`
	RentEndDate   time.Time `json:"rent_end_date"`
	RentStatus    string    `json:"rent_status"`
	RentQty       int       `json:"rent_qty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type RentsCountByIdUserResponse struct {
	Count int `json:"count"`
}

type GetAllByIdBookRequest struct {
	Id string `param:"id" validate:"required"`
}

type GetAllByIdBookResponse struct {
	Data  []RentsGetAllByIdBookResponse `json:"item"`
	Count *int                          `json:"count"`
}

type RentsGetAllByIdBookResponse struct {
	RentId        string    `json:"rent_id"`
	UserId        string    `json:"user_id"`
	BookId        string    `json:"book_id"`
	UserName      string    `json:"user_name"`
	BookName      string    `json:"book_name"`
	RentStartDate time.Time `json:"rent_start_date"`
	RentEndDate   time.Time `json:"rent_end_date"`
	RentStatus    string    `json:"rent_status"`
	RentQty       int       `json:"rent_qty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type RentsCountByIdBookResponse struct {
	Count int `json:"count"`
}

type UpdateRequest struct {
	Id            string  `param:"id" validate:"required"`
	RentStartDate *string `json:"rent_start_date" form:"rent_start_date"`
	RentEndDate   *string `json:"rent_end_date" form:"rent_end_date"`
	RentQty       *int    `json:"rent_qty" form:"rent_qty"`
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
