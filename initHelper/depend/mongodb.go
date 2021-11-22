package depend

import (
	"ControlCenter/config"
	"ControlCenter/dao/mongoDao"
	"context"
)

// MongoDB 初始化 MongoDB 客户端
type MongoDB struct{}

func (d *MongoDB) Init(ctx context.Context, cfg *config.ServiceConfig) error {
	mongoDao.InitMongoClient(cfg.MongoDBConfig)
	return nil
}
