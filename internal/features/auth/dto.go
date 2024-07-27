package auth

type LoginRequest struct {
	UserEmail    *string `json:"user_email" form:"user_email" validate:"required"`
	UserPassword *string `json:"user_password" form:"user_password" validate:"required"`
}

type LoginResponse struct {
	UserId    string `json:"user_id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	IsDelete  bool   `json:"is_delete"`
	Token     Token
}

type LogoutResponse struct {
	Message string `json:"message"`
}
