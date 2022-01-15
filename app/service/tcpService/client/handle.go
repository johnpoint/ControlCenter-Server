package client

import (
	"ControlCenter/app/service/tcpService"
	"fmt"
	"github.com/panjf2000/gnet"
	"time"
)

type Handle struct {
	*gnet.EventServer
}

var ListenerID string

func (h *Handle) OnInitComplete(svr gnet.Server) (action gnet.Action) {
	return
}

func (h *Handle) OnShutdown(svr gnet.Server) {
	return
}

func (h *Handle) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Println(time.Now().Format("20060102 15:04:05"), fmt.Sprintf("[OnOpened] %s", c.RemoteAddr()))
	ListenerID = tcpService.NewListener(c).ID()
	c.SetContext(tcpService.DataStruct{
		ChannelID: ListenerID,
	})
	return
}

func (h *Handle) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	fmt.Println(time.Now().Format("20060102 15:04:05"), fmt.Sprintf("[OnOpened] %s %+v", c.RemoteAddr(), err))
	return
}

func (h *Handle) PreWrite(c gnet.Conn) {
	return
}

func (h *Handle) AfterWrite(c gnet.Conn, b []byte) {
	return
}

func (h *Handle) React(packet []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	return
}

func (h *Handle) Tick() (delay time.Duration, action gnet.Action) {
	return
}
