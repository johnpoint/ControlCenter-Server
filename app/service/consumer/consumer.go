package consumer

import (
	"context"
	"github.com/streadway/amqp"
)

// TaskConsumer 任务队列消费者
func TaskConsumer(ctx context.Context, delivery *amqp.Delivery) error {
	return nil
}
