package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// Site model of Site
type Site struct {
	ID     int64  `gorm:"AUTO_INCREMENT"`
	Name   string `json:"name" xml:"name" form:"name" query:"name"`
	Server int64  `json:"server" xml:"server" form:"server" query:"server"`
	Status int64  `json:"status" xml:"status" form:"status" query:"status"`
	Config string `json:"config" xml:"config" form:"config" query:"config"`
	Cer    int64  `json:"cer" xml:"cer" form:"cer" query:"cer"`
}

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
