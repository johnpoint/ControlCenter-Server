package controller

import (
	"ControlCenter/pkg/errorHelper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResp struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
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
