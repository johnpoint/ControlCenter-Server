package controller

import (
	"ControlCenter/infra"
	"ControlCenter/model/api/request"
	"ControlCenter/model/mongoModel"
	"ControlCenter/pkg/apiMiddleware/session"
	"ControlCenter/pkg/errorHelper"
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
		returnErrorMsg(c, infra.ErrAuthServise)
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
		returnErrorMsg(c, infra.ReqParseError)
		return
	}

	var user mongoModel.ModelUser
	err = new(mongoModel.ModelUser).DB().FindOne(c, bson.M{
		"username": reqData.Username,
	}).Decode(&user)
	if len(user.ID) == 0 {
		encryptPassword := utils.EncodePassword(reqData.Password)
		if len(encryptPassword) == 0 {
			returnErrorMsg(c, errorHelper.WarpErr(infra.DataBaseError, err))
			return
		}
		_, err = new(mongoModel.ModelUser).DB().InsertOne(c, &mongoModel.ModelUser{
			ID:       utils.RandomString(),
			Username: reqData.Username,
			Password: encryptPassword,
		})
		if err != nil {
			returnErrorMsg(c, errorHelper.WarpErr(infra.DataBaseError, err))
			return
		}
	} else {
		fmt.Println(user)
		returnErrorMsg(c, infra.ReqSameUsernameError)
		return
	}
	returnSuccessMsg(c, "", nil)
}
