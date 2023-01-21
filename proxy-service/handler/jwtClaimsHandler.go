package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/loukaspe/auth/proxy/domain"
	"github.com/loukaspe/auth/proxy/service"
	"net/http"
)

type JwtClaimsHandler struct {
	jwtService   *service.JwtService
	loginService *service.LoginService
}

func NewJwtClaimsHandler(
	jwtService *service.JwtService,
	loginService *service.LoginService,
) *JwtClaimsHandler {
	return &JwtClaimsHandler{
		jwtService:   jwtService,
		loginService: loginService,
	}
}

func (j *JwtClaimsHandler) JwtTokenController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")

	var givenUser domain.User

	err := json.NewDecoder(r.Body).Decode(&givenUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if givenUser.Username == "" || givenUser.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Empty username or password is not allowed"))
		return
	}

	loginResponse, err := j.loginService.Login(givenUser.Username, givenUser.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		logrus.WithFields(logrus.Fields{
			"errorMessage": err.Error(),
		}).Error("Error logging in proxy")
		w.Write([]byte("Not Authorized"))
		return
	}

	result, err := j.jwtService.CreateJwtTokenService(loginResponse.User)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.WithFields(logrus.Fields{
			"errorMessage": err.Error(),
		}).Error("Error generating jwt in proxy")
		w.Write([]byte("There is an error during creation of the token"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
	return

}
