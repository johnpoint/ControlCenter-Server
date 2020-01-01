package main

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// User model of user
type User struct {
	ID       int64  `gorm:"AUTO_INCREMENT"`
	Username string `json:"name" xml:"name" form:"name" query:"name"`
	Mail     string `json:"email" xml:"email" form:"email" query:"email"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
	Level    int64
}

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
