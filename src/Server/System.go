package main

import (
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
	conf := loadConfig()
	if user != nil {
		if user.Level <= 0 {
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
			dst, err := os.Create(conf.Database)
			if err != nil {
				return err
			}
			defer dst.Close()

			// Copy
			if _, err = io.Copy(dst, src); err != nil {
				return err
			}
			return c.JSON(http.StatusOK, Callback{Code: 0, Info: "OK"})
		} else {
			return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
		}
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
}

func getBackupFile(c echo.Context) error {
	conf := loadConfig()
	token := c.Param("token")
	getuser := User{Token: token}
	userInfo := getUser(getuser)
	if len(userInfo) == 0 {
		re := Callback{Code: 0, Info: "account or token incorrect"}
		return c.JSON(http.StatusUnauthorized, re)
	} else {
		if userInfo[0].Level <= 0 {
			return c.File(conf.Database)
		} else {
			return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
		}
	}
}

func setSetting(c echo.Context) error {
	user := checkAuth(c)
	name := c.Param("name")
	value := c.Param("value")
	config := SysConfig{Name: name, Value: value, UID: getUser(User{Mail: user.Mail})[0].ID}
	if setConfig(config) {
		return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
}

func getSetting(c echo.Context) error {
	user := checkAuth(c)
	name := c.Param("name")
	config := SysConfig{Name: name, UID: getUser(User{Mail: user.Mail})[0].ID}
	return c.JSON(http.StatusOK, getConfig(config))
}
