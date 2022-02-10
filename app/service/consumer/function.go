package consumer

import (
	"ControlCenter/app/service/performance"
	"ControlCenter/app/service/producer"
	"ControlCenter/dao/redisdao"
	"ControlCenter/model/influxmodel"
	"ControlCenter/model/mongomodel"
	"ControlCenter/pkg/log"
	cProto "ControlCenter/proto/controlproto"
	"ControlCenter/proto/mqproto"
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
	"net"
	"time"
)

func UpdateServerPerformanceData(ctx context.Context, item *cProto.CommandItem) error {
	var data performance.Data
	err := jsoniter.Unmarshal(item.CommandBuf, &data)
	if err != nil {
		return err
	}
	var pData = influxmodel.ModelServerInfo{
		CPU:    data.CPUStat.Percent[0],
		Disk:   data.DiskUsageStat.UsedPercent,
		Memory: data.VirtualMemory.UsedPercent,
		Swap:   data.SwapMemoryStat.UsedPercent,
	}
	dataItem, _ := jsoniter.Marshal(&pData)
	var performanceData = mqproto.PerformanceData{
		ServerId: item.ServerId,
		Buff:     dataItem,
	}
	dataItem, _ = proto.Marshal(&performanceData)
	var mqItem = mqproto.MQItem{
		Buff: dataItem,
	}
	mqByte, _ := proto.Marshal(&mqItem)
	producer.PerformanceProducer <- mqByte

	// 十分之一的机率更新 load
	if rand.Intn(10) == 5 {
		// 流量统计
		var sent, rev uint64
		for i := range data.InterfaceStat {
			for j := range data.InterfaceStat[i].Addrs {
				if !HasLocalIP(net.ParseIP(data.InterfaceStat[i].Addrs[j].Addr)) {
					sent += data.NetInterfaceStat[i].BytesSent
					rev += data.NetInterfaceStat[i].BytesRecv
					break
				}
			}
		}

		// 更新数据库内数据
		var svr mongomodel.ModelServer
		_, err := svr.DB().UpdateOne(ctx, bson.M{
			"_id": item.ServerId,
		}, bson.M{
			"$set": bson.M{
				"load":           data.Load,
				"bytes_sent":     sent,
				"bytes_rev":      rev,
				"last_updated":   time.Now().UnixNano() / 1e6,
				"partition_stat": data.PartitionStat,
			},
		})
		if err != nil {
			log.Error("UpdateServerPerformanceData", log.String("info", err.Error()))
		}
	}

	return nil
}

func HasLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}

func ServerHeartBeat(ctx context.Context, item *cProto.CommandItem) error {
	var reqData cProto.HeatBeat
	err := proto.Unmarshal(item.CommandBuf, &reqData)
	if err != nil {
		return err
	}
	_, err = redisdao.GetClient().Set(ctx, fmt.Sprintf("%s%s", redisdao.ServerUptimeKey, item.ServerId), reqData.Uptime, 0*time.Second).Result()
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
	var svr mongomodel.ModelServer
	_ = svr.DB().FindOne(ctx, bson.M{
		"_id":   reqData.ServerId,
		"token": reqData.Token,
	}).Decode(&svr)
	if len(svr.ID) != 0 {
		_, err = redisdao.GetClient().Set(ctx, fmt.Sprintf("%s%s", redisdao.ServerToken, item.ServerId), reqData.Token, 7*24*time.Hour).Result()
		if err != nil {
			return err
		}
	}

	return nil
}
