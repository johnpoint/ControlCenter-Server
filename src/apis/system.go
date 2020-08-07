package apis

import (
	"encoding/json"
	"github.com/johnpoint/ControlCenter-Server/src/config"
	"github.com/johnpoint/ControlCenter-Server/src/database"
	"github.com/johnpoint/ControlCenter-Server/src/model"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

const VERSION = "2.0.2"

func SysRestart(c echo.Context) error {
	return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: ""})
}

func SetBackupFile(c echo.Context) error {
	user := CheckAuth(c)
	conf := config.LoadConfig()
	if user != nil {
		if user.Level <= 0 {
			file, err := c.FormFile("file")
			if err != nil {
				return err
			}
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			// Destination
			dst, err := os.Create(conf.Database)
			if err != nil {
				return err
			}
			defer dst.Close()

			// Copy
			if _, err = io.Copy(dst, src); err != nil {
				return err
			}
			database.AddLog("System", "setBackupFile:{user:{mail:'"+user.Mail+"'}}", 1)
			return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "OK"})
		} else {
			return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
		}
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}

func GetBackupFile(c echo.Context) error {
	conf := config.LoadConfig()
	token := c.Param("token")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
	}
	getuser := model.User{Token: token}
	userInfo := database.GetUser(getuser)
	if len(userInfo) == 0 {
		re := model.Callback{Code: 0, Info: "account or token incorrect"}
		return c.JSON(http.StatusUnauthorized, re)
	} else {
		if userInfo[0].Level <= 0 {
			database.AddLog("System", "getBackupFile:{user:{mail:'"+userInfo[0].Mail+"'}}", 1)
			return c.File(conf.Database)
		} else {
			return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
		}
	}
}

func SetSetting(c echo.Context) error {
	user := CheckAuth(c)
	name := c.Param("name")
	value := c.Param("value")
	config := model.SysConfig{Name: name, Value: value, UID: database.GetUser(model.User{Mail: user.Mail})[0].ID}
	if database.SetConfig(config) {
		database.AddLog("System", "setSetting:{user:{mail:'"+user.Mail+"'},settings:{name:'"+name+"',value:'"+value+"'}}", 1)
		return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "ERROR"})
}

func GetSetting(c echo.Context) error {
	user := CheckAuth(c)
	name := c.Param("name")
	config := model.SysConfig{Name: name, UID: database.GetUser(model.User{Mail: user.Mail})[0].ID}
	return c.JSON(http.StatusOK, database.GetConfig(config))
}

func Accessible(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>ControlCenter</h1>(´・ω・`) 运行正常<br><hr>Ver: "+VERSION)
}

func GetSystemInfo(c echo.Context) error {
	sysinfo := getSystemInfo()
	return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: sysinfo})
}

func getSystemInfo() string {
	v, _ := mem.VirtualMemory()
	s, _ := mem.SwapMemory()
	cc, _ := cpu.Percent(time.Second, false)
	d, _ := disk.Usage("/")
	n, _ := host.Info()
	l, _ := load.Avg()
	ss := new(model.StatusServer)
	ss.Load = l
	ss.Uptime = n.Uptime
	ss.BootTime = n.BootTime
	ss.Percent.Mem = v.UsedPercent
	ss.Percent.CPU = cc[0]
	ss.Percent.Swap = s.UsedPercent
	ss.Percent.Disk = d.UsedPercent

	ss.Version = VERSION
	b, err := json.Marshal(ss)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}
