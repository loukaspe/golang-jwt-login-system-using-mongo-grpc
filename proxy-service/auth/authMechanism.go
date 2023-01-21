package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/loukaspe/auth/proxy/domain"
	"os"
	"time"
)

const tokenExpirationTime = time.Hour

type AuthMech struct{}

func NewAuthMechanism() *AuthMech {
	return &AuthMech{}
}

func (j *AuthMech) CreateToken(subject string, userInfo interface{}) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		return "", errors.New("missing jwt secret key")
	}

	signingMethod := os.Getenv("JWT_SIGNING_METHOD")
	if signingMethod == "" {
		return "", errors.New("missing jwt signing method")
	}

	token := jwt.New(jwt.GetSigningMethod(signingMethod))

	expiration := time.Now().Add(tokenExpirationTime)

	token.Claims = &domain.JwtClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			Subject:   subject,
		},
		UserInfo: userInfo,
	}

	val, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return val, nil
}

func (j *AuthMech) GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		return nil, errors.New("missing jwt secret key")
	}

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
func (j *AuthMech) SetJWTClaimsContext(
	subject string,
	ctx context.Context,
	claims jwt.MapClaims,
) context.Context {
	return context.WithValue(ctx, subject, claims)
}
