package depend

import (
	"ControlCenter/app/service/consumer"
	"ControlCenter/config"
	"ControlCenter/initHelper/depend/rabbitmq"
	"context"
)

// TaskConsumer 任务队列消费者
type TaskConsumer struct{}

var _ Depend = (*TaskConsumer)(nil)

func (t *TaskConsumer) Init(ctx context.Context, cfg *config.ServiceConfig) error {
	new(rabbitmq.RabbitMQ).
		SetAlarm(&rabbitmq.DefaultAlarm{}).
		SetConfig(cfg.TaskQueue).
		SetHandle(consumer.TaskConsumer).
		StartConsumer()
	return nil
}
