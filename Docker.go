package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func getDockerInfo(c echo.Context) error {
	user := checkAuth(c)
	userInfo := getUser(User{Mail: user.Mail})
	if len(userInfo) != 0 {
		id := c.Param("id")
		if userInfo[0].Level <= 0 {
			if id == "all" {
				docker := getDocker(Docker{})
				return c.JSON(http.StatusOK, docker)
			}
		}
		did, _ := strconv.ParseInt(id, 10, 64)
		docker := getDocker(Docker{UID: userInfo[0].ID, ID: did})
		return c.JSON(http.StatusOK, docker)
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
}

func addDockerInfo(c echo.Context) error {
	user := checkAuth(c)
	userInfo := getUser(User{Mail: user.Mail})
	if len(userInfo) != 0 {
		docker := Docker{}
		if err := c.Bind(&docker); err != nil {
			return err
		}
		docker.UID = userInfo[0].ID
		if addDocker(docker) {
			return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusInternalServerError, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
}

func editDockerInfo(c echo.Context) error {
	user := checkAuth(c)
	userInfo := getUser(User{Mail: user.Mail})
	if len(userInfo) != 0 {
		docker := Docker{}
		if err := c.Bind(&docker); err != nil {
			return err
		}
		docker.ID, _ = strconv.ParseInt(c.Param("id"), 10, 64)
		docker.UID = userInfo[0].ID
		if editDocker(docker) {
			return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusInternalServerError, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
}
