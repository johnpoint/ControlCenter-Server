package depend

import (
	"ControlCenter/app/service/producer"
	"ControlCenter/config"
	"ControlCenter/initHelper/depend/rabbitmq"
	"context"
)

type TaskProducer struct{}

func (t *TaskProducer) Init(ctx context.Context, cfg *config.ServiceConfig) error {
	var err error
	producer.TaskProducer, err = new(rabbitmq.RabbitMQ).
		SetAlarm(&rabbitmq.DefaultAlarm{}).
		SetConfig(cfg.TaskProducer).
		StartProducer()
	if err != nil {
		return err
	}
	return nil
}
