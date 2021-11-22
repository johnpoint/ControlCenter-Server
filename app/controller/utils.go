package controller

import (
	"ControlCenter/pkg/errorHelper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func returnErrorMsg(c *gin.Context, err error) {
	errCode, errMsg := errorHelper.GetErrMsg(err)
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    errCode,
		"message": errMsg,
	})
}
