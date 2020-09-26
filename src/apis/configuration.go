package apis

import (
	"ControlCenter-Server/src/database"
	"ControlCenter-Server/src/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetConfigurationInfo(c echo.Context) error {
	user := CheckAuth(c)
	conf := model.Configuration{}
	if err := c.Bind(&conf); err != nil {
		panic(err)
	}
	if user.Level <= 1 {
		conf.UID = database.GetUser(model.User{Mail: user.Mail})[0].ID
		return c.JSON(http.StatusOK, database.GetConfiguration(conf))
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func AddConfigurationInfo(c echo.Context) error {
	user := CheckAuth(c)
	conf := model.Configuration{}
	if err := c.Bind(&conf); err != nil {
		panic(err)
	}
	if user.Level <= 1 {
		if len(database.GetConfiguration(model.Configuration{UID: database.GetUser(model.User{Mail: user.Mail})[0].ID, Name: conf.Name})) != 0 {
			return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "The configuration already exists"})
		}
		conf.UID = database.GetUser(model.User{Mail: user.Mail})[0].ID
		if database.AddConfiguration(conf) {
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusInternalServerError, model.Callback{Code: 0, Info: "Database error"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func UpdataConfigurationInfo(c echo.Context) error {
	user := CheckAuth(c)
	conf := model.Configuration{}
	conf.ID, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	if err := c.Bind(&conf); err != nil {
		panic(err)
	}
	if user.Level <= 1 {
		if len(database.GetConfiguration(model.Configuration{UID: database.GetUser(model.User{Mail: user.Mail})[0].ID, ID: conf.ID})) == 0 {
			return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "This configuration not exists"})
		}
		if database.UpdateConfiguration(conf, model.Configuration{ID: conf.ID}) {
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusInternalServerError, model.Callback{Code: 0, Info: "Database error"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func DeleteConfigurationInfo(c echo.Context) error {
	//TODO
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func LinkServerConf(c echo.Context) error {
	user := CheckAuth(c)
	if user.Level <= 1 {
		sid := c.Param("ServerID")
		cid := c.Param("FileID")
		Isid, _ := strconv.ParseInt(sid, 10, 64)
		Icid, _ := strconv.ParseInt(cid, 10, 64)
		payload := model.ServerLink{ServerID: Isid, ItemID: Icid, Type: "File"}
		data := database.GetServerLinkedItem(payload)
		if len(data) == 0 {
			if (database.LinkServer(model.ServerLink{ServerID: Isid, ItemID: Icid, Type: "File"})) {
				return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
			}
			return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "ERROR"})
		} else {
			return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "already linked"})
		}
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func UnLinkServerConf(c echo.Context) error {
	user := CheckAuth(c)
	if user.Level <= 1 {
		sid := c.Param("ServerID")
		cid := c.Param("FileID")
		Isid, _ := strconv.ParseInt(sid, 10, 64)
		Icid, _ := strconv.ParseInt(cid, 10, 64)
		payload := model.ServerLink{ServerID: Isid, ItemID: Icid, Type: "File"}
		data := database.UnLinkServer(payload)
		if data {
			if len(database.GetServerLinkedItem(payload)) == 0 {
				return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
			}
		}
		return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}
