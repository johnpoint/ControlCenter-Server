package depend

import (
	"ControlCenter/app/service/performance"
	"ControlCenter/app/service/tcpService"
	tcpClient "ControlCenter/app/service/tcpService/client"
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"ControlCenter/pkg/log"
	"ControlCenter/pkg/utils"
	"ControlCenter/proto/controlProto"
	"context"
	"github.com/golang/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type PerformanceCollector struct{}

var _ bootstrap.Component = (*PerformanceCollector)(nil)

func (p *PerformanceCollector) Init(ctx context.Context) error {
	go runPerformanceCollectLoop()
	return nil
}

func runPerformanceCollectLoop() {
	for {
		id := utils.RandomString()
		log.Info("runPerformanceCollectLoop", log.String("info", "start"), log.String("id", id))
		performanceData, err := performance.NewCollector().Do()
		if err != nil {
			log.Error("runPerformanceCollectLoop", log.String("info", err.Error()))
		} else {
			jsonByte, _ := jsoniter.Marshal(&performanceData)
			var pack = controlProto.CommandItem{
				Command:    controlProto.ServerCommand_CMD_ID_UPDATE_SERVER_INFO,
				ServerId:   config.Config.ServerID,
				CommandBuf: jsonByte,
			}
			itemByte, _ := proto.Marshal(&pack)
			err := tcpService.GetListenerByID(tcpClient.ListenerID).Send(itemByte)
			if err != nil {
				log.Error("runPerformanceCollectLoop", log.String("info", err.Error()))
			}
		}
		log.Info("runPerformanceCollectLoop", log.String("info", "finish"), log.String("id", id))
		time.Sleep(time.Duration(config.Config.CollectionInterval) * time.Second)
	}
}
