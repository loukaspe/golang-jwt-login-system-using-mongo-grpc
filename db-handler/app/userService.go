package app

import (
	"github.com/loukaspe/auth/mongo-handler/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserService struct {
	userDb domain.UserDBInterface
}

func NewUserService(
	userDB domain.UserDBInterface,
) domain.UserServiceInterface {
	return &UserService{userDb: userDB}
}

func (userService UserService) GetUser(id string) (*domain.User, error) {
	filter := bson.M{"_id": id}

	user, err := userService.userDb.GetUser(filter)
	if err != nil {
		return &domain.User{}, err
	}

	return user, nil
}

func (userService UserService) CreateUser(user *domain.User) error {
	err := user.Validate()
	if err != nil {
		return err
	}

	user.CreatedAt = time.Now()
	user.ModifiedAt = time.Now()

	err = user.HashPassword()
	if err != nil {
		return err
	}

	return userService.userDb.CreateUser(user)
}

func (userService UserService) UpdateUser(id string, user *domain.User) error {
	err := user.Validate()
	if err != nil {
		return err
	}

	updater := bson.D{primitive.E{Key: "$set", Value: *user}}
	filter := bson.M{"_id": id}

	user.ModifiedAt = time.Now()

	return userService.userDb.UpdateUser(updater, filter)
}

func (userService UserService) DeleteUser(id string) error {
	filter := bson.M{"_id": id}

	return userService.userDb.DeleteUser(filter)
}

func (userService UserService) Login(username, password string) (*domain.User, error) {
	filter := bson.M{"username": username}

	user, err := userService.userDb.GetUser(filter)
	if err != nil {
		return user, err
	}

	err = user.CheckPassword(password)
	if err != nil {
		return &domain.User{}, err
	}

	return user, nil
}
