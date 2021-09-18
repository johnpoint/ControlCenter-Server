package depend

import (
	"ControlCenter/config"
	"context"
)

type Depend interface {
	Init(ctx context.Context, cfg *config.ServiceConfig) error
	GetEnable() bool
	SetEnable(enable bool)
	GetName() string
	GetDesc() string
}
