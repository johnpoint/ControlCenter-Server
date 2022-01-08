package depend

import (
	"ControlCenter/app/service/producer"
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"ControlCenter/pkg/rabbitmq"
	"context"
)

type TcpServerProducer struct{}

var _ bootstrap.Component = (*TcpServerProducer)(nil)

func (p *TcpServerProducer) Init(ctx context.Context) error {
	var err error
	producer.TcpServerProducer, err = new(rabbitmq.RabbitMQ).
		SetConfig(config.Config.TcpServerMQ).
		StartProducer()
	if err != nil {
		return err
	}
	return nil
}
