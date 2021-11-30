package depend

import (
	"ControlCenter/pkg/bootstrap"
	"context"
)

type TcpClient struct{}

var _ bootstrap.Component = (*TcpClient)(nil)

func (d *TcpClient) Init(ctx context.Context) error {
	return nil
}
