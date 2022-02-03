package depend

import (
	"ControlCenter/config"
	"ControlCenter/dao/redisdao"
	"ControlCenter/pkg/apimiddleware/session"
	"ControlCenter/pkg/bootstrap"
	"context"
	goRedis "github.com/go-redis/redis/v8"
	"time"
)

// Session 初始化 session 组件
type Session struct{}

var _ bootstrap.Component = (*Session)(nil)

func (d *Session) Init(ctx context.Context) error {
	session.Si.
		SetDriver(newSessionDriver(redisdao.GetClient())).
		SetConfig(config.Config.Session)
	return nil
}

func newSessionDriver(c *goRedis.Client) *SessionDriver {
	return &SessionDriver{
		c: c,
	}
}

type SessionDriver struct {
	c *goRedis.Client
}

func (d *SessionDriver) Set(ctx context.Context, uuid, value string, expire time.Duration) {
	d.c.Set(ctx, uuid, value, expire)
}

func (d *SessionDriver) Renew(ctx context.Context, uuid string, expire time.Duration) {
	d.c.Expire(ctx, uuid, expire)
}

func (d *SessionDriver) Get(ctx context.Context, uuid string) string {
	return d.c.Get(ctx, uuid).Val()
}

func (d *SessionDriver) Del(ctx context.Context, uuid string) {
	d.c.Del(ctx, uuid)
}
