package controller

import (
	"ControlCenter/config"
	"ControlCenter/dao/redisdao"
	"ControlCenter/model/mongomodel"
	"ControlCenter/pkg/errorhelper"
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const PaginationDefaultPageSize = 20

type ApiResp struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginationReq struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"page_size"`
}

type PaginationResp struct {
	Total   int64       `json:"total"`
	PerPage int64       `json:"per_page"`
	Page    int64       `json:"page"`
	Data    interface{} `json:"data"`
}

func returnErrorMsg(c *gin.Context, err error) {
	var errMsg string
	errCode, errMsg := errorhelper.DecodeErr(err)
	c.JSON(http.StatusOK, ApiResp{
		Code:    int32(errCode),
		Message: errMsg,
	})
	c.Abort()
}

func returnSuccessMsg(c *gin.Context, message string, data interface{}) {
	if len(message) == 0 {
		message = "OK"
	}
	if data == nil {
		data = gin.H{}
	}
	c.JSON(http.StatusOK, ApiResp{
		Code:    0,
		Message: message,
		Data:    data,
	})
	c.Abort()
}

func getUserIDFromContext(c *gin.Context) (string, bool) {
	userID, exists := c.Get(config.Config.Session.CtxKey)
	if exists {
		return userID.(string), exists
	}
	return "", exists
}

const userInfoCache = "s:ccs-ng:cache:user_info:"

func getUserInfoByUserID(ctx context.Context, userID string) (mongomodel.ModelUser, error) {
	var u mongomodel.ModelUser
	result, _ := redisdao.GetClient().Get(ctx, fmt.Sprintf("%s%s", userInfoCache, userID)).Result()
	err := jsoniter.Unmarshal([]byte(result), &u)
	if err != nil || len(u.ID) == 0 {
		err = u.DB().FindOne(ctx, bson.M{
			"_id": userID,
		}).Decode(&u)
		if err != nil {
			return mongomodel.ModelUser{}, err
		}
		cache, _ := jsoniter.Marshal(&u)
		redisdao.GetClient().Set(ctx, fmt.Sprintf("%s%s", userInfoCache, userID), string(cache), 72*time.Hour)
	}
	return u, nil
}
