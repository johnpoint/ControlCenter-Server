package session

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var Si = new(Session) // SessionInstance session 实例

type Session struct {
	driver Driver
	config *SessionConfig
}

func (s *Session) SetDriver(driver Driver) *Session {
	s.driver = driver
	return s
}

func (s *Session) SetConfig(config *SessionConfig) *Session {
	s.config = config
	return s
}

type Driver interface {
	Set(ctx context.Context, uuid string, expire time.Duration)
	Renew(ctx context.Context, uuid string, expire time.Duration)
	Get(ctx context.Context, uuid string) string
	Del(ctx context.Context, uuid string)
}

func SessionMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		sessionID := c.GetHeader(
			Si.config.HeaderName,
		)
		sessionData := Si.driver.Get(c.Request.Context(), sessionID)
		if len(sessionData) == 0 {
			c.JSON(http.StatusOK, Si.config.ReturnData)
			c.Abort()
			return
		}
		c.Set(Si.config.CtxKey, sessionData)
		Si.driver.Renew(c.Request.Context(), sessionID, time.Duration(Si.config.ExpireTime))
	}
}
