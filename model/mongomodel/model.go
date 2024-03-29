package mongomodel

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Model interface {
	CollectionName() string
	DB() *mongo.Collection
	InitIndex(ctx context.Context) error
}

type DefaultModel struct{}

func (d *DefaultModel) InitIndex(ctx context.Context) error {
	return nil
}

func (d *DefaultModel) CollectionName() string {
	return ""
}

func (d *DefaultModel) DB() *mongo.Collection {
	return nil
}

// 检查是否实现接口
var _ Model = (*DefaultModel)(nil)

var _ Model = (*ModelUser)(nil)
var _ Model = (*ModelServer)(nil)
var _ Model = (*ModelAssets)(nil)
var _ Model = (*ModelTicket)(nil)
var _ Model = (*ModelAssetsOnlineRate)(nil)
