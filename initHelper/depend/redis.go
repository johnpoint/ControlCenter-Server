package depend

import (
	"ControlCenter/config"
	"ControlCenter/dao/redisDao"
	"context"
	goRedis "github.com/go-redis/redis/v8"
)

// Redis 初始化 Redis 客户端
type Redis struct{}

var _ Depend = (*Redis)(nil)

func (d *Redis) Init(ctx context.Context, cfg *config.ServiceConfig) error {
	redisDao.InitClient(&goRedis.Options{
		Network:            cfg.RedisConfig.Network,
		Addr:               cfg.RedisConfig.Addr,
		Username:           cfg.RedisConfig.Username,
		Password:           cfg.RedisConfig.Password,
		DB:                 cfg.RedisConfig.DB,
		MaxRetries:         cfg.RedisConfig.MaxRetries,
		MinRetryBackoff:    cfg.RedisConfig.MinRetryBackoff,
		MaxRetryBackoff:    cfg.RedisConfig.MaxRetryBackoff,
		DialTimeout:        cfg.RedisConfig.DialTimeout,
		ReadTimeout:        cfg.RedisConfig.ReadTimeout,
		WriteTimeout:       cfg.RedisConfig.WriteTimeout,
		PoolFIFO:           cfg.RedisConfig.PoolFIFO,
		PoolSize:           cfg.RedisConfig.PoolSize,
		MinIdleConns:       cfg.RedisConfig.MinIdleConns,
		MaxConnAge:         cfg.RedisConfig.MaxConnAge,
		PoolTimeout:        cfg.RedisConfig.PoolTimeout,
		IdleTimeout:        cfg.RedisConfig.IdleTimeout,
		IdleCheckFrequency: cfg.RedisConfig.IdleCheckFrequency,
	})
	return nil
}
