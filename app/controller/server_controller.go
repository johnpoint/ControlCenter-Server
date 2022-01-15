package controller

import (
	"ControlCenter/app/logic/assets"
	"ControlCenter/dao/influxdbDao"
	"ControlCenter/dao/redisDao"
	"ControlCenter/infra"
	"ControlCenter/model/influxModel"
	"ControlCenter/model/mongoModel"
	"ControlCenter/pkg/errorHelper"
	"ControlCenter/pkg/influxDB"
	"ControlCenter/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
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
		returnErrorMsg(c, errorHelper.WarpErr(infra.ErrAssetsAuthorityError, err))
		return
	}
	svr, ok := model.(*mongoModel.ModelServer)
	if !ok {
		returnErrorMsg(c, errorHelper.WarpErr(infra.ErrAssetsAuthorityError, err))
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
		returnErrorMsg(c, errorHelper.WarpErr(infra.ErrAssetsAuthorityError, err))
		return
	}
	svr, ok := model.(*mongoModel.ModelServer)
	if !ok {
		returnErrorMsg(c, errorHelper.WarpErr(infra.ErrAssetsAuthorityError, err))
		return
	}
	resp.ServerID = svr.ID
	resp.Token = svr.Token
	returnSuccessMsg(c, "OK", resp)
}

type GetServerReq struct {
	RemarkName string           `json:"remark_name"`
	IPv4       string           `json:"ipv4"`
	IPv6       string           `json:"ipv6"`
	Load       *mongoModel.Load `json:"load"`
	State      int              `json:"state"`
}

func GetServerInfo(c *gin.Context) {
	uuid := c.Param("uuid")
	userID, exists := getUserIDFromContext(c)
	if !exists {
		returnErrorMsg(c, infra.ErrNeedVerifyInfo)
		return
	}
	model, err := assets.NewAssetsServer(c, uuid, userID).Get()
	if err != nil {
		returnErrorMsg(c, errorHelper.WarpErr(infra.ErrAssetsAuthorityError, err))
		return
	}
	svr, ok := model.(*mongoModel.ModelServer)
	if !ok {
		returnErrorMsg(c, errorHelper.WarpErr(infra.ErrAssetsAuthorityError, err))
		return
	}

	var state int
	_, err = redisDao.GetClient().Get(c, fmt.Sprintf("%s%s", redisDao.ServerAliveKey, uuid)).Result()
	if err == nil {
		state = 1
	}

	var resp = GetServerReq{
		RemarkName: svr.RemarkName,
		IPv4:       svr.IPv4,
		IPv6:       svr.IPv6,
		Load:       svr.Load,
		State:      state,
	}
	returnSuccessMsg(c, "", resp)
}

type GetServerListReq struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"page_size"`
}

type GetServerListItem struct {
	ID         string `json:"id"`
	IPv4       string `json:"ipv4"`
	IPv6       string `json:"ipv6"`
	RemarkName string `json:"remark_name"`
	Uptime     uint64 `json:"uptime"`
	Editable   bool   `json:"editable"`
	State      int    `json:"state"`
}

func GetServerList(c *gin.Context) {
	var req GetServerListReq
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
	var resp Pagination
	var assetsItem mongoModel.ModelAssets
	var assetsList []*mongoModel.ModelAssets
	opts := options.FindOptions{}
	if req.Page > 0 && req.PageSize > 0 {
		opts.SetSkip((req.Page - 1) * req.PageSize)
		opts.SetLimit(req.PageSize)
		resp.Page = req.Page
		resp.PerPage = req.PageSize
	} else {
		opts.SetLimit(PaginationDefaultPageSize)
		resp.PerPage = PaginationDefaultPageSize
	}
	var filter = bson.M{
		"assets_type": mongoModel.AssetsTypeServer,
		"authority": bson.M{
			"$elemMatch": bson.M{
				"user_id": userID,
			},
		},
	}
	cur, err := assetsItem.DB().Find(c, filter, &opts)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		returnErrorMsg(c, errorHelper.WarpErr(infra.DataBaseError, err))
		return
	}
	err = cur.All(c, &assetsList)
	if err != nil {
		returnErrorMsg(c, errorHelper.WarpErr(infra.DataBaseError, err))
		return
	}
	resp.Total, _ = assetsItem.DB().CountDocuments(c, filter)
	var editableMap = make(map[string]bool)
	var svrIDs []string
	for i := range assetsList {
		svrIDs = append(svrIDs, assetsList[i].ID)
		editableMap[assetsList[i].ID] = false
		for j := range assetsList[i].Authority {
			if assetsList[i].Authority[j].UserID == userID && assetsList[i].Authority[j].Type == mongoModel.AuthorityTypeWrite {
				editableMap[assetsList[i].ID] = true
			}
		}
	}
	if len(svrIDs) == 0 {
		returnSuccessMsg(c, "", resp)
		return
	}
	var svr mongoModel.ModelServer
	var serverList []*mongoModel.ModelServer
	cur, err = svr.DB().Find(c, bson.M{
		"_id": bson.M{
			"$in": svrIDs,
		},
	})
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		returnErrorMsg(c, errorHelper.WarpErr(infra.DataBaseError, err))
		return
	}
	err = cur.All(c, &serverList)
	if err != nil {
		returnErrorMsg(c, errorHelper.WarpErr(infra.DataBaseError, err))
		return
	}
	var respList []*GetServerListItem
	for i := range serverList {
		editable, _ := editableMap[serverList[i].ID]
		uptime, _ := redisDao.GetClient().Get(c, fmt.Sprintf("%s%s", redisDao.ServerUptimeKey, serverList[i].ID)).Result()
		fmt.Println("[uptime]", uptime)
		uptimeUint, _ := strconv.ParseUint(uptime, 10, 64)
		var state int
		_, err = redisDao.GetClient().Get(c, fmt.Sprintf("%s%s", redisDao.ServerAliveKey, serverList[i].ID)).Result()
		if err == nil {
			state = 1
		}
		respList = append(respList, &GetServerListItem{
			ID:         serverList[i].ID,
			IPv4:       serverList[i].IPv4,
			IPv6:       serverList[i].IPv6,
			RemarkName: serverList[i].RemarkName,
			Uptime:     uptimeUint,
			Editable:   editable,
			State:      state,
		})
	}
	resp.Data = respList
	returnSuccessMsg(c, "", resp)
}
