package depend

import (
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"ControlCenter/pkg/grpcClient"
	"context"
	"errors"
)

// GrpcClient 客户端连接服务端的连接
type GrpcClient struct {
	Name string
}

var _ bootstrap.Component = (*GrpcClient)(nil)

func (d *GrpcClient) Init(ctx context.Context) error {
	if grpcConfig, has := config.Config.GrpcConfigMap[d.Name]; has {
		err := grpcClient.AddClient(d.Name, grpcConfig.ClientAddress)
		if err != nil {
			return err
		}
	}

	return errors.New("grpc config not found")
}
