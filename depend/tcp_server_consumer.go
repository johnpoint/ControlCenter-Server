package depend

import (
	"ControlCenter/app/service/consumer"
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"ControlCenter/pkg/rabbitmq"
	"context"
)

type TcpServerConsumer struct{}

var _ bootstrap.Component = (*TcpServerConsumer)(nil)

func (p *TcpServerConsumer) Init(ctx context.Context) error {
	new(rabbitmq.RabbitMQ).
		SetAlarm(&rabbitmq.DefaultAlarm{}). // TODO: 接入 telegram 告警
		SetConfig(config.Config.TcpServerMQ).
		SetHandle(consumer.TcpServerConsumer).
		StartConsumer()
	return nil
}
