package server

import (
	"ControlCenter/app/service/producer"
	"ControlCenter/app/service/tcpservice"
	"ControlCenter/dao/redisdao"
	"ControlCenter/pkg/log"
	"ControlCenter/pkg/utils"
	"ControlCenter/proto/controlproto"
	"context"
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
	log.Info("tcpServer", log.String("step", "[OnOpened]"+c.RemoteAddr().String()))
	cID := utils.RandomString()
	c.SetContext(tcpservice.DataStruct{
		ChannelID: cID,
	})
	connMap.Store(cID, &connMeta{
		c: c,
		l: time.Now().Add(30 * time.Second),
	})
	return
}

func (t *Handle) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	if err != nil {
		log.Error("tcpServer", log.Strings("step", []string{"OnClosed", c.RemoteAddr().String(), err.Error()}))
	} else {
		log.Info("tcpServer", log.String("step", "OnClosed "+c.RemoteAddr().String()))
	}
	r, ok := c.Context().(tcpservice.DataStruct)
	if !ok {
		return gnet.Close
	}
	connMap.Delete(r.ChannelID)
	return
}

func (t *Handle) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	var commandItem controlproto.CommandItem
	err := proto.Unmarshal(frame, &commandItem)
	if err != nil {
		return nil, gnet.Close
	}
	r, ok := c.Context().(tcpservice.DataStruct)
	if !ok {
		return nil, gnet.Close
	}
	connMap.Store(r.ChannelID, &connMeta{
		c: c,
		l: time.Now().Add(30 * time.Second),
	})
	if !r.NotNew {
		r.NotNew = true
		tcpservice.NewListener(c)
		c.SetContext(r)
	}

	if commandItem.Command != controlproto.ServerCommand_CMD_ID_AUTH {
		result, err := redisdao.GetClient().Exists(context.Background(), fmt.Sprintf("%s%s", redisdao.ServerToken, commandItem.ServerId)).Result()
		if result == 0 {
			return
		}

		_, err = redisdao.GetClient().Set(context.Background(), fmt.Sprintf("%s%s", redisdao.ServerAliveKey, commandItem.ServerId), "", 10*time.Second).Result()
		if err != nil {
			// TODO: 日志以及告警
		}
	}

	jsonByte, _ := jsoniter.Marshal(&commandItem)
	log.Info("tcpServer", log.Strings("step", []string{"React", string(jsonByte)}))

	if producer.TcpServerProducer != nil {
		producer.TcpServerProducer <- frame
	} else {
		log.Error("tcpServer", log.Strings("step", []string{"React", "TcpServerProducer is nil"}))
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
