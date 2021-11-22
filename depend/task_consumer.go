package depend

import (
	"ControlCenter/app/service/consumer"
	"ControlCenter/config"
	"ControlCenter/pkg/initHelper"
	rabbitmq2 "ControlCenter/pkg/rabbitmq"
	"context"
)

// TaskConsumer 任务队列消费者
type TaskConsumer struct{}

var _ initHelper.Depend = (*TaskConsumer)(nil)

func (d *TaskConsumer) Init(ctx context.Context, cfg *config.ServiceConfig) error {
	new(rabbitmq2.RabbitMQ).
		SetAlarm(&rabbitmq2.DefaultAlarm{}).
		SetConfig(cfg.TaskQueue).
		SetHandle(consumer.TaskConsumer).
		StartConsumer()
	return nil
}
