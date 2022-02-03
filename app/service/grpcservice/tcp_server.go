package grpcservice

import (
	"ControlCenter/proto/tcpproto"
	"context"
)

type TcpServerService struct{}

func (t TcpServerService) Send(ctx context.Context, pack *tcpproto.TcpProtoPack) (*tcpproto.TcpProtoPack, error) {
	//TODO implement me
	panic("implement me")
}
