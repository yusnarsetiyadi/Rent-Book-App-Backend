package auth

type Auth struct {
	UserId    string
	UserName  string
	UserEmail string
	IsDelete  bool
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
