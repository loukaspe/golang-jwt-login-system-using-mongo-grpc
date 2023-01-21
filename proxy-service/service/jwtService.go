package service

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/loukaspe/auth/proxy/domain"
)

type JwtService struct {
	jwtClaimsMechanism domain.JwtClaimsMechanismInterface
}

func NewJwtService(jwtClaimsMechanismInterface domain.JwtClaimsMechanismInterface) *JwtService {
	return &JwtService{jwtClaimsMechanism: jwtClaimsMechanismInterface}
}

func (j *JwtService) CreateJwtTokenService(user domain.User) (string, error) {
	tokenValue, err := j.jwtClaimsMechanism.CreateToken(user.Username, user)
	if err != nil {
		return "", err
	}
	return tokenValue, nil
}

func (j *JwtService) ClaimsFromJwtTokenService(token string) (jwt.MapClaims, error) {
	claims, err := j.jwtClaimsMechanism.GetClaimsFromToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
