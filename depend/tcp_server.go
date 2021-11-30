package depend

import (
	"ControlCenter/app/service/tcpService"
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"context"
)

type TcpServer struct{}

var _ bootstrap.Component = (*TcpServer)(nil)

func (d *TcpServer) Init(ctx context.Context) error {
	tcpService.RunServer(config.Config.TcpServerListen)
	return nil
}
