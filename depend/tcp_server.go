package depend

import (
	"ControlCenter/app/service/tcpService"
	"ControlCenter/config"
	"ControlCenter/pkg/initHelper"
	"context"
)

type TcpServer struct{}

var _ initHelper.Depend = (*TcpServer)(nil)

func (d *TcpServer) Init(ctx context.Context, cfg *config.ServiceConfig) error {
	tcpService.RunServer(cfg.TcpServerListen)
	return nil
}
