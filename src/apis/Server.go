package apis

import (
	"crypto/md5"
	"fmt"
	. "github.com/johnpoint/ControlCenter-Server/src/database"
	"github.com/johnpoint/ControlCenter-Server/src/model"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func SetupServer(c echo.Context) error {
	token := c.Param("token")
	checkU := GetUser(model.User{Token: token})
	if len(checkU) != 1 {
		return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "User Not Found"})
	}
	server := model.Server{UID: checkU[0].ID}
	if err := c.Bind(&server); err != nil {
		log.Print(err)
	}
	check := GetServer(model.Server{Ipv4: server.Ipv4})
	if len(check) != 0 {
		return c.JSON(http.StatusBadRequest, model.Callback{Code: 0, Info: "Server already exists"})
	}
	NowTime := time.Now().Unix()
	data := []byte(strconv.FormatInt(NowTime, 10))
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	server.Token = md5str1
	if !AddServer(server) {
		return c.JSON(http.StatusBadGateway, model.Callback{Code: 0, Info: "ERROR"})
	}
	AddLog("Server", "setupServer:{server:{ip:'"+server.Ipv4+"',token: '"+server.Token+"'}}", 1)
	return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: md5str1})
}

func GetServerUpdate(c echo.Context) error {
	token := c.Param("token")
	if (len(GetServer(model.Server{Token: token})) == 0) {
		return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "Unauthorized"})
	}
	data := model.UpdateInfo{}
	check := GetServer(model.Server{Token: token})
	if len(check) == 1 {
		getCerID := GetLinkCer(model.ServerLink{ServerID: check[0].ID})
		if len(getCerID) != 0 {
			CerData := []model.DataCertificate{}
			SiteData := []model.DataSite{}
			for i := 0; i < len(getCerID); i++ {
				if getCerID[i].SiteID != 0 {
					site := GetSite(model.Site{ID: getCerID[i].SiteID})[0]
					SiteData = append(SiteData, model.DataSite{ID: site.ID, Config: site.Config, Domain: site.Name, CerID: site.Cer})
				} else if getCerID[i].CertificateID != 0 {
					cer := GetCer(model.Certificate{ID: getCerID[i].CertificateID})[0]
					CerData = append(CerData, model.DataCertificate{ID: cer.ID, Domain: cer.Name, FullChain: cer.Fullchain, Key: cer.Key})
				}
			}
			data.Code = 200
			data.Certificates = CerData
			data.Sites = SiteData
		}
	}
	return c.JSON(http.StatusOK, data)
}

func GetServerInfo(c echo.Context) error {
	user := CheckAuth(c)
	server := model.Server{}
	if err := c.Bind(&server); err != nil {
		log.Print(err)
	}
	if user.Level <= 1 {
		server.UID = GetUser(model.User{Mail: user.Mail})[0].ID
		return c.JSON(http.StatusOK, GetServer(server))
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func GetCertificateLinked(c echo.Context) error {
	user := CheckAuth(c)
	server := model.Server{}
	if err := c.Bind(&server); err != nil {
		log.Print(err)
	}
	if user.Level <= 1 {
		server.UID = GetUser(model.User{Mail: user.Mail})[0].ID
		data := GetServer(server)
		if len(data) != 0 {
			return c.JSON(http.StatusOK, GetLinkCer(model.ServerLink{ServerID: data[0].ID}))
		}
		return c.JSON(http.StatusOK, GetLinkCer(model.ServerLink{ServerID: 0}))
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func GetSiteLinked(c echo.Context) error {
	user := CheckAuth(c)
	server := model.Server{}
	if err := c.Bind(&server); err != nil {
		log.Print(err)
	}
	if user.Level <= 1 {
		server.UID = GetUser(model.User{Mail: user.Mail})[0].ID
		data := GetServer(server)
		if len(data) != 0 {
			return c.JSON(http.StatusOK, GetLinkSite(model.ServerLink{ServerID: data[0].ID}))
		}
		return c.JSON(http.StatusOK, GetLinkSite(model.ServerLink{ServerID: 0}))
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func UpdateServerInfo(c echo.Context) error {
	user := CheckAuth(c)
	if user.Level <= 1 {
		users := GetUser(model.User{Mail: user.Mail})
		server := model.Server{}
		if err := c.Bind(&server); err != nil {
			log.Print(err)
		}
		if UpdateServer(model.Server{ID: server.ID, UID: users[0].ID}, server) {
			AddLog("Server", "updateServerInfo:{ip:'"+server.Ipv4+"',user: '"+user.Mail+"'}", 1)
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusBadRequest, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func ServerUpdate(c echo.Context) error {
	Token := c.Param("Token")
	if (len(GetServer(model.Server{Token: Token})) == 0) {
		return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
	}
	server := model.Server{}
	if err := c.Bind(&server); err != nil {
		log.Print(err)
	}
	log.Print(server.Ipv4 + "\t✓")
	status := GetServer(model.Server{Ipv4: server.Ipv4, Token: server.Token})[0].Online
	if status == 2 {
		server.Online = 3
	} else if status == 3 {
		server.Online = 3
	} else {
		server.Online = 1
	}
	if UpdateServer(model.Server{Ipv4: server.Ipv4, Token: server.Token}, server) {
		return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusBadRequest, model.Callback{Code: 0, Info: "ERROR"})
}

func ServerGetCertificate(c echo.Context) error {
	Token := c.Param("Token")
	if (len(GetServer(model.Server{Token: Token})) == 0) {
		return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
	}
	id := c.Param("id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	return c.JSON(http.StatusOK, GetCer(model.Certificate{ID: id64}))
}

func RemoveServer(c echo.Context) error {
	user := CheckAuth(c)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if user.Level <= 1 {
		if DelServer(id, GetUser(model.User{Mail: user.Mail})[0].ID) {
			AddLog("Server", "removeServer: {user:'"+user.Mail+"',server:{id:"+strconv.FormatInt(id, 10)+"}}", 1)
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func CheckOnline() {
	UpdateServer(model.Server{Online: 1}, model.Server{Online: -1})
	time.Sleep(time.Duration(60) * time.Second)
	offlineServer := GetServer(model.Server{Online: -1})
	onlineServer := GetServer(model.Server{Online: 3})
	pushNotification(offlineServer, " × ")
	UpdateServer(model.Server{Online: -1}, model.Server{Online: 2})
	pushNotification(onlineServer, " ✓ ")
	UpdateServer(model.Server{Online: 3}, model.Server{Online: 1})
	return
	// -1 默认 --> 推送
	// 1 在线
	// 2 等待上线
	// 3 上线 --> 推送
}

func GetNow(c echo.Context) error {
	token := c.Param("token")
	servers := GetServer(model.Server{Token: token})
	if len(servers) != 0 {
		if servers[0].Update == 1 {
			log.Print(servers[0].Ipv4 + "\t↑")
			UpdateServer(model.Server{Token: token}, model.Server{Update: -1})
			return c.JSON(http.StatusOK, model.Callback{Code: 211, Info: "Update"})
		}
		if servers[0].Update == 2 {
			log.Print(servers[0].Ipv4 + "\t☯")
			UpdateServer(model.Server{Token: token}, model.Server{Update: -2})
			return c.JSON(http.StatusOK, model.Callback{Code: 210, Info: "Exit"})
		}
		if servers[0].Update == 3 {
			log.Print(servers[0].Ipv4 + "\t↕")
			UpdateServer(model.Server{Token: token}, model.Server{Update: -3})
			return c.JSON(http.StatusOK, model.Callback{Code: 212, Info: "Sync"})
		}
		return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}
