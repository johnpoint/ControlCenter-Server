package client

import (
	"ControlCenter/app/service/tcpservice"
	"ControlCenter/config"
	"ControlCenter/pkg/log"
	"fmt"
	"github.com/panjf2000/gnet"
	"time"
)

func InitClient() {
	log.Info("InitClient", log.String("info", "InitClient"))
	client, err := gnet.NewClient(&Handle{},
		gnet.WithCodec(&tcpservice.TcpCodec{}),
		gnet.WithTCPNoDelay(gnet.TCPNoDelay),
		gnet.WithTCPKeepAlive(time.Minute*5))
	if err != nil {
		return
	}

	err = client.Start()
	if err != nil {
		log.Error("InitClient", log.String("info", err.Error()))
		return
	}

	_, err = client.Dial("tcp", fmt.Sprintf("%s", config.Config.RemoteAddress))
	if err != nil {
		log.Error("InitClient", log.String("info", err.Error()))
		return
	}

	return
}
