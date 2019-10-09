package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

const salt = "NFUCA"

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Mail  string `json:"mail"`
	Level int64  `json:"level"`
	jwt.StandardClaims
}

func oaLogin(c echo.Context) error {
	u := User{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	data := []byte(u.Mail + salt + u.Password)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	getuser := User{Mail: u.Mail}
	user := getUser(getuser)
	if len(user) == 0 {
		re := Callback{Code: 0, Info: "account or password incorrect"}
		return c.JSON(http.StatusOK, re)
	}
	if user[0].Password == md5str1 {
		claims := &jwtCustomClaims{
			user[0].Username,
			user[0].Mail,
			user[0].Level,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("NFUCA"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "account or password incorrect"})

}

func oaRegister(c echo.Context) error {
	u := User{}
	var re Callback
	if err := c.Bind(&u); err != nil {
		return err
	}
	checkUser := getUser(User{Mail: u.Mail})
	if len(checkUser) != 0 {
		re = Callback{Code: 0, Info: "This account has been used"}
	} else {
		data := []byte(u.Mail + salt + u.Password)
		has := md5.Sum(data)
		md5str1 := fmt.Sprintf("%x", has)
		newUser := User{Username: u.Username, Mail: u.Mail, Password: md5str1, Level: 1}
		if addUser(newUser) {
			re = Callback{Code: 200, Info: "OK"}
		} else {
			re = Callback{Code: 0, Info: "ERROR"}
		}
	}

	return c.JSON(http.StatusOK, re)
}

func oaGetJwt(user User) string {
	return ""
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func checkAuth(c echo.Context) jwt.Claims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims
	return claims
}

func checkPower(c echo.Context) error {
	return c.JSON(http.StatusOK, checkAuth(c))
}
