package client

import (
	"ControlCenter/app/service/tcpservice"
	"ControlCenter/pkg/log"
	"errors"
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
	log.Info("OnOpened", log.String("info", c.RemoteAddr().String()))
	ListenerID = tcpservice.NewListener(c).ID()
	c.SetContext(tcpservice.DataStruct{
		ChannelID: ListenerID,
	})
	return
}

func (h *Handle) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	if err == nil {
		err = errors.New("nil")
	}
	log.Info("OnClosed", log.String("info", c.RemoteAddr().String()), log.String("err", err.Error()))
	time.Sleep(3 * time.Second)
	InitClient()
	return gnet.Shutdown
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
