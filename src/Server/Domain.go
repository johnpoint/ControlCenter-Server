package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// Domain model of domain
type Domain struct {
	ID     int64  `gorm:"AUTO_INCREMENT"`
	Name   string `json:"name" xml:"name" form:"name" query:"name"`
	Status string `json:"status" xml:"status" form:"status" query:"status"`
	Cer    string `json:"cer" xml:"cer" form:"cer" query:"cer"`
	Key    string `json:"key" xml:"key" form:"key" query:"key"`
}

func getDomainInfo(c echo.Context) error {
	user := checkAuth(c)
	if user.Level == 1 {
		return c.JSON(http.StatusOK, getDomain(Domain{}))
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func updateDomainInfo(c echo.Context) error {
	user := checkAuth(c)
	if user.Level == 1 {
		domain := Domain{}
		if err := c.Bind(&domain); err != nil {
			panic(err)
		}
		if updateDomain(Domain{Name: domain.Name}, domain) {
			return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusBadRequest, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}
