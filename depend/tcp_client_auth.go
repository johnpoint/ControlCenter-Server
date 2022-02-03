package depend

import (
	"ControlCenter/app/service/tcpservice"
	tcpClient "ControlCenter/app/service/tcpservice/client"
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"ControlCenter/pkg/utils"
	"ControlCenter/proto/controlproto"
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"
)

type TcpClientAuth struct{}

var _ bootstrap.Component = (*TcpClientAuth)(nil)

func (d *TcpClientAuth) Init(ctx context.Context) error {
	var req = controlproto.AuthRequest{
		Token:    config.Config.Token,
		ServerId: config.Config.ServerID,
	}
	itemByte, _ := proto.Marshal(&req)
	var item = controlproto.CommandItem{
		Command:    controlproto.ServerCommand_CMD_ID_AUTH,
		CommandBuf: itemByte,
		ServerId:   config.Config.ServerID,
		SequenceId: utils.RandomString(),
	}
	itemByte, _ = proto.Marshal(&item)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "[TcpClientAuth]")
	l := tcpservice.GetListenerByID(tcpClient.ListenerID)
	if l != nil {
		err := l.Send(itemByte)
		if err != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf("[TcpClientAuth] %s", err.Error()))
		}
	}
	return nil
}
