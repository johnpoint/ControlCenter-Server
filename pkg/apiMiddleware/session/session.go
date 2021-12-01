package session

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var Si = new(Session) // SessionInstance session 实例

var (
	NeedDriver = errors.New("need driver")
	NeedConfig = errors.New("need config")
)

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

func (s *Session) SetReturnData(data interface{}) *Session {
	s.config.ReturnData = data
	return s
}

func (s *Session) validate() error {
	if s.driver == nil {
		return NeedDriver
	}
	if s.config == nil {
		return NeedConfig
	}
	if s.config.ReturnData == nil {
		s.config.ReturnData = gin.H{"msg": "session middleware intercept"}
	}
	return nil
}

type Driver interface {
	Set(ctx context.Context, uuid string, expire time.Duration)
	Renew(ctx context.Context, uuid string, expire time.Duration)
	Get(ctx context.Context, uuid string) string
	Del(ctx context.Context, uuid string)
}

func MiddlewareFunc() func(c *gin.Context) {
	if Si.validate() != nil {
		return func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "session middleware error"})
			return
		}
	}
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
