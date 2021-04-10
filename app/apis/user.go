package apis

import (
	"ControlCenter-Server/app/config"
	"ControlCenter-Server/app/database"
	"ControlCenter-Server/app/model"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func GetUserInfo(c echo.Context) error {
	user := CheckAuth(c)
	if user != nil {
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}

func UpdateUserInfo(c echo.Context) error {
	user := CheckAuth(c)
	if user != nil {
		database.AddLog("User", "updateUserInfo:{user:{mail:'"+user.Mail+"'}}", 1)
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}

func ReSetPassword(c echo.Context) error {
	conf := config.Cfg
	salt := conf.Salt
	user := CheckAuth(c)
	Getdata := model.ReSetPassword{}
	if err := c.Bind(&Getdata); err != nil {
		return err
	}
	if user != nil {
		data := []byte(user.Mail + salt + Getdata.Oldpass)
		has := md5.Sum(data)
		oldpass := fmt.Sprintf("%x", has)
		u := database.GetUser(model.User{Mail: user.Mail})
		if u[0].Password != oldpass {
			return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
		}
		data = []byte(user.Mail + salt + Getdata.Newpass)
		has = md5.Sum(data)
		newpass := fmt.Sprintf("%x", has)
		if (database.UpdateUser(u[0], model.User{Password: newpass})) {
			database.AddLog("User", "reSetPassword:{user:{mail:'"+user.Mail+"'}}", 1)
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}

func GetUserToken(c echo.Context) error {
	user := CheckAuth(c)
	data := database.GetUser(model.User{Mail: user.Mail})
	if len(data) != 0 {
		return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: data[0].Token})
	}
	return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "User Not Found"})
}

func GetNewToken(c echo.Context) error {
	user := CheckAuth(c)
	timeUnixNano := time.Now().UnixNano()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(timeUnixNano, 10))
	newToken := fmt.Sprintf("%x", h.Sum(nil))
	if (database.UpdateUser(model.User{Mail: user.Mail}, model.User{Token: newToken})) {
		database.AddLog("User", "getNewToken:{user:{mail:'"+user.Mail+"'}}", 1)
		return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "ERROR"})
}

func ChangeLevel(c echo.Context) error {
	user := CheckAuth(c)
	if user != nil {
		if database.GetUser(model.User{Mail: user.Mail})[0].Level <= 0 {
			uid, _ := strconv.ParseInt(c.Param("uid"), 10, 64)
			level, _ := strconv.ParseInt(c.Param("level"), 10, 64)
			userTarget := database.GetUser(model.User{ID: uid})
			if len(userTarget) != 0 {
				if userTarget[0].Level == 0 {
					return c.JSON(http.StatusForbidden, model.Callback{Code: 0, Info: "No permission"})
				}
			}
			if database.UpdateUser(model.User{ID: uid}, model.User{Level: level}) {
				database.AddLog("User", "changeLevel:{user:{mail:'"+user.Mail+"'},target:{id:"+strconv.FormatInt(uid, 10)+",level:"+strconv.FormatInt(level, 10)+"}}", 1)
				return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
			}
			return c.JSON(http.StatusInternalServerError, model.Callback{Code: 0, Info: "ERROR"})
		}
		return c.JSON(http.StatusForbidden, model.Callback{Code: 0, Info: "No permission"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}

func GetUserList(c echo.Context) error {
	user := CheckAuth(c)
	if user != nil {
		if database.GetUser(model.User{Mail: user.Mail})[0].Level <= 0 {
			users := database.GetUser(model.User{})
			for i := 0; i < len(users); i++ {
				users[i].Password = "*********"
				users[i].Token = "*********"
			}
			return c.JSON(http.StatusOK, users)
		}
		return c.JSON(http.StatusForbidden, model.Callback{Code: 0, Info: "No permission"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}
