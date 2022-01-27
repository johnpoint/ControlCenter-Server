package mongoModel

import (
	"ControlCenter/dao/mongoDao"
	"go.mongodb.org/mongo-driver/mongo"
)

type ModelServer struct {
	ID               string              `json:"_id" bson:"_id"`
	RemarkName       string              `json:"remark_name" bson:"remark_name"`
	Uptime           int64               `json:"uptime" bson:"uptime"`
	Load             *Load               `json:"load" bson:"load"`
	Token            string              `json:"token" bson:"token"`
	NetworkInterface []*NetworkInterface `json:"network_interface" bson:"network_interface"`
	BytesSent        int64               `json:"bytes_sent" bson:"bytes_sent"`
	BytesRev         int64               `json:"bytes_rev" bson:"bytes_rev"`
	LastUpdated      int64               `json:"last_updated" bson:"last_updated"`
}

type NetworkInterface struct {
	Name    string   `json:"name"`
	Address []string `json:"address"`
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
