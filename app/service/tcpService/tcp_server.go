package tcpService

import (
	"ControlCenter/pkg/utils"
	serverProto "ControlCenter/proto/server_info"
	"github.com/golang/protobuf/proto"
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
	go scanIdleConn()
	return
}

func (t *TcpServer) OnShutdown(server gnet.Server) {
	return
}

func (t *TcpServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	cID := utils.RandomString()
	c.SetContext(&DataStruct{
		channelID: cID,
	})
	connMap.Store(cID, &connMeta{
		c: c,
		l: time.Now().Add(30 * time.Second),
	})
	return
}

func (t *TcpServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	r, ok := c.Context().(*DataStruct)
	if !ok {
		return gnet.Close
	}
	connMap.Delete(r.channelID)
	return
}

func (t *TcpServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	var commandItem serverProto.CommandItem
	err := proto.Unmarshal(frame, &commandItem)
	if err != nil {
		return nil, gnet.Close
	}
	r, ok := c.Context().(*DataStruct)
	if !ok {
		return nil, gnet.Close
	}
	connMap.Store(r.channelID, &connMeta{
		c: c,
		l: time.Now().Add(30 * time.Second),
	})
	if !r.notNew {
		r.notNew = true
		NewListener(c)
	}
	return
}

func scanIdleConn() {
	for {
		var needDisConnect []gnet.Conn
		connMap.Range(func(key, value interface{}) bool {
			if r, ok := value.(connMeta); ok {
				if r.l.After(time.Now()) {
					needDisConnect = append(needDisConnect, r.c)
				}
			}
			return true
		})
		for i := range needDisConnect {
			needDisConnect[i].Close()
		}
		time.Sleep(15 * time.Second)
	}
}
