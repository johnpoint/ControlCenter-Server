package apis

import (
	"github.com/johnpoint/ControlCenter-Server/src/database"
	"github.com/johnpoint/ControlCenter-Server/src/model"
	"github.com/labstack/echo"
	"net/http"
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
	//TODO
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}
