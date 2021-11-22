package depend

import (
	"ControlCenter/app/service/producer"
	"ControlCenter/config"
	rabbitmq2 "ControlCenter/pkg/rabbitmq"
	"context"
)

// TaskProducer 下发任务生产者
type TaskProducer struct{}

func (d *TaskProducer) Init(ctx context.Context, cfg *config.ServiceConfig) error {
	var err error
	producer.TaskProducer, err = new(rabbitmq2.RabbitMQ).
		SetAlarm(&rabbitmq2.DefaultAlarm{}).
		SetConfig(cfg.TaskQueue).
		StartProducer()
	if err != nil {
		return err
	}
	return nil
}
