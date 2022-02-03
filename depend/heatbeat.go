package depend

import (
	"ControlCenter/app/service/tcpservice"
	tcpClient "ControlCenter/app/service/tcpservice/client"
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"ControlCenter/pkg/log"
	"ControlCenter/proto/controlproto"
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/shirou/gopsutil/host"
	"time"
)

type HeartBeat struct{}

var _ bootstrap.Component = (*HeartBeat)(nil)

func (d *HeartBeat) Init(ctx context.Context) error {
	go heartBeatLoop()
	return nil
}

func heartBeatLoop() {
	time.Sleep(10 * time.Second)
	var uptime uint64
	var heartItem controlproto.HeatBeat
	var loopTime = -1
	for {
		if loopTime >= config.Config.HeartBeatFixInterval || loopTime == -1 {
			loopTime = 0
			serverInfo, err := host.Info()
			if err != nil {
				panic(err)
			}
			uptime = serverInfo.Uptime
		}
		heartItem.Uptime = uptime
		itemByte, _ := proto.Marshal(&heartItem)
		var pack = controlproto.CommandItem{
			Command:    controlproto.ServerCommand_CMD_ID_HEARTBEAT,
			ServerId:   config.Config.ServerID,
			CommandBuf: itemByte,
		}
		itemByte, _ = proto.Marshal(&pack)
		log.Info("heartBeatLoop", log.Uint64("info", uptime))
		err := tcpservice.GetListenerByID(tcpClient.ListenerID).Send(itemByte)
		if err != nil {
			log.Error("heartBeatLoop", log.String("info", err.Error()))
		}
		time.Sleep(time.Second * config.Config.HeartBeatDuration)
		uptime += uint64(config.Config.HeartBeatDuration)
	}
}
