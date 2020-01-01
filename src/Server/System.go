package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func sysRestart(c echo.Context) error {
	return c.JSON(http.StatusOK, Callback{Code: 200, Info: ""})
}

func setBackupFile(c echo.Context) error {
	user := checkAuth(c)
	if user != nil {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create("test.db")
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, Callback{Code: 0, Info: "OK"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
}

func getBackupFile(c echo.Context) error {
	conf := loadConfig()
	salt := conf.Salt
	mail := c.Param("mail")
	pass := c.Param("pass")
	data := []byte(mail + salt + pass)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	getuser := User{Mail: mail, Password: md5str1}
	userInfo := getUser(getuser)
	if len(userInfo) == 0 {
		re := Callback{Code: 0, Info: "account or password incorrect"}
		return c.JSON(http.StatusUnauthorized, re)
	}
	return c.File("test.db")
}
