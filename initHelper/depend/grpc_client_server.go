package depend

import (
	"ControlCenter/config"
	"ControlCenter/initHelper/depend/grpcClient"
	"context"
)

type GrpcClientServer struct{}

func (g *GrpcClientServer) Init(ctx context.Context, cfg *config.ServiceConfig) error {
	err := grpcClient.AddClient("server", cfg.GrpcClientServer)
	if err != nil {
		return err
	}
	return nil
}
