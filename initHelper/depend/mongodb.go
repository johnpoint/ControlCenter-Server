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
