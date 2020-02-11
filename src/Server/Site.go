package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func getSiteInfo(c echo.Context) error {
	user := checkAuth(c)
	if user != nil {
		site := Site{}
		if err := c.Bind(&site); err != nil {
			panic(err)
		}
		return c.JSON(http.StatusOK, getSite(site))
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
}

func addSiteInfo(c echo.Context) error {
	site := Site{}
	if err := c.Bind(&site); err != nil {
		panic(err)
	}
	if !addSite(site) {
		return c.JSON(http.StatusBadGateway, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
}

func deleteSiteInfo(c echo.Context) error {
	site := Site{}
	if err := c.Bind(&site); err != nil {
		panic(err)
	}
	if !delSite(site) {
		return c.JSON(http.StatusBadGateway, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
}
