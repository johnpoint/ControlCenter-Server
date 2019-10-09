package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type User struct {
	ID       int64  `gorm:"AUTO_INCREMENT"`
	Username string `json:"name" xml:"name" form:"name" query:"name"`
	Mail     string `json:"email" xml:"email" form:"email" query:"email"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
	Level    int64
}

func getUserInfo(c echo.Context) error {
	user := checkAuth(c).(jwt.MapClaims)
	if user != nil {
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "ERROR"})
}
