package apis

import (
	"crypto/md5"
	"fmt"
	"github.com/johnpoint/ControlCenter-Server/src/config"
	"github.com/johnpoint/ControlCenter-Server/src/database"
	"github.com/johnpoint/ControlCenter-Server/src/model"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Mail  string `json:"mail"`
	Level int64  `json:"level"`
	jwt.StandardClaims
}

func OaLogin(c echo.Context) error {
	conf := config.LoadConfig()
	salt := conf.Salt
	u := model.User{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	data := []byte(u.Mail + salt + u.Password)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	getuser := model.User{Mail: u.Mail}
	user := database.GetUser(getuser)
	if len(user) == 0 {
		re := model.Callback{Code: 0, Info: "account or password incorrect"}
		return c.JSON(http.StatusOK, re)
	}
	if user[0].Password == md5str1 {
		claims := &JwtCustomClaims{
			user[0].Username,
			user[0].Mail,
			user[0].Level,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte(salt))
		if err != nil {
			return err
		}
		database.AddLog("Auth", "Login:{user:{id:"+strconv.FormatInt(user[0].ID, 10)+",mail:'"+user[0].Mail+"',level:"+strconv.FormatInt(user[0].Level, 10)+"},token:'"+t+"'}", 1)
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
	return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "account or password incorrect"})

}

func OaRegister(c echo.Context) error {
	conf := config.LoadConfig()
	salt := conf.Salt
	u := model.User{}
	var re model.Callback
	if err := c.Bind(&u); err != nil {
		return err
	}
	checkUser := database.GetUser(model.User{Mail: u.Mail})
	if len(checkUser) != 0 {
		re = model.Callback{Code: 0, Info: "This account has been used"}
	} else {
		data := []byte(u.Mail + salt + u.Password)
		has := md5.Sum(data)
		md5str1 := fmt.Sprintf("%x", has)
		newUser := model.User{Username: u.Username, Mail: u.Mail, Password: md5str1, Level: 1}
		if database.AddUser(newUser) {
			database.AddLog("Auth", "Register:{user:{mail:'"+u.Mail+"'}", 1)
			re = model.Callback{Code: 200, Info: "OK"}
		} else {
			re = model.Callback{Code: 0, Info: "ERROR"}
		}
	}

	return c.JSON(http.StatusOK, re)
}

func CheckAuth(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	if len(database.GetUser(model.User{Mail: claims.Mail, Level: claims.Level})) == 0 {
		return nil
	}
	return claims
}

func CheckPower(c echo.Context) error {
	return c.JSON(http.StatusOK, CheckAuth(c))
}

func Accessible(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>ControlCenter</h1>(´・ω・`) 运行正常<br><hr>Ver: 1.9.0")
}
