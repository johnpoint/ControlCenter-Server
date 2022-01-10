package depend

import (
	"ControlCenter/config"
	"ControlCenter/dao/mongoDao"
	"ControlCenter/pkg/bootstrap"
	"context"
)

// MongoDB 初始化 MongoDB 客户端
type MongoDB struct{}

var _ bootstrap.Component = (*MongoDB)(nil)

func (d *MongoDB) Init(ctx context.Context) error {
	err := mongoDao.InitMongoClient(config.Config.MongoDBConfig)
	if err != nil {
		return err
	}
	return nil
}
