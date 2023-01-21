package db

import (
	"context"
	"github.com/loukaspe/auth/mongo-handler/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserStore struct {
	client       *mongo.Client
	collection   *mongo.Collection
	databaseName string
}

func NewMongoUserStore(
	client *mongo.Client,
	collectionName string,
	databaseName string,
) (domain.UserDBInterface, error) {
	collection := client.Database(databaseName).Collection(collectionName)

	return &UserStore{
		client:     client,
		collection: collection,
	}, nil
}

func (userStore UserStore) CreateUser(user *domain.User) error {
	_, err := userStore.collection.InsertOne(context.TODO(), user)
	return err
}

func (userStore UserStore) UpdateUser(
	updater bson.D,
	filter bson.M,
) error {
	opts := options.Update().SetUpsert(true)
	_, err := userStore.collection.UpdateOne(context.TODO(), filter, updater, opts)
	return err
}

func (userStore UserStore) GetUser(filter bson.M) (*domain.User, error) {
	result := &domain.User{}

	err := userStore.collection.FindOne(context.TODO(), filter).Decode(result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (userStore UserStore) DeleteUser(filter bson.M) error {
	_, err := userStore.collection.DeleteOne(context.TODO(), filter)
	return err
}
