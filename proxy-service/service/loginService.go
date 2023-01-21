package service

import (
	"github.com/loukaspe/auth/proxy/domain"
)

type LoginService struct {
	loginClient domain.LoginClientInterface
}

func (s LoginService) Login(username, password string) (domain.LoginResponse, error) {
	return s.loginClient.Login(username, password)
}

func NewLoginService(client domain.LoginClientInterface) *LoginService {
	return &LoginService{loginClient: client}
}
