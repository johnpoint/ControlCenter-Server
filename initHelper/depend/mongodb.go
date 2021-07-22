package depend

import (
	"ControlCenter-Server/config"
	"ControlCenter-Server/dao/mongoDao"
	"context"
)

type MongoDB struct {
	Enable bool
}

func (r *MongoDB) Init(ctx context.Context) error {
	mongoDao.InitMongoClient(&config.Config.MongoDBConfig)
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
