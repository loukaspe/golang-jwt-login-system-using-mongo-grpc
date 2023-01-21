package seed

import (
	"github.com/loukaspe/auth/mongo-handler/domain"
)

type Seeder struct {
	userService domain.UserServiceInterface
}

func NewSeeder(service domain.UserServiceInterface) *Seeder {
	return &Seeder{userService: service}
}

func (seeder Seeder) Seed() error {
	user := domain.User{
		Username: "username",
		Password: "password",
		Info:     "It Worked!",
	}

	return seeder.userService.CreateUser(&user)
}
