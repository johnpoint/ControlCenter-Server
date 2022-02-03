package depend

import (
	tcpClient "ControlCenter/app/service/tcpservice/client"
	"ControlCenter/pkg/bootstrap"
	"context"
)

type TcpClient struct{}

var _ bootstrap.Component = (*TcpClient)(nil)

func (d *TcpClient) Init(ctx context.Context) error {
	tcpClient.InitClient()
	return nil
}
