package tcpService

import (
	"github.com/panjf2000/gnet"
	"time"
)

func RunServer(address string) {
	go func() {
		err := gnet.Serve(new(TcpServer), "tcp://"+address, gnet.WithMulticore(true), gnet.WithCodec(&TcpCodec{}),
			gnet.WithReusePort(true), gnet.WithTCPKeepAlive(time.Second))
		if err != nil {
			panic(err)
		}
	}()
}

type TcpServer struct {
	*gnet.EventServer
}

func (t *TcpServer) OnInitComplete(server gnet.Server) (action gnet.Action) {
	return
}

func (t *TcpServer) OnShutdown(server gnet.Server) {
	return
}

func (t *TcpServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	c.SetContext(DataStruct{})
	return
}

func (t *TcpServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	return
}

func (t *TcpServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	return
}
