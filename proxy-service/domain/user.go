package domain

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Info     string `json:"info"`
}

type LoginResponse struct {
	User User
}

type LoginClientInterface interface {
	Login(username, password string) (LoginResponse, error)
}
