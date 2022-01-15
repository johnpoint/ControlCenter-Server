package mongoModel

import (
	"ControlCenter/dao/mongoDao"
	"go.mongodb.org/mongo-driver/mongo"
)

type ModelServer struct {
	ID         string `json:"_id" bson:"_id"`
	RemarkName string `json:"remark_name" bson:"remark_name"`
	IPv4       string `json:"ipv4" bson:"ipv4"`
	IPv6       string `json:"ipv6" bson:"ipv6"`
	Uptime     int64  `json:"uptime" bson:"uptime"`
	Load       *Load  `json:"load" bson:"load"`
	Token      string `json:"token" bson:"token"`
}

type Load struct {
	Load1  float64 `json:"load1" bson:"load1"`
	Load5  float64 `json:"load5" bson:"load5"`
	Load15 float64 `json:"load15" bson:"load15"`
}

func (m *ModelServer) CollectionName() string {
	return "server"
}

func (m *ModelServer) DB() *mongo.Collection {
	return mongoDao.Collection(m.CollectionName())
}
