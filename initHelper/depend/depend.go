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
