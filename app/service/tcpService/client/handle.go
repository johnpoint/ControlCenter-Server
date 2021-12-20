package client

import (
	"github.com/panjf2000/gnet"
	"time"
)

type ClientHandle struct{}

func (h *ClientHandle) OnInitComplete(svr gnet.Server) (action gnet.Action) {
	return
}

func (h *ClientHandle) OnShutdown(svr gnet.Server) {
	return
}

func (h *ClientHandle) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	return
}

func (h *ClientHandle) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	return
}

func (h *ClientHandle) PreWrite(c gnet.Conn) {
	return
}

func (h *ClientHandle) AfterWrite(c gnet.Conn, b []byte) {
	return
}

func (h *ClientHandle) React(packet []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	return
}

func (h *ClientHandle) Tick() (delay time.Duration, action gnet.Action) {
	return
}
