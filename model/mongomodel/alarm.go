package mongomodel

import (
	"ControlCenter/dao/mongodao"
	"go.mongodb.org/mongo-driver/mongo"
)

type AlarmType = int32

const (
	AlarmTypeWebhook AlarmType = iota + 1
)

type Alarm struct {
	ID        string    `bson:"_id"`
	UserID    string    `bson:"user_id"`
	AlarmType AlarmType `bson:"alarm_type"`
	Enable    int       `bson:"enable"` // 0禁用 1启用
}

func (a *Alarm) CollectionName() string {
	return "alarm"
}

func (a *Alarm) DB() *mongo.Collection {
	return mongodao.Collection(a.CollectionName())
}
