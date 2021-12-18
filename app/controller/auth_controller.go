package controller

import (
	"ControlCenter/config"
	"ControlCenter/dao/mongoDao"
	"ControlCenter/infra"
	"ControlCenter/model/api/request"
	"ControlCenter/model/mongoModel"
	"ControlCenter/pkg/apiMiddleware/session"
	"ControlCenter/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type loginResp struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	var reqData request.LoginReq
	err := c.BindJSON(&reqData)
	if err != nil {
		returnErrorMsg(c, infra.ErrAuthInfoInvalid)
		return
	}
	var user mongoModel.ModelUser
	err = mongoDao.Client(user.CollectionName()).FindOne(c, bson.M{
		"username": fmt.Sprintf("%s", reqData.Username),
		"password": utils.Md5(fmt.Sprintf("%s%s", reqData.Password, config.Config.Salt)),
	}).Decode(&user)
	if err != nil {
		returnErrorMsg(c, infra.ErrAuthInfoInvalid)
		return
	}
	uuid := session.Si.NewSession(c, utils.RandomString(), user.ID)
	if len(uuid) == 0 {
		returnErrorMsg(c, infra.ErrAuthServise)
		return
	}
	c.SetCookie("SESSION", uuid, 0, "/", config.Config.URL, true, true)
	returnSuccessMsg(c, "", loginResp{
		Token: uuid,
	})
	return
}
