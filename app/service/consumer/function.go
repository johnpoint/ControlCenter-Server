package consumer

import (
	"ControlCenter/app/service/performance"
	"ControlCenter/app/service/producer"
	"ControlCenter/dao/redisDao"
	"ControlCenter/model/influxModel"
	"ControlCenter/model/mongoModel"
	cProto "ControlCenter/proto/controlProto"
	"ControlCenter/proto/mqProto"
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UpdateServerPerformanceData(ctx context.Context, item *cProto.CommandItem) error {
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
}

func ServerHeartBeat(ctx context.Context, item *cProto.CommandItem) error {
	var reqData cProto.HeatBeat
	err := proto.Unmarshal(item.CommandBuf, &reqData)
	if err != nil {
		return err
	}
	_, err = redisDao.GetClient().Set(ctx, fmt.Sprintf("%s%s", redisDao.ServerUptimeKey, item.ServerId), reqData.Uptime, 0*time.Second).Result()
	if err != nil {
		return err
	}
	return nil
}

func ServerAuth(ctx context.Context, item *cProto.CommandItem) error {
	var reqData cProto.AuthRequest
	err := proto.Unmarshal(item.CommandBuf, &reqData)
	if err != nil {
		return err
	}
	var svr mongoModel.ModelServer
	_ = svr.DB().FindOne(ctx, bson.M{
		"_id":   reqData.ServerId,
		"token": reqData.Token,
	}).Decode(&svr)
	if len(svr.ID) != 0 {
		_, err = redisDao.GetClient().Set(ctx, fmt.Sprintf("%s%s", redisDao.ServerToken, item.ServerId), reqData.Token, 7*24*time.Hour).Result()
		if err != nil {
			return err
		}
	}

	return nil
}
