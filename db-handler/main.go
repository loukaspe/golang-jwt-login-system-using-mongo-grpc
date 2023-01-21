package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/loukaspe/auth/mongo-handler/app"
	db "github.com/loukaspe/auth/mongo-handler/db/mongo"
	"github.com/loukaspe/auth/mongo-handler/domain"
	"github.com/loukaspe/auth/mongo-handler/helper"
	mongohandler "github.com/loukaspe/auth/mongo-handler/proto/mongo-handler"
	"github.com/loukaspe/auth/mongo-handler/rpc"
	"github.com/loukaspe/auth/mongo-handler/seed"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	mongoConnectionHelper := helper.NewMongoConnectionHelper(
		os.Getenv("MONGO_URI"),
		os.Getenv("MONGO_DB"),
	)
	mongoClient, err := mongoConnectionHelper.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}

	userDb, err := db.NewMongoUserStore(
		mongoClient,
		os.Getenv("MONGO_USERS_COLLECTION"),
		os.Getenv("MONGO_DB"),
	)

	userService := app.NewUserService(userDb)

	seeder := seed.NewSeeder(userService)

	// Go routine that initiates data to the DB
	go seedDbWithData(seeder)

	grpcServe(userService)
}

func seedDbWithData(
	seeder *seed.Seeder,
) {
	err := seeder.Seed()
	if err != nil {
		panic(err)
	}
}

func grpcServe(
	userService domain.UserServiceInterface,
) {
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)

	s := rpc.Server{
		UserService: userService,
	}

	mongohandler.RegisterMongoHandlerServer(grpcServer, s)
	log.Print(fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDR"), os.Getenv("SERVER_PORT")))
	listener, err := net.Listen(
		"tcp",
		fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDR"), os.Getenv("SERVER_PORT")),
	)
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal(err)
	}
}
