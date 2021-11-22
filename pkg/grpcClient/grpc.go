package grpcClient

import (
	"ControlCenter/pkg/grpcClient/grpc_retry"
	"context"
	"errors"
	"fmt"
	"gitlab.heywoods.cn/go-sdk/omega/net/grpcOption"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"sync"
	"time"
)

var clientMap = make(map[string]*Client)
var mapLock sync.Mutex

func GetClient(alias string) (*grpc.ClientConn, error) {
	if client, has := clientMap[alias]; has {
		return client.GetConn()
	}
	return nil, NilGrpcConn
}

func AddClient(alias, address string) error {
	if _, has := clientMap[alias]; has {
		return errors.New("alias already exist")
	}
	mapLock.Lock()
	defer mapLock.Unlock()
	c := Client{
		Address: address,
	}
	err := c.Init()
	if err != nil {
		return err
	}
	clientMap[alias] = &c
	return nil
}

type Client struct {
	conn    *grpc.ClientConn
	Address string
	lock    *sync.Mutex
}

func (c *Client) Init() error {
	if c.lock == nil {
		c.lock = new(sync.Mutex)
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	if len(c.Address) == 0 {
		return errors.New("address is empty")
	}
	if c.conn == nil || c.conn.GetState() == connectivity.Shutdown {
		// 调用 CRM gRPC 接口
		retryOps := []grpc_retry.CallOption{
			grpc_retry.WithMax(3),                                         // 设定最大重试次数
			grpc_retry.WithAlarm(true, connAlarm),                         // 设定重试最后一次仍然失败告警
			grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)), // 设置重试间隔
			grpc_retry.WithCodes(codes.Unavailable),                       // 设置需要重试的状态码
		}
		retryInterceptor := grpc_retry.UnaryClientInterceptor(retryOps...)
		conn, err := grpc.Dial(
			c.Address,
			grpc.WithInsecure(),
			grpcOption.DialOption(),
			grpc.WithUnaryInterceptor(retryInterceptor),
		)
		if err != nil {
			return err
		}
		c.conn = conn
	}
	return nil
}

func (c *Client) GetConn() (*grpc.ClientConn, error) {
	if c.conn == nil {
		err := c.Init()
		if err != nil {
			return nil, NilGrpcConn
		}
	}
	if c.conn.GetState() == connectivity.Shutdown {
		c.conn.Close()
		err := c.Init()
		if err != nil {
			return nil, NilGrpcConn
		}
	}
	return c.conn, nil
}

func connAlarm(ctx context.Context, lastErr error, method string) error {
	fmt.Printf("[grpc] method: %s err: %+v\n", method, lastErr)
	return nil
}

var NilGrpcConn = errors.New("grpc conn is nil")
