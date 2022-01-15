package consumer

import (
	"ControlCenter/proto/controlProto"
	"context"
)

var funcMap = map[controlProto.ServerCommand]func(ctx context.Context, item *controlProto.CommandItem) error{
	controlProto.ServerCommand_CMD_ID_UPDATE_SERVER_INFO: UpdateServerPerformanceData,
	controlProto.ServerCommand_CMD_ID_HEARTBEAT:          ServerHeartBeat,
	controlProto.ServerCommand_CMD_ID_AUTH:               ServerAuth,
}
