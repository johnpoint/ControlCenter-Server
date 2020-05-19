package apis

import (
	"github.com/labstack/echo"
	"main/src/model"
	"net/http"
	"strconv"
)

func getDockerInfo(c echo.Context) error {
	user := checkAuth(c)
	userInfo := getUser(model.User{Mail: user.Mail})
	if len(userInfo) != 0 {
		id := c.Param("id")
		if userInfo[0].Level <= 0 {
			if id == "all" {
				docker := getDocker(model.Docker{})
				return c.JSON(http.StatusOK, docker)
			}
		}
		did, _ := strconv.ParseInt(id, 10, 64)
		docker := getDocker(model.Docker{UID: userInfo[0].ID, ID: did})
		return c.JSON(http.StatusOK, docker)
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}

func addDockerInfo(c echo.Context) error {
	user := checkAuth(c)
	userInfo := getUser(model.User{Mail: user.Mail})
	if len(userInfo) != 0 {
		docker := model.Docker{}
		if err := c.Bind(&docker); err != nil {
			return err
		}
		docker.UID = userInfo[0].ID
		if addDocker(docker) {
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusInternalServerError, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}

func editDockerInfo(c echo.Context) error {
	user := checkAuth(c)
	userInfo := getUser(model.User{Mail: user.Mail})
	if len(userInfo) != 0 {
		docker := model.Docker{}
		if err := c.Bind(&docker); err != nil {
			return err
		}
		docker.ID, _ = strconv.ParseInt(c.Param("id"), 10, 64)
		docker.UID = userInfo[0].ID
		if editDocker(docker) {
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusInternalServerError, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}
