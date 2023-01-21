package domain

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const HashCost = 10

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Username   string             `bson:"username"`
	Password   string             `bson:"password"`
	Info       string             `bson:"info"`
	CreatedAt  time.Time          `bson:"created_at"`
	ModifiedAt time.Time          `bson:"modified_at"`
}

type UserServiceInterface interface {
	Login(username, password string) (*User, error)
	CreateUser(user *User) error
	GetUser(id string) (*User, error)
	UpdateUser(id string, user *User) error
	DeleteUser(id string) error
}
type UserDBInterface interface {
	GetUser(bson.M) (*User, error)
	CreateUser(*User) error
	UpdateUser(updater bson.D, filter bson.M) error
	DeleteUser(bson.M) error
}

func (user *User) CheckPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
}

func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), HashCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) Validate() error {
	if user.Password == "" {
		return errors.New("required password")
	}
	if user.Username == "" {
		return errors.New("required username")
	}

	return nil
}
