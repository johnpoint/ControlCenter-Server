package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func serverUpdate(c echo.Context) error {
	server := Server{}
	if err := c.Bind(&server); err != nil {
		panic(err)
	}
	if updateServer(Server{Ipv4: server.Ipv4, Token: server.Token}, server) {
		return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusBadRequest, Callback{Code: 0, Info: "ERROR"})
}

func serverGetCertificate(c echo.Context) error {
	return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
}
