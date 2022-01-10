package mongoDao

import (
	"ControlCenter/config"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var MongoClientNotInit = errors.New("mongo client not init")

var mongoClient *mongo.Client

func InitMongoClient(config *config.MongoDBConfig) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.URL))
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	mongoClient = client
	return nil
}

func getMongoClient() (*mongo.Client, error) {
	if mongoClient == nil {
		return nil, MongoClientNotInit
	}
	return mongoClient, nil
}

func Collection(name string) *mongo.Collection {
	if client, err := getMongoClient(); err == nil {
		return client.Database(config.Config.MongoDBConfig.Database).Collection(name)
	} else {
		panic(err)
	}
}
