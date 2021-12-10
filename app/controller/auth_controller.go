package controller

import (
	"ControlCenter/dao/mongoDao"
	"ControlCenter/model/api/request"
	"ControlCenter/model/mongoModel"
	"ControlCenter/pkg/errorHelper"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var reqData request.LoginReq
	err := c.BindJSON(&reqData)
	if err != nil {
		returnErrorMsg(c, errorHelper.ErrAuthInfoInvalid)
		return
	}
	var user mongoModel.ModelUser
	mongoDao.Client(user.CollectionName())
}
