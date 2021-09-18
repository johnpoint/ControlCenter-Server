package mongoDao

import (
	"ControlCenter/config"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var MongoClientNotInit = errors.New("mongo client not init")

var MongoClient *mongo.Client

func InitMongoClient(config *config.MongoDBConfig) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.URL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func getMongoClient() (*mongo.Client, error) {
	if MongoClient == nil {
		return nil, MongoClientNotInit
	}
	return MongoClient, nil
}

func Client(collection string) *mongo.Collection {
	if client, err := getMongoClient(); err == nil {
		return client.Database(config.Config.MongoDBConfig.Database).Collection(collection)
	} else {
		panic(err)
	}
}
