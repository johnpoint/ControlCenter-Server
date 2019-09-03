package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func getServerInfo(c echo.Context) error {
	user := checkAuth(c).(jwt.MapClaims)
	server := Server{}
	if err := c.Bind(&server); err != nil {
		panic(err)
	}
	if user["level"].(float64) == 1 {
		return c.JSON(http.StatusOK, getServer(server))
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func setupServer(c echo.Context) error {
	server := Server{}
	if err := c.Bind(&server); err != nil {
		panic(err)
	}
	check := getServer(Server{Ipv4: server.Ipv4})
	if len(check) != 0 {
		return c.JSON(http.StatusBadRequest, Callback{Code: 0, Info: "Server already exists"})
	}
	time := time.Now().Unix()
	data := []byte(strconv.FormatInt(time, 10))
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	server.Token = md5str1
	if !addServer(server) {
		return c.JSON(http.StatusBadGateway, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusOK, Callback{Code: 200, Info: md5str1})
}

func getDomainInfo(c echo.Context) error {
	user := checkAuth(c).(jwt.MapClaims)
	if user["level"].(float64) == 1 {
		return c.JSON(http.StatusOK, getDomain(Domain{}))
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func updateDomainInfo(c echo.Context) error {
	user := checkAuth(c).(jwt.MapClaims)
	if user["level"].(float64) == 1 {
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

func updateServerInfo(c echo.Context) error {
	user := checkAuth(c).(jwt.MapClaims)
	if user["level"].(float64) == 1 {
		server := Server{}
		if err := c.Bind(&server); err != nil {
			panic(err)
		}
		if updateServer(Server{Ipv4: server.Ipv4, Token: server.Token}, server) {
			return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusBadRequest, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func getUserInfo(c echo.Context) error {
	user := checkAuth(c).(jwt.MapClaims)
	if user != nil {
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
}
