package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

type ContextJwtClaimsRetrieverInterface interface {
	ExtractInfo(claims interface{}) (string, error)
}

type ContextJwtClaimsRetriever struct{}

func NewContextJwtClaimsRetriever() ContextJwtClaimsRetrieverInterface {
	return &ContextJwtClaimsRetriever{}
}

func (retriever ContextJwtClaimsRetriever) ExtractInfo(claims interface{}) (string, error) {
	jwtClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("error getting jwt claims map")
	}

	userInfo, ok := jwtClaims["UserInfo"].(map[string]interface{})
	if !ok {
		return "", errors.New("error getting jwt claims user info")
	}

	info := userInfo["info"].(string)
	if info == "" {
		return "", errors.New("error getting jwt claims map")
	}

	return info, nil
}
