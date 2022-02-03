package consumer

import (
	"ControlCenter/proto/controlproto"
	"context"
)

var funcMap = map[controlproto.ServerCommand]func(ctx context.Context, item *controlproto.CommandItem) error{
	controlproto.ServerCommand_CMD_ID_UPDATE_SERVER_INFO: UpdateServerPerformanceData,
	controlproto.ServerCommand_CMD_ID_HEARTBEAT:          ServerHeartBeat,
	controlproto.ServerCommand_CMD_ID_AUTH:               ServerAuth,
}
