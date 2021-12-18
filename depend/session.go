package depend

import (
	"ControlCenter/config"
	"ControlCenter/dao/redisDao"
	"ControlCenter/pkg/apiMiddleware/session"
	"ControlCenter/pkg/bootstrap"
	"context"
	goRedis "github.com/go-redis/redis/v8"
	"time"
)

// Session 初始化 session 组件
type Session struct{}

var _ bootstrap.Component = (*Session)(nil)

func (d *Session) Init(ctx context.Context) error {
	var err error
	client, err = redisDao.GetClient()
	if err != nil {
		return err
	}
	session.Si.
		SetDriver(&SessionDriver{}).
		SetConfig(config.Config.Session)
	return nil
}

var client *goRedis.Client

type SessionDriver struct{}

func (d *SessionDriver) Set(ctx context.Context, uuid, value string, expire time.Duration) {
	client.Set(ctx, uuid, value, expire)
}

func (d *SessionDriver) Renew(ctx context.Context, uuid string, expire time.Duration) {
	client.Expire(ctx, uuid, expire)
}

func (d *SessionDriver) Get(ctx context.Context, uuid string) string {
	return client.Get(ctx, uuid).Val()
}

func (d *SessionDriver) Del(ctx context.Context, uuid string) {
	client.Del(ctx, uuid)
}
