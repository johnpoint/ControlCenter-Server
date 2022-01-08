package depend

import (
	"ControlCenter/app/service/performance"
	"ControlCenter/app/service/tcpService"
	tcpClient "ControlCenter/app/service/tcpService/client"
	"ControlCenter/config"
	"ControlCenter/pkg/bootstrap"
	"ControlCenter/proto/controlProto"
	"context"
	"fmt"
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
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "runPerformanceCollectLoop")
		performanceData, err := performance.NewCollector().Do()
		if err != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf("[runPerformanceCollectLoop] %s", err.Error()))
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
				fmt.Println(time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf("[runPerformanceCollectLoop] %s", err.Error()))
			}
		}
		time.Sleep(time.Duration(config.Config.CollectionInterval) * time.Second)
	}
}
