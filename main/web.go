package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func getUpdate(c echo.Context) error {
	user := checkAuth(c).(jwt.MapClaims)
	if user["level"].(float64) == 1 {
		return c.JSON(http.StatusOK, getServer(Server{Hostname: "*"}))
	}
	return echo.ErrUnauthorized
}
