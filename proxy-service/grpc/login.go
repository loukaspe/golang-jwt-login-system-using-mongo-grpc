package grpc

import (
	"context"
	"github.com/loukaspe/auth/proxy/domain"
	mongohandler "github.com/loukaspe/auth/proxy/proto/login"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
)

const defaultLoginCallTimeout = 3 * time.Second

type GrpcLoginClientWrapper struct {
	loginResponse domain.LoginResponse
	grpcClient    mongohandler.MongoHandlerClient
}

func (wrapper *GrpcLoginClientWrapper) Login(
	username,
	password string,
) (domain.LoginResponse, error) {
	req := &mongohandler.LoginRequest{
		Username: username,
		Password: password,
	}

	conn, err := grpc.DialContext(
		context.Background(),
		os.Getenv("GRPC_MONGO_HANDLER_SERVICE_URL"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(),
	)
	defer conn.Close()
	wrapper.grpcClient = mongohandler.NewMongoHandlerClient(conn)

	ctx, cancelFunc := context.WithTimeout(
		context.Background(),
		defaultLoginCallTimeout,
	)

	defer cancelFunc()

	response, err := wrapper.grpcClient.Login(ctx, req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"errorMessage": err.Error(),
		}).Error("Error from login grpc client in proxy")
		return domain.LoginResponse{}, err
	}

	wrapper.loginResponse.User.Username = response.Username
	wrapper.loginResponse.User.Info = response.Info

	return wrapper.loginResponse, nil
}

func NewGrpcLoginClient() domain.LoginClientInterface {
	return &GrpcLoginClientWrapper{}
}
