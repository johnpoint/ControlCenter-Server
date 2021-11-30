package depend

import (
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"ControlCenter/pkg/grpcClient"
	"context"
)

// GrpcClientServer 客户端连接服务端的连接
type GrpcClientServer struct{}

var _ bootstrap.Component = (*GrpcClientServer)(nil)

func (d *GrpcClientServer) Init(ctx context.Context) error {
	err := grpcClient.AddClient("server", config.Config.GrpcClientServer)
	if err != nil {
		return err
	}
	return nil
}
