package rpc

import (
	"context"
	"github.com/loukaspe/auth/mongo-handler/domain"
	mongohandler "github.com/loukaspe/auth/mongo-handler/proto/mongo-handler"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	mongohandler.UnimplementedMongoHandlerServer
	UserService domain.UserServiceInterface
}

func (s Server) Login(
	ctx context.Context,
	loginRequest *mongohandler.LoginRequest,
) (*mongohandler.LoginResponse, error) {
	if loginRequest.Username == "" || loginRequest.Password == "" {
		return &mongohandler.LoginResponse{}, status.Error(codes.InvalidArgument, "Bad Login Token Request")
	}

	user, err := s.UserService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return &mongohandler.LoginResponse{}, err
	}

	return &mongohandler.LoginResponse{
		Username: user.Username,
		Info:     user.Info,
	}, err
}
