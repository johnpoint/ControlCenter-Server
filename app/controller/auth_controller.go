package controller

import (
	"ControlCenter/model/api/request"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var reqData request.LoginReq
	err := c.BindJSON(&reqData)
	if err != nil {
		returnErrorMsg(c, nil, 40001)
		return
	}
}
