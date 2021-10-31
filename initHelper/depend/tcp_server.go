package depend

import (
	"ControlCenter/config"
	"context"
)

type TcpServer struct{}

var _ Depend = (*TcpServer)(nil)

func (d *TcpServer) Init(ctx context.Context, cfg *config.ServiceConfig) error {
	return nil
}
