package tcpService

import (
	"ControlCenter/pkg/utils"
	"github.com/panjf2000/gnet"
	"sync"
)

var ListenerMap sync.Map

func GetListenerByID(id string) *Listener {
	value, ok := ListenerMap.Load(id)
	if !ok {
		return nil
	}
	if v, ok := value.(*Listener); ok {
		return v
	}
	return nil
}

type Listener struct {
	c          gnet.Conn
	listenerID string
	rev        chan []byte
}

func NewListener(c gnet.Conn) *Listener {
	var l Listener
	l.listenerID = utils.RandomString()
	l.c = c
	go l.RevLoop()
	ListenerMap.Store(l.listenerID, &l)
	return &l
}

func (l *Listener) RevLoop() {
	defer ListenerMap.Delete(l.listenerID)
	l.rev = make(chan []byte)
	for {
		select {
		case b, ok := <-l.rev:
			if !ok {
				return
			}
			l.c.AsyncWrite(b)
		}
	}
}
