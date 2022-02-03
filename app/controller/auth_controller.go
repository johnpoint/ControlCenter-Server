package controller

import (
	"ControlCenter/infra"
	"ControlCenter/model/api/request"
	"ControlCenter/model/mongomodel"
	"ControlCenter/pkg/apimiddleware/session"
	"ControlCenter/pkg/errorhelper"
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
	var user mongomodel.ModelUser
	err = user.DB().FindOne(c, bson.M{
		"username": fmt.Sprintf("%s", reqData.Username),
	}).Decode(&user)
	if err != nil {
		returnErrorMsg(c, infra.ErrAuthInfoInvalid)
		return
	}
	if !utils.EqualPassword(reqData.Password, user.Password) {
		returnErrorMsg(c, infra.ErrAuthInfoInvalid)
		return
	}
	uuid := session.Si.NewSession(c, utils.RandomString(), user.ID)
	if len(uuid) == 0 {
		returnErrorMsg(c, infra.ErrAuthService)
		return
	}
	returnSuccessMsg(c, "", loginResp{
		Token: uuid,
	})
	return
}

func Register(c *gin.Context) {
	var reqData request.Register
	err := c.BindJSON(&reqData)
	if err != nil {
		returnErrorMsg(c, errorhelper.WarpErr(infra.ReqParseError, err))
		return
	}

	var user mongomodel.ModelUser
	err = new(mongomodel.ModelUser).DB().FindOne(c, bson.M{
		"username": reqData.Username,
	}).Decode(&user)
	if len(user.ID) == 0 {
		encryptPassword := utils.EncodePassword(reqData.Password)
		if len(encryptPassword) == 0 {
			returnErrorMsg(c, errorhelper.WarpErr(infra.DataBaseError, err))
			return
		}
		_, err = new(mongomodel.ModelUser).DB().InsertOne(c, &mongomodel.ModelUser{
			ID:       utils.RandomString(),
			Username: reqData.Username,
			Password: encryptPassword,
			Power:    mongomodel.UserPowerUser,
			Nickname: reqData.Nickname,
		})
		if err != nil {
			returnErrorMsg(c, errorhelper.WarpErr(infra.DataBaseError, err))
			return
		}
	} else {
		returnErrorMsg(c, infra.ReqSameUsernameError)
		return
	}
	returnSuccessMsg(c, "", nil)
}
