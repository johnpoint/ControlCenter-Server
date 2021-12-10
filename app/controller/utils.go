package controller

import (
	"ControlCenter/pkg/errorHelper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func returnErrorMsg(c *gin.Context, err error) {
	var errMsg string
	errCode, errMsg := errorHelper.DecodeErr(err)
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    errCode,
		"message": errMsg,
	})
	c.Abort()
}
