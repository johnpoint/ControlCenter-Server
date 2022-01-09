package performance

import (
	"ControlCenter/dao/influxdbDao"
	"ControlCenter/model/influxModel"
	"context"
	"errors"
	influxAPIWrite "github.com/influxdata/influxdb-client-go/v2/api/write"
	"time"
)

type Archiver struct {
	ctx      context.Context
	serverID string
	data     *influxModel.ModelServerInfo
	point    *influxAPIWrite.Point
}

func NewArchiver(ctx context.Context, id string) *Archiver {
	return &Archiver{
		serverID: id,
		ctx:      ctx,
	}
}

func (a *Archiver) check() error {
	if len(a.serverID) == 0 {
		return errors.New("server is empty")
	}
	if a.data == nil {
		return errors.New("data is nil")
	}
	return nil
}

func (a *Archiver) buildPoint() {
	pt := influxAPIWrite.NewPointWithMeasurement(a.data.Measurement())
	pt.SetTime(time.Now())
	pt.AddTag("server_id", a.serverID)
	pt.AddField("cpu", a.data.CPU)
	pt.AddField("disk", a.data.Disk)
	pt.AddField("memory", a.data.Memory)
	pt.AddField("swap", a.data.Swap)
	a.point = pt
}

func (a *Archiver) Save() error {
	if err := a.check(); err != nil {
		return err
	}
	a.buildPoint()
	err := influxdbDao.GetWriteAPIBlocking(a.data).WritePoint(
		a.ctx,
		a.point,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *Archiver) SetData(data *influxModel.ModelServerInfo) *Archiver {
	a.data = data
	return a
}
