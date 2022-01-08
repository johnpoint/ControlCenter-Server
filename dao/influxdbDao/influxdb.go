package influxdbDao

import "github.com/influxdata/influxdb-client-go/v2"

var client influxdb2.Client

func SetClient(c influxdb2.Client) {
	client = c
}

func GetClient() influxdb2.Client {
	return client
}
