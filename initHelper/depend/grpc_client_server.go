package depend

import (
	"ControlCenter/config"
	"ControlCenter/initHelper/depend/grpcClient"
	"context"
)

// GrpcClientServer 客户端连接服务端的连接
type GrpcClientServer struct{}

func (g *GrpcClientServer) Init(ctx context.Context, cfg *config.ServiceConfig) error {
	err := grpcClient.AddClient("server", cfg.GrpcClientServer)
	if err != nil {
		return err
	}
	return nil
}
