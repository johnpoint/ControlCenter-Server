package depend

import (
	"ControlCenter/app/service/tcpService"
	tcpClient "ControlCenter/app/service/tcpService/client"
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"ControlCenter/proto/controlProto"
	"context"
	"fmt"
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
	var uptime uint64
	var heartItem controlProto.HeatBeat
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
		var pack = controlProto.CommandItem{
			Command:    controlProto.ServerCommand_CMD_ID_HEARTBEAT,
			ServerId:   config.Config.ServerID,
			CommandBuf: itemByte,
		}
		itemByte, _ = proto.Marshal(&pack)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "heartBeatLoop", uptime)
		err := tcpService.GetListenerByID(tcpClient.ListenerID).Send(itemByte)
		if err != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf("[heartBeatLoop] %s", err.Error()))
		}
		time.Sleep(time.Second * config.Config.HeartBeatDuration)
		uptime += uint64(config.Config.HeartBeatDuration)
	}
}
