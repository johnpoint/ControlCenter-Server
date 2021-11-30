package depend

import (
	"ControlCenter/app/service/producer"
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	rabbitmq2 "ControlCenter/pkg/rabbitmq"
	"context"
)

// TaskProducer 下发任务生产者
type TaskProducer struct{}

var _ bootstrap.Component = (*TaskProducer)(nil)

func (d *TaskProducer) Init(ctx context.Context) error {
	var err error
	producer.TaskProducer, err = new(rabbitmq2.RabbitMQ).
		SetAlarm(&rabbitmq2.DefaultAlarm{}).
		SetConfig(config.Config.TaskQueue).
		StartProducer()
	if err != nil {
		return err
	}
	return nil
}
