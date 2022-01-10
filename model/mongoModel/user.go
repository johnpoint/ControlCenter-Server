package mongoModel

import (
	"ControlCenter/dao/mongoDao"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ModelUser struct {
	ID       string `json:"_id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func (m *ModelUser) CollectionName() string {
	return "user"
}

func (m *ModelUser) DB() *mongo.Collection {
	return mongoDao.Collection(m.CollectionName())
}

func (m *ModelUser) InsertOne(ctx context.Context) error {
	_, err := mongoDao.Collection(m.CollectionName()).InsertOne(ctx, bson.M{"$set": m})
	if err != nil {
		return err
	}
	return nil
}
