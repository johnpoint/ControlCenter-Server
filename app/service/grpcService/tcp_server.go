package grpcService

import (
	"ControlCenter/proto/tcpProto"
	"context"
)

type TcpServerService struct{}

func (t TcpServerService) Send(ctx context.Context, pack *tcpProto.TcpProtoPack) (*tcpProto.TcpProtoPack, error) {
	//TODO implement me
	panic("implement me")
}
