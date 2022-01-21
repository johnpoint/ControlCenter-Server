package depend

import (
	"ControlCenter/app/service/tcpService"
	tcpClient "ControlCenter/app/service/tcpService/client"
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"ControlCenter/pkg/utils"
	"ControlCenter/proto/controlProto"
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"
)

type TcpClientAuth struct{}

var _ bootstrap.Component = (*TcpClientAuth)(nil)

func (d *TcpClientAuth) Init(ctx context.Context) error {
	var req = controlProto.AuthRequest{
		Token:    config.Config.Token,
		ServerId: config.Config.ServerID,
	}
	itemByte, _ := proto.Marshal(&req)
	var item = controlProto.CommandItem{
		Command:    controlProto.ServerCommand_CMD_ID_AUTH,
		CommandBuf: itemByte,
		ServerId:   config.Config.ServerID,
		SequenceId: utils.RandomString(),
	}
	itemByte, _ = proto.Marshal(&item)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "[TcpClientAuth]")
	err := tcpService.GetListenerByID(tcpClient.ListenerID).Send(itemByte)
	if err != nil {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf("[TcpClientAuth] %s", err.Error()))
	}
	return nil
}