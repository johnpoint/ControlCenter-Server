package mongomodel

import (
	"ControlCenter/dao/mongodao"
	"go.mongodb.org/mongo-driver/mongo"
)

type AssetsType int32

const (
	AuthorityTypeRead = iota + 1
	AuthorityTypeWrite
)

const (
	AssetsTypeServer AssetsType = iota + 1
)

type Authority struct {
	UserID string `json:"user_id" bson:"user_id"`
	Type   int    `json:"type" bson:"type"`
}

type ModelAssets struct {
	ID         string       `json:"id" bson:"_id"`
	RemarkName string       `json:"remark_name" bson:"remark_name"`
	AssetsType AssetsType   `json:"assets_type" bson:"assets_type"`
	Owner      string       `json:"owner" bson:"owner"`
	Authority  []*Authority `json:"authority" bson:"authority"`
	CreateAt   int64        `json:"create_at" bson:"create_at"`
}

func (a *ModelAssets) CollectionName() string {
	return "assets"
}

func (a *ModelAssets) DB() *mongo.Collection {
	return mongodao.Collection(a.CollectionName())
}
