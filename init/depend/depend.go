package depend

import (
	"context"
)

type Depend interface {
	Init(ctx context.Context) error
	GetEnable() bool
	SetEnable(enable bool)
	GetName() string
	GetDesc() string
}
