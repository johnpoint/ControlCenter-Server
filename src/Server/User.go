package main

import (
	"crypto/md5"
	"fmt"
	"io"
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
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
}

func updateUserInfo(c echo.Context) error {
	user := checkAuth(c)
	if user != nil {
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
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
		u := getUser(User{Mail: user.Mail})
		if u[0].Password != oldpass {
			return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
		}
		data = []byte(user.Mail + salt + newPass)
		has = md5.Sum(data)
		newpass := fmt.Sprintf("%x", has)
		if (updateUser(u[0], User{Password: newpass})) {
			return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusOK, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
}

func getUserToken(c echo.Context) error {
	user := checkAuth(c)
	data := getUser(User{Mail: user.Mail})
	if len(data) != 0 {
		return c.JSON(http.StatusOK, Callback{Code: 200, Info: data[0].Token})
	}
	return c.JSON(http.StatusOK, Callback{Code: 0, Info: "User Not Found"})
}

func getNewToken(c echo.Context) error {
	user := checkAuth(c)
	timeUnixNano := time.Now().UnixNano()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(timeUnixNano, 10))
	newToken := fmt.Sprintf("%x", h.Sum(nil))
	if (updateUser(User{Mail: user.Mail}, User{Token: newToken})) {
		return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusOK, Callback{Code: 0, Info: "ERROR"})
}

func changeLevel(c echo.Context) error {
	user := checkAuth(c)
	if user != nil {
		if getUser(User{Mail: user.Mail})[0].Level <= 0 {
			uid, _ := strconv.ParseInt(c.Param("uid"), 10, 64)
			level, _ := strconv.ParseInt(c.Param("level"), 10, 64)
			userTarget := getUser(User{ID: uid})
			if len(userTarget) != 0 {
				if userTarget[0].Level == 0 {
					return c.JSON(http.StatusForbidden, Callback{Code: 0, Info: "No permission"})
				}
			}
			if updateUser(User{ID: uid}, User{Level: level}) {
				return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
			}
			return c.JSON(http.StatusInternalServerError, Callback{Code: 0, Info: "ERROR"})
		}
		return c.JSON(http.StatusForbidden, Callback{Code: 0, Info: "No permission"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
}

func getUserList(c echo.Context) error {
	user := checkAuth(c)
	if user != nil {
		if getUser(User{Mail: user.Mail})[0].Level <= 0 {
			users := getUser(User{})
			for i := 0; i < len(users); i++ {
				users[i].Password = "*********"
				users[i].Token = "*********"
			}
			return c.JSON(http.StatusOK, users)
		}
		return c.JSON(http.StatusForbidden, Callback{Code: 0, Info: "No permission"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
}
