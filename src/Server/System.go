package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func sysRestart(c echo.Context) error {
	return c.JSON(http.StatusOK, Callback{Code: 200, Info: ""})
}
