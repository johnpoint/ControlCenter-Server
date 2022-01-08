package depend

import (
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"context"
)

type Config struct {
	Path string
}

var _ bootstrap.Component = (*Config)(nil)

func (d *Config) Init(ctx context.Context) error {
	return config.Config.SetPath(d.Path).ReadConfig()
}
