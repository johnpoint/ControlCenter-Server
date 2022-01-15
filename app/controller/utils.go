package controller

import (
	"ControlCenter/config"
	"ControlCenter/pkg/errorHelper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResp struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Pagination struct {
	Total   int64       `json:"total"`
	PerPage int64       `json:"per_page"`
	Page    int64       `json:"page"`
	Data    interface{} `json:"data"`
}

func returnErrorMsg(c *gin.Context, err error) {
	var errMsg string
	errCode, errMsg := errorHelper.DecodeErr(err)
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
