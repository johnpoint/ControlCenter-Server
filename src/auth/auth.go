package main

import (
	"crypto/md5"
	"fmt"
	"main/src/model"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Mail  string `json:"mail"`
	Level int64  `json:"level"`
	jwt.StandardClaims
}

func oaLogin(c echo.Context) error {
	conf := loadConfig()
	salt := conf.Salt
	u := model.User{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	data := []byte(u.Mail + salt + u.Password)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	getuser := model.User{Mail: u.Mail}
	user := getUser(getuser)
	if len(user) == 0 {
		re := model.Callback{Code: 0, Info: "account or password incorrect"}
		return c.JSON(http.StatusOK, re)
	}
	if user[0].Password == md5str1 {
		claims := &jwtCustomClaims{
			user[0].Username,
			user[0].Mail,
			user[0].Level,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte(salt))
		if err != nil {
			return err
		}
		addLog("Auth", "Login:{user:{id:"+strconv.FormatInt(user[0].ID, 10)+",mail:'"+user[0].Mail+"',level:"+strconv.FormatInt(user[0].Level, 10)+"},token:'"+t+"'}", 1)
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
	return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "account or password incorrect"})

}

func oaRegister(c echo.Context) error {
	conf := loadConfig()
	salt := conf.Salt
	u := model.User{}
	var re model.Callback
	if err := c.Bind(&u); err != nil {
		return err
	}
	checkUser := getUser(model.User{Mail: u.Mail})
	if len(checkUser) != 0 {
		re = model.Callback{Code: 0, Info: "This account has been used"}
	} else {
		data := []byte(u.Mail + salt + u.Password)
		has := md5.Sum(data)
		md5str1 := fmt.Sprintf("%x", has)
		newUser := model.User{Username: u.Username, Mail: u.Mail, Password: md5str1, Level: 1}
		if addUser(newUser) {
			addLog("Auth", "Register:{user:{mail:'"+u.Mail+"'}", 1)
			re = model.Callback{Code: 200, Info: "OK"}
		} else {
			re = model.Callback{Code: 0, Info: "ERROR"}
		}
	}

	return c.JSON(http.StatusOK, re)
}

func checkAuth(c echo.Context) *jwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	if len(getUser(model.User{Mail: claims.Mail, Level: claims.Level})) == 0 {
		return nil
	}
	return claims
}

func checkPower(c echo.Context) error {
	return c.JSON(http.StatusOK, checkAuth(c))
}
