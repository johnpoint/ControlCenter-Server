package apis

import (
	. "github.com/johnpoint/ControlCenter-Server/src/auth"
	. "github.com/johnpoint/ControlCenter-Server/src/database"
	"github.com/johnpoint/ControlCenter-Server/src/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetDomainInfo(c echo.Context) error {
	user := CheckAuth(c)
	if user.Level == 1 {
		return c.JSON(http.StatusOK, GetDomain(model.Domain{}))
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func UpdateDomainInfo(c echo.Context) error {
	user := CheckAuth(c)
	if user.Level == 1 {
		domain := model.Domain{}
		if err := c.Bind(&domain); err != nil {
			panic(err)
		}
		if UpdateDomain(model.Domain{Name: domain.Name}, domain) {
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusBadRequest, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}
