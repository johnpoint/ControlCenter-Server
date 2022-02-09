package depend

import (
	"ControlCenter/app/service/tcpservice"
	tcpClient "ControlCenter/app/service/tcpservice/client"
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"ControlCenter/pkg/log"
	"ControlCenter/pkg/utils"
	"ControlCenter/proto/controlproto"
	"context"
	"github.com/golang/protobuf/proto"
	"time"
)

type TcpClientAuth struct{}

var _ bootstrap.Component = (*TcpClientAuth)(nil)

func (d *TcpClientAuth) Init(ctx context.Context) error {
	time.Sleep(5 * time.Second)
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
	l := tcpservice.GetListenerByID(tcpClient.ListenerID)
	if l != nil {
		err := l.Send(itemByte)
		if err != nil {
			log.Error("TcpClientAuth", log.String("info", err.Error()))
		}
	}
	return nil
}
