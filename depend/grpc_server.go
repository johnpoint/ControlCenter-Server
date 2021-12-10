package depend

import (
	"ControlCenter/app/service/grpcService"
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"context"
	"errors"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	Name    string
	AddFunc func(grpcServer *grpc.Server)
}

var _ bootstrap.Component = (*GrpcServer)(nil)

func (d *GrpcServer) Init(ctx context.Context) error {
	if grpcConfig, has := config.Config.GrpcConfigMap[d.Name]; has {
		err := grpcService.RunGrpcServer(grpcConfig.ServerListen, d.AddFunc)
		if err != nil {
			return err
		}
	}

	return errors.New("grpc config not found")
}
