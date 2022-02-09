package tcpservice

import (
	"ControlCenter/pkg/log"
	"ControlCenter/pkg/utils"
	"errors"
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

func (l *Listener) ID() string {
	return l.listenerID
}

func (l *Listener) Send(b []byte) error {
	if l.rev != nil {
		l.rev <- b
		return nil
	}
	return errors.New("chan is nil")
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
			err := l.c.AsyncWrite(b)
			if err != nil {
				log.Error("Listener.RevLoop", log.String("info", err.Error()))
			}
		}
	}
}
