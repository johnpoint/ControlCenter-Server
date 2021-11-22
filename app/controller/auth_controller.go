package controller

import (
	"ControlCenter/model/api/request"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var reqData request.LoginReq
	err := c.BindJSON(&reqData)
	if err != nil {
		return
	}
}
