package depend

import (
	"ControlCenter/app/service/tcpservice/server"
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"context"
)

type TcpServer struct{}

var _ bootstrap.Component = (*TcpServer)(nil)

func (d *TcpServer) Init(ctx context.Context) error {
	server.RunServer(config.Config.TcpServerListen)
	return nil
}
