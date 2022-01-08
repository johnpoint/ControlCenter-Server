package server

import (
	"ControlCenter/app/service/tcpService"
	"fmt"
	"github.com/panjf2000/gnet"
	"sync"
	"time"
)

var connMap sync.Map

type connMeta struct {
	c gnet.Conn
	l time.Time
}

func RunServer(address string) {
	go func() {
		err := gnet.Serve(new(Handle), fmt.Sprintf("tcp://%s", address), gnet.WithMulticore(true), gnet.WithCodec(&tcpService.TcpCodec{}),
			gnet.WithReusePort(true),
			gnet.WithTCPKeepAlive(time.Minute*5),
			gnet.WithTCPNoDelay(gnet.TCPNoDelay))
		if err != nil {
			panic(err)
		}
	}()
}
