package grpcservice

import (
	serverInfo "ControlCenter/proto/controlproto"
	"context"
)

type PushToServerService struct{}

func (p *PushToServerService) PushTask(ctx context.Context, item *serverInfo.CommandItem) (*serverInfo.CommandItem, error) {
	//TODO implement me
	panic("implement me")
}
