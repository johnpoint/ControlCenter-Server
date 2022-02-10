package controller

import (
	"ControlCenter/infra"
	"ControlCenter/model/mongomodel"
	"ControlCenter/pkg/errorhelper"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AssetsListReq struct {
	PaginationReq
}

type AssetsListItem struct {
	ID         string                `json:"id"`
	Type       mongomodel.AssetsType `json:"type"`
	Authority  int                   `json:"authority"`
	RemarkName string                `json:"remark_name"`
}

func AssetsList(c *gin.Context) {
	var req AssetsListReq
	err := c.Bind(&req)
	if err != nil {
		returnErrorMsg(c, errorhelper.WarpErr(infra.ReqParseError, err))
		return
	}
	userID, exists := getUserIDFromContext(c)
	if !exists {
		returnErrorMsg(c, infra.ErrNeedVerifyInfo)
		return
	}
	var assets mongomodel.ModelAssets
	var opt options.FindOptions
	var resp PaginationResp
	if req.Page > 0 && req.PageSize > 0 {
		opt.SetSkip((req.Page - 1) * req.PageSize)
		opt.SetLimit(req.PageSize)
		resp.Page = req.Page
		resp.PerPage = req.PageSize
	} else {
		opt.SetLimit(PaginationDefaultPageSize)
		resp.PerPage = PaginationDefaultPageSize
	}
	filter := bson.M{
		"authority": bson.M{
			"$elemMatch": bson.M{
				"user_id": userID,
			},
		},
	}
	find, err := assets.DB().Find(c, filter, &opt)
	if err != nil {
		returnErrorMsg(c, errorhelper.WarpErr(infra.DataBaseError, err))
		return
	}
	var assetsList []*mongomodel.ModelAssets
	err = find.All(c, &assetsList)
	if err != nil {
		returnErrorMsg(c, errorhelper.WarpErr(infra.DataBaseError, err))
		return
	}
	resp.Total, _ = assets.DB().CountDocuments(c, filter)
	if err != nil {
		returnErrorMsg(c, errorhelper.WarpErr(infra.DataBaseError, err))
		return
	}
	var respList []*AssetsListItem
	for i := range assetsList {
		var authority int
		for j := range assetsList[i].Authority {
			if assetsList[i].Authority[j].UserID == userID && authority < assetsList[i].Authority[j].Type {
				authority = assetsList[i].Authority[j].Type
			}
		}
		respList = append(respList, &AssetsListItem{
			ID:         assetsList[i].ID,
			RemarkName: assetsList[i].RemarkName,
			Authority:  authority,
			Type:       assetsList[i].AssetsType,
		})
	}
	resp.Data = respList
	returnSuccessMsg(c, "", &resp)
}
