package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func setupServer(c echo.Context) error {
	token := c.Param("token")
	checkU := getUser(User{Token: token})
	if len(checkU) != 1 {
		return c.JSON(http.StatusOK, Callback{Code: 0, Info: "User Not Found"})
	}
	server := Server{UID: checkU[0].ID}
	if err := c.Bind(&server); err != nil {
		panic(err)
	}
	check := getServer(Server{Ipv4: server.Ipv4})
	if len(check) != 0 {
		return c.JSON(http.StatusBadRequest, Callback{Code: 0, Info: "Server already exists"})
	}
	NowTime := time.Now().Unix()
	data := []byte(strconv.FormatInt(NowTime, 10))
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	server.Token = md5str1
	if !addServer(server) {
		return c.JSON(http.StatusBadGateway, Callback{Code: 0, Info: "ERROR"})
	}
	addLog("Server", "setupServer:{server:{ip:'"+server.Ipv4+"',token: '"+server.Token+"'}}", 1)
	return c.JSON(http.StatusOK, Callback{Code: 200, Info: md5str1})
}

func getServerUpdate(c echo.Context) error {
	token := c.Param("token")
	if (len(getServer(Server{Token: token})) == 0) {
		return c.JSON(http.StatusOK, Callback{Code: 0, Info: "Unauthorized"})
	}
	data := UpdateInfo{}
	check := getServer(Server{Token: token})
	if len(check) == 1 {
		getCerID := getLinkCer(ServerLink{ServerID: check[0].ID})
		if len(getCerID) != 0 {
			CerData := []DataCertificate{}
			SiteData := []DataSite{}
			for i := 0; i < len(getCerID); i++ {
				if getCerID[i].SiteID != 0 {
					site := getSite(Site{ID: getCerID[i].SiteID})[0]
					SiteData = append(SiteData, DataSite{ID: site.ID, Config: site.Config, Domain: site.Name, CerID: site.Cer})
				} else if getCerID[i].CertificateID != 0 {
					cer := getCer(Certificate{ID: getCerID[i].CertificateID})[0]
					CerData = append(CerData, DataCertificate{ID: cer.ID, Domain: cer.Name, FullChain: cer.Fullchain, Key: cer.Key})
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
	server := Server{}
	if err := c.Bind(&server); err != nil {
		panic(err)
	}
	if user.Level <= 1 {
		server.UID = getUser(User{Mail: user.Mail})[0].ID
		return c.JSON(http.StatusOK, getServer(server))
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func getCertificateLinked(c echo.Context) error {
	user := checkAuth(c)
	server := Server{}
	if err := c.Bind(&server); err != nil {
		panic(err)
	}
	if user.Level <= 1 {
		server.UID = getUser(User{Mail: user.Mail})[0].ID
		data := getServer(server)
		if len(data) != 0 {
			return c.JSON(http.StatusOK, getLinkCer(ServerLink{ServerID: data[0].ID}))
		}
		return c.JSON(http.StatusOK, getLinkCer(ServerLink{ServerID: 0}))
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func getSiteLinked(c echo.Context) error {
	user := checkAuth(c)
	server := Server{}
	if err := c.Bind(&server); err != nil {
		panic(err)
	}
	if user.Level <= 1 {
		server.UID = getUser(User{Mail: user.Mail})[0].ID
		data := getServer(server)
		if len(data) != 0 {
			return c.JSON(http.StatusOK, getLinkSite(ServerLink{ServerID: data[0].ID}))
		}
		return c.JSON(http.StatusOK, getLinkSite(ServerLink{ServerID: 0}))
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func updateServerInfo(c echo.Context) error {
	user := checkAuth(c)
	if user.Level <= 1 {
		server := Server{}
		if err := c.Bind(&server); err != nil {
			panic(err)
		}
		if updateServer(Server{Ipv4: server.Ipv4, Token: server.Token}, server) {
			addLog("Server", "updateServerInfo:{ip:'"+server.Ipv4+"',user: '"+user.Mail+"'}", 1)
			return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusBadRequest, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func serverUpdate(c echo.Context) error {
	Token := c.Param("Token")
	if (len(getServer(Server{Token: Token})) == 0) {
		return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
	}
	server := Server{}
	if err := c.Bind(&server); err != nil {
		panic(err)
	}
	fmt.Println(":: Get Server update From :" + server.Ipv4)
	status := getServer(Server{Ipv4: server.Ipv4, Token: server.Token})[0].Online
	if status == 2 {
		server.Online = 3
	} else if status == 3 {
		server.Online = 3
	} else {
		server.Online = 1
	}
	if updateServer(Server{Ipv4: server.Ipv4, Token: server.Token}, server) {
		return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusBadRequest, Callback{Code: 0, Info: "ERROR"})
}

func serverGetCertificate(c echo.Context) error {
	Token := c.Param("Token")
	if (len(getServer(Server{Token: Token})) == 0) {
		return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
	}
	id := c.Param("id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	return c.JSON(http.StatusOK, getCer(Certificate{ID: id64}))
}

func removeServer(c echo.Context) error {
	user := checkAuth(c)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if user.Level <= 1 {
		if delServer(id, getUser(User{Mail: user.Mail})[0].ID) {
			addLog("Server", "removeServer: {user:'"+user.Mail+"',server:{id:"+strconv.FormatInt(id, 10)+"}}", 1)
			return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusOK, Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}

func checkOnline() {

	updateServer(Server{Online: 1}, Server{Online: -1})
	time.Sleep(time.Duration(120) * time.Second)
	offlineServer := getServer(Server{Online: -1})
	onlineServer := getServer(Server{Online: 3})
	pushNotification(offlineServer, "offline")
	updateServer(Server{Online: -1}, Server{Online: 2})
	pushNotification(onlineServer, "online")
	updateServer(Server{Online: 3}, Server{Online: 1})
	return
	// -1 默认 --> 推送
	// 1 在线
	// 2 等待上线
	// 3 上线 --> 推送
}

func getNow(c echo.Context) error {
	token := c.Param("token")
	servers := getServer(Server{Token: token})
	if len(servers) != 0 {
		if servers[0].Update == 1 {
			fmt.Println(servers[0].Ipv4)
			updateServer(Server{Token: token}, Server{Update: -1})
			return c.JSON(http.StatusOK, Callback{Code: 211, Info: "Update"})
		}
		return c.JSON(http.StatusOK, Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusUnauthorized, Callback{Code: 0, Info: "Unauthorized"})
}
