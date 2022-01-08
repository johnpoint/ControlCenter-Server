package server

import (
	"ControlCenter/app/service/producer"
	"ControlCenter/app/service/tcpService"
	"ControlCenter/pkg/utils"
	"ControlCenter/proto/controlProto"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/panjf2000/gnet"
	"google.golang.org/protobuf/proto"
	"time"
)

type Handle struct {
	*gnet.EventServer
}

func (t *Handle) OnInitComplete(server gnet.Server) (action gnet.Action) {
	//go scanIdleConn()
	return
}

func (t *Handle) OnShutdown(server gnet.Server) {
	return
}

func (t *Handle) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Println(time.Now().Format("20060102 15:04:05"), fmt.Sprintf("[OnOpened] %s", c.RemoteAddr()))
	cID := utils.RandomString()
	c.SetContext(&tcpService.DataStruct{
		ChannelID: cID,
	})
	connMap.Store(cID, &connMeta{
		c: c,
		l: time.Now().Add(30 * time.Second),
	})
	return
}

func (t *Handle) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	fmt.Println(time.Now().Format("20060102 15:04:05"), fmt.Sprintf("[OnClosed] %s %+v", c.RemoteAddr(), err))
	r, ok := c.Context().(tcpService.DataStruct)
	if !ok {
		return gnet.Close
	}
	connMap.Delete(r.ChannelID)
	return
}

func (t *Handle) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Println(len(frame))
	var commandItem controlProto.CommandItem
	err := proto.Unmarshal(frame, &commandItem)
	if err != nil {
		panic(err)
		return nil, gnet.Close
	}
	r, ok := c.Context().(tcpService.DataStruct)
	if !ok {
		return nil, gnet.Close
	}
	connMap.Store(r.ChannelID, &connMeta{
		c: c,
		l: time.Now().Add(30 * time.Second),
	})
	if !r.NotNew {
		r.NotNew = true
		tcpService.NewListener(c)
		c.SetContext(r)
	}

	jsonByte, _ := jsoniter.Marshal(&commandItem)
	fmt.Println(time.Now().Format("20060102 15:04:05"), fmt.Sprintf("[React] %s", string(jsonByte)))
	if producer.TcpServerProducer != nil {
		producer.TcpServerProducer <- frame
	} else {
		fmt.Println(time.Now().Format("20060102 15:04:05"), "[React] TcpServerProducer is nil")
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
