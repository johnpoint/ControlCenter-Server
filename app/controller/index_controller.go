package controller

import (
	"ControlCenter/infra"
	"ControlCenter/model/mongomodel"
	"ControlCenter/pkg/errorhelper"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type IndexResp struct {
	Assets   []*Assets `json:"assets"`
	UserInfo *UserInfo `json:"user_info"`
}

type UserInfo struct {
	Username string               `json:"username"`
	Power    mongomodel.UserPower `json:"power"`
	Nickname string               `json:"nickname"`
}

type Assets struct {
	AssetsType int64  `json:"assets_type" bson:"assets_type"`
	ID         string `json:"-" bson:"_id"`
	Num        int64  `json:"num" bson:"num"`
}

func Index(c *gin.Context) {
	var resp IndexResp
	userID, exists := getUserIDFromContext(c)
	if !exists {
		returnErrorMsg(c, infra.ErrNeedVerifyInfo)
		return
	}
	user, err := getUserInfoByUserID(c, userID)
	if err != nil {
		returnErrorMsg(c, infra.ErrNeedVerifyInfo)
		return
	}
	var assets mongomodel.ModelAssets
	var data = make([]*Assets, 0)
	cur, err := assets.DB().Aggregate(c, []bson.M{
		{
			"$match": bson.M{
				"authority": bson.M{
					"$elemMatch": bson.M{
						"user_id": userID,
					},
				},
			},
		},
		{
			"$group": bson.M{
				"_id": bson.M{
					"$toString": "$assets_type",
				},
				"assets_type": bson.M{
					"$max": "$assets_type",
				},
				"num": bson.M{
					"$sum": 1,
				},
			},
		},
	})
	if err != nil {
		returnErrorMsg(c, errorhelper.WarpErr(infra.DataBaseError, err))
		return
	}
	err = cur.All(c, &data)
	if err != nil {
		returnErrorMsg(c, errorhelper.WarpErr(infra.DataBaseError, err))
		return
	}
	resp.Assets = data
	resp.UserInfo = &UserInfo{
		Username: user.Username,
		Power:    user.Power,
		Nickname: user.Nickname,
	}

	returnSuccessMsg(c, "", resp)
}
