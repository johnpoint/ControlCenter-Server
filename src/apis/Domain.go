package main

import (
	"main/src/model"
	"net/http"

	"github.com/labstack/echo"
)

func getDomainInfo(c echo.Context) error {
	user := checkAuth(c)
	if user.Level == 1 {
		return c.JSON(http.StatusOK, getDomain(model.Domain{}))
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func updateDomainInfo(c echo.Context) error {
	user := checkAuth(c)
	if user.Level == 1 {
		domain := model.Domain{}
		if err := c.Bind(&domain); err != nil {
			panic(err)
		}
		if updateDomain(model.Domain{Name: domain.Name}, domain) {
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusBadRequest, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}
