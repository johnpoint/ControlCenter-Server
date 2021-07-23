package mongodb

import (
	"ControlCenter-Server/dao/mongoDao"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type ModelUser struct {
	ID       string `json:"_id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func (m *ModelUser) CollectionName() string {
	return "model_user"
}

func (m *ModelUser) InsertOne(ctx context.Context) error {
	_, err := mongoDao.Client(m.CollectionName()).InsertOne(ctx, bson.M{"$set": m})
	if err != nil {
		return err
	}
	return nil
}
