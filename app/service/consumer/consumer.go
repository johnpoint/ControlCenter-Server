package consumer

import (
	"ControlCenter/app/service/performance"
	"ControlCenter/model/influxmodel"
	"ControlCenter/pkg/log"
	"ControlCenter/pkg/rabbitmq"
	"ControlCenter/proto/controlproto"
	"ControlCenter/proto/mqproto"
	"context"
	"github.com/golang/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
	"github.com/streadway/amqp"
)

// TaskConsumer 任务队列消费者
func TaskConsumer(ctx context.Context, delivery *amqp.Delivery) rabbitmq.Action {
	//defer delivery.Ack(true)
	return rabbitmq.Ack
}

// PerformanceConsumer 性能数据消费者
func PerformanceConsumer(ctx context.Context, delivery *amqp.Delivery) rabbitmq.Action {
	//defer delivery.Ack(true)
	var data mqproto.MQItem
	err := proto.Unmarshal(delivery.Body, &data)
	if err != nil {
		return rabbitmq.Ack
	}

	var performancePack mqproto.PerformanceData
	err = proto.Unmarshal(data.Buff, &performancePack)
	if err != nil {
		return rabbitmq.Ack
	}

	var serverInfo influxmodel.ModelServerInfo
	err = jsoniter.Unmarshal(performancePack.Buff, &serverInfo)
	if err != nil {
		return rabbitmq.Ack
	}

	err = performance.NewArchiver(ctx, performancePack.ServerId).SetData(&serverInfo).Save()
	if err != nil {
		return rabbitmq.Ack
	}

	return rabbitmq.Ack
}

// TcpServerConsumer 消费 tcp 服务器收到的包
func TcpServerConsumer(ctx context.Context, delivery *amqp.Delivery) rabbitmq.Action {
	//defer delivery.Ack(true)
	var cmdItem controlproto.CommandItem
	err := proto.Unmarshal(delivery.Body, &cmdItem)
	if err != nil {
		return rabbitmq.Ack
	}

	jsonItem, _ := jsoniter.Marshal(&cmdItem)
	log.Info("TcpServerConsumer", log.String("info", string(jsonItem)))
	fun, has := funcMap[cmdItem.Command]
	if has {
		err = fun(ctx, &cmdItem)
		if err != nil {
			return rabbitmq.Ack
		}
	} else {
		return rabbitmq.Ack
	}
	return rabbitmq.Ack
}
