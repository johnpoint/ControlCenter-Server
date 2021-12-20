package client

import (
	"ControlCenter/app/service/tcpService"
	"github.com/panjf2000/gnet"
)

var tcpClient *gnet.Client

func InitClient() {
	client, err := gnet.NewClient(&ClientHandle{}, gnet.WithCodec(&tcpService.TcpCodec{}))
	if err != nil {
		return
	}
	tcpClient = client
	return
}
