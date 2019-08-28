package main

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

const salt = "NFUCA"

func oaLogin(c echo.Context) error {
	u := User{}
	var re Callback
	if err := c.Bind(&u); err != nil {
		return err
	}
	data := []byte(u.Mail + salt + u.Password)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	getuser := User{Mail: u.Mail}
	user := getUser(getuser)
	if len(user) == 0 {
		re = Callback{Code: 0, Info: "ERROR"}
		return c.JSON(http.StatusOK, re)
	}
	if user[0].Password == md5str1 {
		re = Callback{Code: 200, Info: "OK"}
	} else {
		re = Callback{Code: 0, Info: "account or password incorrect"}
	}
	return c.JSON(http.StatusOK, re)
}

func oaRegister(c echo.Context) error {
	u := User{}
	var re Callback
	if err := c.Bind(&u); err != nil {
		return err
	}
	checkUser := getUser(User{Mail: u.Mail})
	if len(checkUser) != 0 {
		re = Callback{Code: 0, Info: "This account has been used"}
	} else {
		data := []byte(u.Mail + salt + u.Password)
		has := md5.Sum(data)
		md5str1 := fmt.Sprintf("%x", has)
		newUser := User{Username: u.Username, Mail: u.Mail, Password: md5str1, Level: 1}
		if addUser(newUser) {
			re = Callback{Code: 200, Info: "OK"}
		} else {
			re = Callback{Code: 0, Info: "ERROR"}
		}
	}

	return c.JSON(http.StatusOK, re)
}

func oaGetJwt(user User) string {
	return ""
}
