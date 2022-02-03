package influxdbdao

import (
	"ControlCenter/config"
	"ControlCenter/model/influxmodel"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

var client influxdb2.Client

func SetClient(c influxdb2.Client) {
	client = c
}

func GetClient() influxdb2.Client {
	return client
}

func GetWriteAPIBlocking(model influxmodel.Model) api.WriteAPIBlocking {
	return client.WriteAPIBlocking(config.Config.InfluxDB.Org, model.BucketName())
}

func GetQuery() api.QueryAPI {
	return client.QueryAPI(config.Config.InfluxDB.Org)
}
