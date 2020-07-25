package apis

import (
	"github.com/johnpoint/ControlCenter-Server/src/database"
	"github.com/johnpoint/ControlCenter-Server/src/model"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetDockerInfo(c echo.Context) error {
	user := CheckAuth(c)
	userInfo := database.GetUser(model.User{Mail: user.Mail})
	if len(userInfo) != 0 {
		id := c.Param("id")
		if userInfo[0].Level <= 0 {
			if id == "all" {
				docker := database.GetDocker(model.Docker{})
				return c.JSON(http.StatusOK, docker)
			}
		}
		did, _ := strconv.ParseInt(id, 10, 64)
		docker := database.GetDocker(model.Docker{UID: userInfo[0].ID, ID: did})
		return c.JSON(http.StatusOK, docker)
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}

func AddDockerInfo(c echo.Context) error {
	user := CheckAuth(c)
	userInfo := database.GetUser(model.User{Mail: user.Mail})
	if len(userInfo) != 0 {
		docker := model.Docker{}
		if err := c.Bind(&docker); err != nil {
			return err
		}
		docker.UID = userInfo[0].ID
		if database.AddDocker(docker) {
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusInternalServerError, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}

func EditDockerInfo(c echo.Context) error {
	user := CheckAuth(c)
	userInfo := database.GetUser(model.User{Mail: user.Mail})
	if len(userInfo) != 0 {
		docker := model.Docker{}
		if err := c.Bind(&docker); err != nil {
			return err
		}
		docker.ID, _ = strconv.ParseInt(c.Param("id"), 10, 64)
		docker.UID = userInfo[0].ID
		if database.EditDocker(docker) {
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusInternalServerError, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}
