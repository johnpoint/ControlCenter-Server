package controller

import (
	"ControlCenter/pkg/errorHelper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func returnErrorMsg(c *gin.Context, err error, errCode int) {
	var errMsg string
	errMsg = errorHelper.GetErrMessage(err)
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    errCode,
		"message": errMsg,
	})
	c.Abort()
}
