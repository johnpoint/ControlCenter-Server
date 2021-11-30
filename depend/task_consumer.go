package depend

import (
	"ControlCenter/app/service/consumer"
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	rabbitmq2 "ControlCenter/pkg/rabbitmq"
	"context"
)

// TaskConsumer 任务队列消费者
type TaskConsumer struct{}

var _ bootstrap.Component = (*TaskConsumer)(nil)

func (d *TaskConsumer) Init(ctx context.Context) error {
	new(rabbitmq2.RabbitMQ).
		SetAlarm(&rabbitmq2.DefaultAlarm{}).
		SetConfig(config.Config.TaskQueue).
		SetHandle(consumer.TaskConsumer).
		StartConsumer()
	return nil
}
