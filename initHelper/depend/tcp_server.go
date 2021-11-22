package depend

import (
	"ControlCenter/app/service/tcpService"
	"ControlCenter/config"
	"context"
)

type TcpServer struct{}

func (d *TcpServer) Init(ctx context.Context, cfg *config.ServiceConfig) error {
	tcpService.RunServer(cfg.TcpServerListen)
	return nil
}
