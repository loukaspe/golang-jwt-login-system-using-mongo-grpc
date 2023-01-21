package handler

import (
	"github.com/loukaspe/auth/proxy/domain"
	"net/http"
	"strings"
)

type AuthMechanismInterface interface {
	AddJwtAuthorization(next http.Handler) http.Handler
}

type AuthMiddleware struct {
	claimsMechanism domain.JwtClaimsMechanismInterface
}

func NewAuthMiddleware(claimsMechanism domain.JwtClaimsMechanismInterface) *AuthMiddleware {
	return &AuthMiddleware{claimsMechanism: claimsMechanism}
}
func (a *AuthMiddleware) AddJwtAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer") {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := a.claimsMechanism.GetClaimsFromToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		r = r.WithContext(
			a.claimsMechanism.SetJWTClaimsContext(
				authHeader,
				r.Context(),
				claims,
			),
		)
		next.ServeHTTP(w, r)
	})
}
