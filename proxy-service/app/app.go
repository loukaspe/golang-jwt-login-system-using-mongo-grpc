package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/loukaspe/auth/proxy/auth"
	"github.com/loukaspe/auth/proxy/grpc"
	"github.com/loukaspe/auth/proxy/handler"
	"github.com/loukaspe/auth/proxy/helper"
	"github.com/loukaspe/auth/proxy/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func Start() {
	router := mux.NewRouter()
	jwtClaimsRetriever := helper.NewContextJwtClaimsRetriever()

	loginGrpcClient := grpc.NewGrpcLoginClient()

	loginService := service.NewLoginService(loginGrpcClient)

	infoController := handler.NewInfoHandler(jwtClaimsRetriever)

	jwtMechanism := auth.NewAuthMechanism()
	jwtService := service.NewJwtService(jwtMechanism)
	jwtMiddleware := handler.NewAuthMiddleware(jwtMechanism)
	jwtController := handler.NewJwtClaimsHandler(jwtService, loginService)

	router.HandleFunc("/login", jwtController.JwtTokenController).Methods(http.MethodPost)

	protected := router.PathPrefix("/").Subrouter()
	protected.Use(jwtMiddleware.AddJwtAuthorization)

	protected.HandleFunc("/info", infoController.InfoController).Methods(http.MethodGet)

	err := godotenv.Load()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"errorMessage": err.Error(),
		}).Error("Error starting proxy service")
	}

	address := os.Getenv("SERVER_ADDR")
	port := os.Getenv("SERVER_PORT")

	logrus.Print(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
