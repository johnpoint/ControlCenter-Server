package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"main/src/model"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func getUserInfo(c echo.Context) error {
	user := checkAuth(c)
	if user != nil {
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}

func updateUserInfo(c echo.Context) error {
	user := checkAuth(c)
	if user != nil {
		addLog("User", "updateUserInfo:{user:{mail:'"+user.Mail+"'}}", 1)
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}

func reSetPassword(c echo.Context) error {
	conf := loadConfig()
	salt := conf.Salt
	oldPass := c.Param("oldpass")
	newPass := c.Param("newpass")
	user := checkAuth(c)
	if user != nil {
		data := []byte(user.Mail + salt + oldPass)
		has := md5.Sum(data)
		oldpass := fmt.Sprintf("%x", has)
		u := getUser(model.User{Mail: user.Mail})
		if u[0].Password != oldpass {
			return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
		}
		data = []byte(user.Mail + salt + newPass)
		has = md5.Sum(data)
		newpass := fmt.Sprintf("%x", has)
		if (updateUser(u[0], model.User{Password: newpass})) {
			addLog("User", "reSetPassword:{user:{mail:'"+user.Mail+"'}}", 1)
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}

func getUserToken(c echo.Context) error {
	user := checkAuth(c)
	data := getUser(model.User{Mail: user.Mail})
	if len(data) != 0 {
		return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: data[0].Token})
	}
	return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "User Not Found"})
}

func getNewToken(c echo.Context) error {
	user := checkAuth(c)
	timeUnixNano := time.Now().UnixNano()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(timeUnixNano, 10))
	newToken := fmt.Sprintf("%x", h.Sum(nil))
	if (updateUser(model.User{Mail: user.Mail}, model.User{Token: newToken})) {
		addLog("User", "getNewToken:{user:{mail:'"+user.Mail+"'}}", 1)
		return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "ERROR"})
}

func changeLevel(c echo.Context) error {
	user := checkAuth(c)
	if user != nil {
		if getUser(model.User{Mail: user.Mail})[0].Level <= 0 {
			uid, _ := strconv.ParseInt(c.Param("uid"), 10, 64)
			level, _ := strconv.ParseInt(c.Param("level"), 10, 64)
			userTarget := getUser(model.User{ID: uid})
			if len(userTarget) != 0 {
				if userTarget[0].Level == 0 {
					return c.JSON(http.StatusForbidden, model.Callback{Code: 0, Info: "No permission"})
				}
			}
			if updateUser(model.User{ID: uid}, model.User{Level: level}) {
				addLog("User", "changeLevel:{user:{mail:'"+user.Mail+"'},target:{id:"+strconv.FormatInt(uid, 10)+",level:"+strconv.FormatInt(level, 10)+"}}", 1)
				return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
			}
			return c.JSON(http.StatusInternalServerError, model.Callback{Code: 0, Info: "ERROR"})
		}
		return c.JSON(http.StatusForbidden, model.Callback{Code: 0, Info: "No permission"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}

func getUserList(c echo.Context) error {
	user := checkAuth(c)
	if user != nil {
		if getUser(model.User{Mail: user.Mail})[0].Level <= 0 {
			users := getUser(model.User{})
			for i := 0; i < len(users); i++ {
				users[i].Password = "*********"
				users[i].Token = "*********"
			}
			return c.JSON(http.StatusOK, users)
		}
		return c.JSON(http.StatusForbidden, model.Callback{Code: 0, Info: "No permission"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}
