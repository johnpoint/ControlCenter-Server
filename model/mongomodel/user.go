package mongomodel

import (
	"ControlCenter/dao/mongodao"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserPower = uint64

const (
	UserPowerGuest     = iota + 1 // 游客
	UserPowerUser                 // 用户
	UserPowerOperation            // 运营人员
	UserPowerAdmin                // 管理员
	UserPowerSystem               // 系统
)

type ModelUser struct {
	ID       string    `json:"_id" bson:"_id"`
	Username string    `json:"username" bson:"username"`
	Password string    `json:"password" bson:"password"`
	Power    UserPower `json:"power" bson:"power"`
	Nickname string    `json:"nickname" bson:"nickname"`
}

func (m *ModelUser) InitIndex(ctx context.Context) error {
	return nil
}

func (m *ModelUser) CollectionName() string {
	return "user"
}

func (m *ModelUser) DB() *mongo.Collection {
	return mongodao.Collection(m.CollectionName())
}

func (m *ModelUser) InsertOne(ctx context.Context) error {
	_, err := mongodao.Collection(m.CollectionName()).InsertOne(ctx, bson.M{"$set": m})
	if err != nil {
		return err
	}
	return nil
}
