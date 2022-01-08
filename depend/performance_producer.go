package depend

import (
	"ControlCenter/app/service/producer"
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"ControlCenter/pkg/rabbitmq"
	"context"
)

type PerformanceProducer struct{}

var _ bootstrap.Component = (*PerformanceProducer)(nil)

func (p *PerformanceProducer) Init(ctx context.Context) error {
	var err error
	producer.PerformanceProducer, err = new(rabbitmq.RabbitMQ).
		SetConfig(config.Config.PerformanceMQ).
		StartProducer()
	if err != nil {
		return err
	}
	return nil
}
