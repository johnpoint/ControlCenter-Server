package depend

import (
	"ControlCenter/config"
	"ControlCenter/dao/mongoDao"
	"context"
)

type MongoDB struct {
	Enable bool
}

func (r *MongoDB) Init(ctx context.Context, cfg *config.ServiceConfig) error {
	mongoDao.InitMongoClient(cfg.MongoDBConfig)
	return nil
}

func (r *MongoDB) GetEnable() bool {
	return r.Enable
}

func (r *MongoDB) SetEnable(enable bool) {
	r.Enable = enable
}

func (r *MongoDB) GetName() string {
	return "MongoDB"
}

func (r *MongoDB) GetDesc() string {
	return "初始化 MongoDB 连接"
}
