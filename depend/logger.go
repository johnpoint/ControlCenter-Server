package depend

import (
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"ControlCenter/pkg/log"
	"context"
)

type Logger struct{}

var _ bootstrap.Component = (*Logger)(nil)

func (d *Logger) Init(ctx context.Context) error {
	var options = make([]log.Option, 0)
	if config.Config.Environment == "local" {
		options = append(options, log.WithConsoleEncoding())
	} else {
		options = append(options, log.WithJSONEncoding())
	}
	log.OverrideLoggerWithOption(map[string]interface{}{
		"service-name": config.Config.ServiceName,
	}, options...)
	return nil
}
