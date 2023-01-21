package helper

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type MongoConnectionHelper struct {
	uri                 string
	db                  string
	clientInstance      *mongo.Client
	clientInstanceError error
	mongoOnce           sync.Once
}

func NewMongoConnectionHelper(
	uri,
	db string,
) *MongoConnectionHelper {
	return &MongoConnectionHelper{
		uri: uri,
		db:  db,
	}
}

func (helper *MongoConnectionHelper) GetMongoClient() (*mongo.Client, error) {
	helper.mongoOnce.Do(func() {
		helper.clientInstance = nil
		clientOptions := options.Client().ApplyURI(helper.uri)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			helper.clientInstanceError = err
			return
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			helper.clientInstanceError = err
			return
		}
		helper.clientInstance = client
	})
	return helper.clientInstance, helper.clientInstanceError
}

func (helper *MongoConnectionHelper) GetMongoDatabaseName() string {
	return helper.db
}
