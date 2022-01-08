package consumer

import (
	"ControlCenter/app/service/performance"
	"ControlCenter/app/service/producer"
	"ControlCenter/model/influxModel"
	"ControlCenter/proto/controlProto"
	"ControlCenter/proto/mqProto"
	"context"
	"github.com/golang/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
)

var funcMap = map[controlProto.ServerCommand]func(ctx context.Context, item *controlProto.CommandItem) error{
	controlProto.ServerCommand_CMD_ID_UPDATE_SERVER_INFO: func(ctx context.Context, item *controlProto.CommandItem) error {
		var data performance.Data
		err := jsoniter.Unmarshal(item.CommandBuf, &data)
		if err != nil {
			return err
		}
		var pData = influxModel.ModelServerInfo{
			CPU:    data.CPUStat.Percent[0],
			Disk:   data.DiskUsageStat.UsedPercent,
			Memory: data.VirtualMemory.UsedPercent,
			Swap:   data.SwapMemoryStat.UsedPercent,
		}
		dataItem, _ := jsoniter.Marshal(&pData)
		var performanceData = mqProto.PerformanceData{
			ServerId: item.ServerId,
			Buff:     dataItem,
		}
		dataItem, _ = proto.Marshal(&performanceData)
		var mqItem = mqProto.MQItem{
			Buff: dataItem,
		}
		mqByte, _ := proto.Marshal(&mqItem)
		producer.PerformanceProducer <- mqByte
		return nil
	},
}
