package rent

import "time"

type Rents struct {
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
