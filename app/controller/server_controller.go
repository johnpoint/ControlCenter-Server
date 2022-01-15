package controller

import (
	"ControlCenter/app/logic/assets"
	"ControlCenter/dao/influxdbDao"
	"ControlCenter/infra"
	"ControlCenter/model/influxModel"
	"ControlCenter/model/mongoModel"
	"ControlCenter/pkg/errorHelper"
	"ControlCenter/pkg/influxDB"
	"ControlCenter/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type ServerChartReq struct {
	ID   string `json:"id,omitempty"`
	From int64  `json:"from,omitempty"`
	To   int64  `json:"to,omitempty"`
}

type ServerChartResp struct {
	XAxis  []string                       `json:"XAxis"`
	Points map[string][]*PerformancePoint `json:"Points"`
}

type PerformancePoint struct {
	Time  string  `json:"time"`
	Value float64 `json:"value"`
}

type PerformanceResult struct {
	Table       int       `json:"table"`
	Stop        time.Time `json:"_stop"`
	Time        time.Time `json:"_time"`
	Measurement string    `json:"_measurement"`
	Result      string    `json:"result"`
	Value       float64   `json:"_value"`
	Field       string    `json:"_field"`
	ServerId    string    `json:"server_id"`
	Start       time.Time `json:"_start"`
}

func ServerChartController(c *gin.Context) {
	var req ServerChartReq
	err := c.Bind(&req)
	if err != nil {
		returnErrorMsg(c, errorHelper.WarpErr(infra.ReqParseError, err))
		return
	}
	userID, exists := getUserIDFromContext(c)
	if !exists {
		returnErrorMsg(c, infra.ErrNeedVerifyInfo)
		return
	}
	model, err := assets.NewAssetsServer(c, req.ID, userID).Get()
	if err != nil {
		returnErrorMsg(c, errorHelper.WarpErr(infra.DataBaseError, err))
		return
	}
	svr, ok := model.(*mongoModel.ModelServer)
	if !ok {
		returnErrorMsg(c, errorHelper.WarpErr(infra.DataBaseError, err))
		return
	}
	if req.To == 0 {
		req.To = time.Now().Unix()
	}
	if req.From == 0 {
		req.From = time.Now().Add(-1 * time.Hour).Unix()
	}
	var query = influxDB.NewQuery(new(influxModel.ModelServerInfo).BucketName()).
		AddRange(
			fmt.Sprintf("%d", req.From),
			fmt.Sprintf("%d", req.To)).
		AddFilter(fmt.Sprintf(`fn: (r) => r["server_id"] == "%s"`, svr.ID)).QL()

	result, err := influxdbDao.GetQuery().Query(c, query)
	if err != nil {
		returnErrorMsg(c, errorHelper.WarpErr(infra.DataBaseError, err))
		return
	}

	var results []*PerformanceResult
	for result.Next() {
		itemByte, _ := jsoniter.Marshal(result.Record().Values())
		var r PerformanceResult
		err := jsoniter.Unmarshal(itemByte, &r)
		if err != nil {
			continue
		}
		results = append(results, &r)
	}

	var resp = ServerChartResp{
		Points: map[string][]*PerformancePoint{},
	}
	var xAxis = make(map[string]struct{})
	for i := range results {
		timeStr := results[i].Time.Format("2006-01-02 15:04:05")
		if _, has := xAxis[timeStr]; !has {
			resp.XAxis = append(resp.XAxis, timeStr)
			xAxis[timeStr] = struct{}{}
		}
		resp.Points[results[i].Field] = append(resp.Points[results[i].Field], &PerformancePoint{
			Value: results[i].Value, Time: timeStr,
		})
	}

	returnSuccessMsg(c, "OK", resp)
}

type SetUpNewServerReq struct {
	RemarkName string `json:"remark_name"`
}

type SetUpNewServerResp struct {
	Token    string `json:"token"`
	ServerID string `json:"server_id"`
}

func SetUpNewServer(c *gin.Context) {
	var req SetUpNewServerReq
	err := c.Bind(&req)
	if err != nil {
		returnErrorMsg(c, errorHelper.WarpErr(infra.ReqParseError, err))
		return
	}
	userID, exists := getUserIDFromContext(c)
	if !exists {
		returnErrorMsg(c, infra.ErrNeedVerifyInfo)
		return
	}
	var resp SetUpNewServerResp
	var assetsServer = assets.NewAssetsServer(c, utils.RandomString(), userID)
	err = assetsServer.Add(&mongoModel.ModelServer{
		ID:         utils.RandomString(),
		RemarkName: req.RemarkName,
		Token:      utils.RandomString(),
	})
	if err != nil {
		returnErrorMsg(c, errorHelper.WarpErr(infra.DataBaseError, err))
		return
	}
	model, err := assetsServer.Get()
	if err != nil {
		returnErrorMsg(c, errorHelper.WarpErr(infra.DataBaseError, err))
		return
	}
	svr, ok := model.(*mongoModel.ModelServer)
	if !ok {
		returnErrorMsg(c, errorHelper.WarpErr(infra.DataBaseError, err))
		return
	}
	resp.ServerID = svr.ID
	resp.Token = svr.Token
	returnSuccessMsg(c, "OK", resp)
}
