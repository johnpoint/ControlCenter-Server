package depend

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

var _ Depend = (*Api)(nil)
var _ Depend = (*GrpcClientServer)(nil)
var _ Depend = (*MongoDB)(nil)
var _ Depend = (*Redis)(nil)
var _ Depend = (*TaskConsumer)(nil)
var _ Depend = (*TaskProducer)(nil)
var _ Depend = (*TcpServer)(nil)
