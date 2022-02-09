package mongomodel

import (
	"ControlCenter/dao/mongodao"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ModelAssetsOnlineRate struct {
	ID         string     `json:"id" bson:"_id"`
	AID        string     `json:"aid" bson:"aid"`                 // 资产ID
	Rate       float64    `json:"rate" bson:"rate"`               // 存储在线率(保留两位小数*100)
	Date       int64      `json:"date" bson:"date"`               // 日期
	AssetsType AssetsType `json:"assets_type" bson:"assets_type"` // 资产类型
	DeleteAt   time.Time  `json:"delete_at" bson:"delete_at"`     // 自动删除索引
}

func (m *ModelAssetsOnlineRate) InitIndex(ctx context.Context) error {
	_, err := m.DB().Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "delete_at", Value: 1}},     // 设置索引列
			Options: options.Index().SetExpireAfterSeconds(0), // 设置过期时间
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *ModelAssetsOnlineRate) CollectionName() string {
	return "assets_online_rate"
}

func (m *ModelAssetsOnlineRate) DB() *mongo.Collection {
	return mongodao.Collection(m.CollectionName())
}
