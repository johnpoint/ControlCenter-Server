package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"gitlab.heywoods.cn/go-sdk/omega/component/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/streadway/amqp"
)

type consumer struct {
	logger  *log.Logger
	alarm   Alarm
	channel *channel
	close   bool
	wait    *sync.WaitGroup
}

func (c *consumer) Validate() error {
	if c.channel == nil {
		return errors.New("channel is nil")
	}
	if c.channel.config == nil {
		return errors.New("channel config is nil")
	}
	if c.channel.config.QueueName == "" {
		return errors.New("queue is nil")
	}
	if c.wait == nil {
		c.wait = &sync.WaitGroup{}
	}
	return nil
}

func (c *consumer) GetConn() <-chan amqp.Delivery {
	reconnectCount := 0
	maxReconnectCount := 3
	alarmFlag := false
	for {
		time.Sleep(time.Duration(reconnectCount*reconnectCount) * time.Second)
		err := c.channel.Init()
		if err != nil {
			if reconnectCount >= maxReconnectCount {
				if !alarmFlag {
					fmt.Printf("RabbitMQ-Consumer err: %+v", err)
					if c.alarm != nil {
						_ = c.alarm.SetMsg(map[string]string{
							"Title":   "RabbitMQ 连接失败超出阈值",
							"Address": c.channel.config.Address,
							"Queue":   c.channel.config.QueueName,
						})
						_ = c.alarm.Do()
					}
					alarmFlag = true
				}
				continue
			}
			// 指数退让重试
			reconnectCount++
		} else {
			delivery, err := c.channel.Chan.Consume(
				c.channel.config.QueueName, // queue
				"",                         // consumer
				c.channel.config.AutoAck,   // auto-ack
				c.channel.config.Exclusive, // exclusive
				c.channel.config.NoLocal,   // no-local
				c.channel.config.NoWait,    // no-wait
				c.channel.config.Args,      // args
			)
			if err != nil {
				// 失败重新连接MQ
				reconnectCount++
				continue
			}
			if alarmFlag {
				fmt.Printf("RabbitMQ-Consumer err: %+v", err)
				if c.alarm != nil {
					_ = c.alarm.SetMsg(map[string]string{
						"Title":   "RabbitMQ 重连成功",
						"Address": c.channel.config.Address,
						"Queue":   c.channel.config.QueueName,
					})
				}
				_ = c.alarm.Do()
			}
			return delivery
		}
	}
}

func (c *consumer) Run(handles func(context.Context, *amqp.Delivery) error) {
RECONNECT:
	delivery := c.GetConn()

	cc := make(chan *amqp.Error)
	ctx := context.Background()
	c.channel.Chan.NotifyClose(cc)
	for {
		select {
		case msg, ok := <-delivery:
			if !ok {
				fmt.Println("RabbitMQ-Consumer err: Consumer Msg Not ok")
				if !c.close {
					goto RECONNECT
				} else {
					return
				}
			}
			if !c.close {
				c.wait.Add(1)
				err := handles(ctx, &msg)
				if err != nil {
					fmt.Printf("RabbitMQ-Consumer err: %+v\n", err)
				}
				c.wait.Done()
			} else {
				return
			}

		case <-cc:
			if !c.close {
				fmt.Printf("RabbitMQ-Consumer err: Consumer Close\n")
				goto RECONNECT
			} else {
				return
			}
		}
	}
}

func (c *consumer) gracefulShutdown() {
	// 阻塞，直到接收到shutdown的信号
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL)
	_ = <-ch
	//关闭后，Run方法中处理消息的协程将会关闭，不再处理新消息。
	c.close = true
	c.wait.Wait()
}
