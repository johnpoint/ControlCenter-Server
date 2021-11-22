package initHelper

import (
	"ControlCenter/config"
	"context"
)

type Depend interface {
	Init(ctx context.Context, cfg *config.ServiceConfig) error
}

type DefaultDepend struct{}

func (d *DefaultDepend) Init(ctx context.Context, cfg *config.ServiceConfig) error {
	return nil
}

// 检查接口是否实现
var _ Depend = (*DefaultDepend)(nil)
