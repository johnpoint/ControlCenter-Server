package depend

import (
	"ControlCenter/config"
	"ControlCenter/dao/influxdbDao"
	"ControlCenter/pkg/bootstrap"
	"context"
	"errors"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type Influxdb struct{}

var _ bootstrap.Component = (*Influxdb)(nil)

func (d *Influxdb) Init(ctx context.Context) error {
	client := influxdb2.NewClient(config.Config.InfluxDB.Address, config.Config.InfluxDB.Token)
	influxdbDao.SetClient(client)
	ping, err := client.Ping(ctx)
	if err != nil {
		return err
	}
	if !ping {
		return errors.New("no pong")
	}
	return nil
}
