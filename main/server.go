package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func serverUpdate(c echo.Context) error {
	token := c.Param("token")
	if (len(getServer(Server{Token: token})) == 0) {
		return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
	}
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
	token := c.Param("token")
	if (len(getServer(Server{Token: token})) == 0) {
		return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
	}
	id := c.Param("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, getCer(Certificate{ID: id64}))
}
