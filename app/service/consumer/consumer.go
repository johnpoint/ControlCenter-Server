package consumer

import (
	"ControlCenter/app/service/performance"
	"ControlCenter/model/influxModel"
	"ControlCenter/proto/controlProto"
	"ControlCenter/proto/mqProto"
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
	"github.com/streadway/amqp"
	"time"
)

// TaskConsumer 任务队列消费者
func TaskConsumer(ctx context.Context, delivery *amqp.Delivery) error {
	defer delivery.Ack(true)
	return nil
}

// PerformanceConsumer 性能数据消费者
func PerformanceConsumer(ctx context.Context, delivery *amqp.Delivery) error {
	defer delivery.Ack(true)
	var data mqProto.MQItem
	err := proto.Unmarshal(delivery.Body, &data)
	if err != nil {
		return err
	}

	var performancePack mqProto.PerformanceData
	err = proto.Unmarshal(data.Buff, &performancePack)
	if err != nil {
		return err
	}

	var serverInfo influxModel.ModelServerInfo
	err = jsoniter.Unmarshal(performancePack.Buff, &serverInfo)
	if err != nil {
		return err
	}

	err = performance.NewArchiver(ctx, performancePack.ServerId).SetData(&serverInfo).Save()
	if err != nil {
		return err
	}

	return nil
}

// TcpServerConsumer 消费 tcp 服务器收到的包
func TcpServerConsumer(ctx context.Context, delivery *amqp.Delivery) error {
	defer delivery.Ack(true)
	var cmdItem controlProto.CommandItem
	err := proto.Unmarshal(delivery.Body, &cmdItem)
	if err != nil {
		return err
	}

	jsonItem, _ := jsoniter.Marshal(&cmdItem)
	fmt.Println(time.Now().Format("20060102 15:04:05"), string(jsonItem))
	fun, has := funcMap[cmdItem.Command]
	if has {
		err = fun(ctx, &cmdItem)
		if err != nil {
			return err
		}
	} else {
		return errors.New("func not found")
	}
	return nil
}
