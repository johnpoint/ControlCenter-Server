package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Service struct {
	Id       int64
	Name     string
	Enable   string
	Disable  string
	Status   int64
	Serverid int64
}

func getService(c echo.Context) error {
	return c.JSON(http.StatusOK, Callback{Code: 200, Info: ""})
}
