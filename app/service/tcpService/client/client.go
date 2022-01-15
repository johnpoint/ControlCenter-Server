package client

import (
	"ControlCenter/app/service/tcpService"
	"ControlCenter/config"
	"fmt"
	"github.com/panjf2000/gnet"
	"time"
)

func InitClient() {
	client, err := gnet.NewClient(&Handle{},
		gnet.WithCodec(&tcpService.TcpCodec{}),
		gnet.WithTCPNoDelay(gnet.TCPNoDelay),
		gnet.WithTCPKeepAlive(time.Minute*5))
	if err != nil {
		return
	}

	err = client.Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = client.Dial("tcp", fmt.Sprintf("%s", config.Config.RemoteAddress))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
