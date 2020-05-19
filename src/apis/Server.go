package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"main/src/model"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func setupServer(c echo.Context) error {
	token := c.Param("token")
	checkU := getUser(model.User{Token: token})
	if len(checkU) != 1 {
		return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "User Not Found"})
	}
	server := model.Server{UID: checkU[0].ID}
	if err := c.Bind(&server); err != nil {
		log.Print(err)
	}
	check := getServer(model.Server{Ipv4: server.Ipv4})
	if len(check) != 0 {
		return c.JSON(http.StatusBadRequest, model.Callback{Code: 0, Info: "Server already exists"})
	}
	NowTime := time.Now().Unix()
	data := []byte(strconv.FormatInt(NowTime, 10))
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	server.Token = md5str1
	if !addServer(server) {
		return c.JSON(http.StatusBadGateway, model.Callback{Code: 0, Info: "ERROR"})
	}
	addLog("Server", "setupServer:{server:{ip:'"+server.Ipv4+"',token: '"+server.Token+"'}}", 1)
	return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: md5str1})
}

func getServerUpdate(c echo.Context) error {
	token := c.Param("token")
	if (len(getServer(model.Server{Token: token})) == 0) {
		return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "Unauthorized"})
	}
	data := model.UpdateInfo{}
	check := getServer(model.Server{Token: token})
	if len(check) == 1 {
		getCerID := getLinkCer(model.ServerLink{ServerID: check[0].ID})
		if len(getCerID) != 0 {
			CerData := []model.DataCertificate{}
			SiteData := []model.DataSite{}
			for i := 0; i < len(getCerID); i++ {
				if getCerID[i].SiteID != 0 {
					site := getSite(model.Site{ID: getCerID[i].SiteID})[0]
					SiteData = append(SiteData, model.DataSite{ID: site.ID, Config: site.Config, Domain: site.Name, CerID: site.Cer})
				} else if getCerID[i].CertificateID != 0 {
					cer := getCer(model.Certificate{ID: getCerID[i].CertificateID})[0]
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

func getServerInfo(c echo.Context) error {
	user := checkAuth(c)
	server := model.Server{}
	if err := c.Bind(&server); err != nil {
		log.Print(err)
	}
	if user.Level <= 1 {
		server.UID = getUser(model.User{Mail: user.Mail})[0].ID
		return c.JSON(http.StatusOK, getServer(server))
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func getCertificateLinked(c echo.Context) error {
	user := checkAuth(c)
	server := model.Server{}
	if err := c.Bind(&server); err != nil {
		log.Print(err)
	}
	if user.Level <= 1 {
		server.UID = getUser(model.User{Mail: user.Mail})[0].ID
		data := getServer(server)
		if len(data) != 0 {
			return c.JSON(http.StatusOK, getLinkCer(model.ServerLink{ServerID: data[0].ID}))
		}
		return c.JSON(http.StatusOK, getLinkCer(model.ServerLink{ServerID: 0}))
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func getSiteLinked(c echo.Context) error {
	user := checkAuth(c)
	server := model.Server{}
	if err := c.Bind(&server); err != nil {
		log.Print(err)
	}
	if user.Level <= 1 {
		server.UID = getUser(model.User{Mail: user.Mail})[0].ID
		data := getServer(server)
		if len(data) != 0 {
			return c.JSON(http.StatusOK, getLinkSite(model.ServerLink{ServerID: data[0].ID}))
		}
		return c.JSON(http.StatusOK, getLinkSite(model.ServerLink{ServerID: 0}))
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func updateServerInfo(c echo.Context) error {
	user := checkAuth(c)
	if user.Level <= 1 {
		users := getUser(model.User{Mail: user.Mail})
		server := model.Server{}
		if err := c.Bind(&server); err != nil {
			log.Print(err)
		}
		if updateServer(model.Server{ID: server.ID, UID: users[0].ID}, server) {
			addLog("Server", "updateServerInfo:{ip:'"+server.Ipv4+"',user: '"+user.Mail+"'}", 1)
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusBadRequest, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func serverUpdate(c echo.Context) error {
	Token := c.Param("Token")
	if (len(getServer(model.Server{Token: Token})) == 0) {
		return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
	}
	server := model.Server{}
	if err := c.Bind(&server); err != nil {
		log.Print(err)
	}
	log.Print(server.Ipv4 + "\t✓")
	status := getServer(model.Server{Ipv4: server.Ipv4, Token: server.Token})[0].Online
	if status == 2 {
		server.Online = 3
	} else if status == 3 {
		server.Online = 3
	} else {
		server.Online = 1
	}
	if updateServer(model.Server{Ipv4: server.Ipv4, Token: server.Token}, server) {
		return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusBadRequest, model.Callback{Code: 0, Info: "ERROR"})
}

func serverGetCertificate(c echo.Context) error {
	Token := c.Param("Token")
	if (len(getServer(model.Server{Token: Token})) == 0) {
		return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
	}
	id := c.Param("id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	return c.JSON(http.StatusOK, getCer(model.Certificate{ID: id64}))
}

func removeServer(c echo.Context) error {
	user := checkAuth(c)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if user.Level <= 1 {
		if delServer(id, getUser(model.User{Mail: user.Mail})[0].ID) {
			addLog("Server", "removeServer: {user:'"+user.Mail+"',server:{id:"+strconv.FormatInt(id, 10)+"}}", 1)
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func checkOnline() {
	updateServer(model.Server{Online: 1}, model.Server{Online: -1})
	time.Sleep(time.Duration(60) * time.Second)
	offlineServer := getServer(model.Server{Online: -1})
	onlineServer := getServer(model.Server{Online: 3})
	pushNotification(offlineServer, " × ")
	updateServer(model.Server{Online: -1}, model.Server{Online: 2})
	pushNotification(onlineServer, " ✓ ")
	updateServer(model.Server{Online: 3}, model.Server{Online: 1})
	return
	// -1 默认 --> 推送
	// 1 在线
	// 2 等待上线
	// 3 上线 --> 推送
}

func getNow(c echo.Context) error {
	token := c.Param("token")
	servers := getServer(model.Server{Token: token})
	if len(servers) != 0 {
		if servers[0].Update == 1 {
			log.Print(servers[0].Ipv4 + "\t↑")
			updateServer(model.Server{Token: token}, model.Server{Update: -1})
			return c.JSON(http.StatusOK, model.Callback{Code: 211, Info: "Update"})
		}
		if servers[0].Update == 2 {
			log.Print(servers[0].Ipv4 + "\t☯")
			updateServer(model.Server{Token: token}, model.Server{Update: -2})
			return c.JSON(http.StatusOK, model.Callback{Code: 210, Info: "Exit"})
		}
		if servers[0].Update == 3 {
			log.Print(servers[0].Ipv4 + "\t↕")
			updateServer(model.Server{Token: token}, model.Server{Update: -3})
			return c.JSON(http.StatusOK, model.Callback{Code: 212, Info: "Sync"})
		}
		return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}