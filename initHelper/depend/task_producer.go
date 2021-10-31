package depend

import (
	"ControlCenter/app/service/producer"
	"ControlCenter/config"
	"ControlCenter/initHelper/depend/rabbitmq"
	"context"
)

// TaskProducer 下发任务生产者
type TaskProducer struct{}

var _ Depend = (*TaskProducer)(nil)

func (d *TaskProducer) Init(ctx context.Context, cfg *config.ServiceConfig) error {
	var err error
	producer.TaskProducer, err = new(rabbitmq.RabbitMQ).
		SetAlarm(&rabbitmq.DefaultAlarm{}).
		SetConfig(cfg.TaskQueue).
		StartProducer()
	if err != nil {
		return err
	}
	return nil
}
