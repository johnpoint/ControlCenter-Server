package apis

import (
	"ControlCenter-Server/src/database"
	"ControlCenter-Server/src/model"
	"ControlCenter-Server/src/push"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

type DataCertificate struct {
	ID        int64
	Domain    string
	FullChain string
	Key       string
}

func SetupServer(c echo.Context) error {
	token := c.Param("token")
	checkU := database.GetUser(model.User{Token: token})
	if len(checkU) != 1 {
		return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "User Not Found"})
	}
	server := model.Server{UID: checkU[0].ID}
	if err := c.Bind(&server); err != nil {
		log.Print(err)
	}
	check := database.GetServer(model.Server{Ipv4: server.Ipv4})
	if len(check) != 0 {
		return c.JSON(http.StatusBadRequest, model.Callback{Code: 0, Info: "Server already exists"})
	}
	NowTime := time.Now().Unix()
	data := []byte(strconv.FormatInt(NowTime, 10))
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	server.Token = md5str1
	if !database.AddServer(server) {
		return c.JSON(http.StatusBadGateway, model.Callback{Code: 0, Info: "ERROR"})
	}
	database.AddLog("Server", "setupServer:{server:{ip:'"+server.Ipv4+"',token: '"+server.Token+"'}}", 1)
	return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: md5str1})
}

func GetServerUpdate(c echo.Context) error {
	token := c.Param("token")
	if (len(database.GetServer(model.Server{Token: token})) == 0) {
		return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "Unauthorized"})
	}
	data := model.UpdateInfo{}
	check := database.GetServer(model.Server{Token: token})
	if len(check) == 1 {
		getCerID := database.GetServerLinkedItem(model.ServerLink{ServerID: check[0].ID, Type: "Certificate"})
		if len(getCerID) != 0 {
			CerData := []model.DataCertificate{}
			for i := 0; i < len(getCerID); i++ {
				if getCerID[i].ItemID != 0 {
					cer := database.GetCer(model.Certificate{ID: getCerID[i].ItemID})[0]
					CerData = append(CerData, model.DataCertificate{ID: cer.ID, Domain: cer.Name, FullChain: cer.Fullchain, Key: cer.Key})
				}
			}
			data.Certificates = CerData
		}
		getConfID := database.GetServerLinkedItem(model.ServerLink{ServerID: check[0].ID, Type: "File"})
		if len(getConfID) != 0 {
			ConfFile := []model.File{}
			for i := 0; i < len(getConfID); i++ {
				if getConfID[i].ItemID != 0 {
					conf := database.GetConfiguration(model.Configuration{ID: getConfID[i].ItemID})[0]
					ConfFile = append(ConfFile, model.File{ID: conf.ID, Name: conf.Name, Value: conf.Value, Path: conf.Path})
				}
			}
			data.ConfFile = ConfFile
		}
		data.Code = 200
	}
	return c.JSON(http.StatusOK, data)
}

func APIv2(c echo.Context) error {
	flag := 0
	token := c.Param("token")
	u := model.User{Token: token}
	users := database.GetUser(u)
	if len(users) != 1 {
		flag = 1
	}
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		if flag != 0 {
			err := websocket.Message.Send(ws, "Unauthorized, Bye")
			if err != nil {
				c.Logger().Error(err)
			}
			ws.Close()
		} else {
			err := websocket.Message.Send(ws, "Certified, Welcome")
			if err != nil {
				c.Logger().Error(err)
			}
			for {
				// Read
				msg := ""
				err = websocket.Message.Receive(ws, &msg)
				if err != nil {
					ws.Close()
					c.Logger().Error(err)
					break
				}

				switch msg {
				case "serverList":
					list := GetServerList(users[0])
					data, err := json.Marshal(list)
					if err != nil {
						err := websocket.Message.Send(ws, "[]")
						if err != nil {
							c.Logger().Error(err)
						}
						break
					}
					err = websocket.Message.Send(ws, string(data))
					if err != nil {
						c.Logger().Error(err)
					}
					break
				}
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func GetServerList(u model.User) []model.Server {
	if u.Level <= 1 {
		server := model.Server{}
		server.UID = u.ID
		return database.GetServer(server)
	}
	return []model.Server{}
}

func GetServerInfo(c echo.Context) error {
	user := CheckAuth(c)
	server := model.Server{}
	if err := c.Bind(&server); err != nil {
		log.Print(err)
	}
	if user.Level <= 1 {
		server.UID = database.GetUser(model.User{Mail: user.Mail})[0].ID
		return c.JSON(http.StatusOK, database.GetServer(server))
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
		server.UID = database.GetUser(model.User{Mail: user.Mail})[0].ID
		data := database.GetServer(server)
		if len(data) != 0 {
			return c.JSON(http.StatusOK, database.GetServerLinkedItem(model.ServerLink{ServerID: data[0].ID, Type: "Certificate"}))
		}
		return c.JSON(http.StatusOK, model.Callback{Code: 404, Info: "Server Not Found"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func GetConfigurationLinked(c echo.Context) error {
	user := CheckAuth(c)
	server := model.Server{}
	if err := c.Bind(&server); err != nil {
		log.Print(err)
	}
	if user.Level <= 1 {
		server.UID = database.GetUser(model.User{Mail: user.Mail})[0].ID
		data := database.GetServer(server)
		if len(data) != 0 {
			return c.JSON(http.StatusOK, database.GetServerLinkedItem(model.ServerLink{ServerID: data[0].ID, Type: "File"}))
		}
		return c.JSON(http.StatusOK, model.Callback{Code: 404, Info: "Server Not Found"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func GetServerEvents(c echo.Context) error {
	user := CheckAuth(c)
	server := model.Server{}
	if err := c.Bind(&server); err != nil {
		log.Print(err)
	}
	if user.Level <= 1 {
		server.UID = database.GetUser(model.User{Mail: user.Mail})[0].ID
		data := database.GetServer(server)
		if len(data) != 0 {
			return c.JSON(http.StatusOK, database.GetEvent(0, data[0].ID, 0, "", 0))
		}
		return c.JSON(http.StatusOK, model.Callback{Code: 404, Info: "Server Not Found"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func UpdateServerInfo(c echo.Context) error {
	user := CheckAuth(c)
	if user.Level <= 1 {
		users := database.GetUser(model.User{Mail: user.Mail})
		server := model.Server{}
		if err := c.Bind(&server); err != nil {
			log.Print(err)
		}
		if database.UpdateServer(model.Server{ID: server.ID, UID: users[0].ID}, server) {
			database.AddLog("Server", "updateServerInfo:{ip:'"+server.Ipv4+"',user: '"+user.Mail+"'}", 1)
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusBadRequest, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func ServerUpdate(c echo.Context) error {
	Token := c.Param("Token")
	if (len(database.GetServer(model.Server{Token: Token})) == 0) {
		return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
	}
	server := model.Server{}
	if err := c.Bind(&server); err != nil {
		log.Print(err)
	}
	log.Print(server.Ipv4 + "\t✓")
	dbServer := database.GetServer(model.Server{Ipv4: server.Ipv4, Token: server.Token})
	if dbServer[0].Online == 2 {
		server.Online = 3
	} else if dbServer[0].Online == 3 {
		server.Online = 3
	} else {
		server.Online = 1
	}
	if database.UpdateServer(model.Server{ID: dbServer[0].ID}, model.Server{Status: server.Status, Online: server.Online}) {
		return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusBadRequest, model.Callback{Code: 0, Info: "ERROR"})
}

func RemoveServer(c echo.Context) error {
	user := CheckAuth(c)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if user.Level <= 1 {
		if database.DelServer(id, database.GetUser(model.User{Mail: user.Mail})[0].ID) {
			database.AddLog("Server", "removeServer: {user:'"+user.Mail+"',server:{id:"+strconv.FormatInt(id, 10)+"}}", 1)
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
		return c.JSON(http.StatusOK, model.Callback{Code: 0, Info: "ERROR"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func CheckOnline() {
	// -1 默认 --> 推送
	// 1 在线
	// 2 等待上线
	// 3 上线 --> 推送
	database.UpdateServer(model.Server{Online: 1}, model.Server{Online: -1})
	time.Sleep(time.Duration(60) * time.Second)
	offlineServer := database.GetServer(model.Server{Online: -1})
	onlineServer := database.GetServer(model.Server{Online: 3})
	go func() {
		if !push.PushNotification(offlineServer, -1) {
			log.Print("PUSH FAIL")
		}
	}()
	database.UpdateServer(model.Server{Online: -1}, model.Server{Online: 2})
	go func() {
		if !push.PushNotification(onlineServer, 1) {
			log.Print("PUSH FAIL")
		}
	}()
	database.UpdateServer(model.Server{Online: 3}, model.Server{Online: 1})
	return
}

func GetNow(c echo.Context) error {
	token := c.Param("token")
	servers := database.GetServer(model.Server{Token: token})
	if len(servers) != 0 {
		eventList := database.GetEvent(0, servers[0].ID, 0, "", 1)
		if len(eventList) != 0 {
			b := database.FinishEvent(eventList[0].ID)
			if b {
				return c.JSON(http.StatusOK, model.Callback{Code: eventList[0].Code, Info: eventList[0].Info})
			}
			b = database.DeleteEvent(eventList[0].ID)
			if b {
				return c.JSON(http.StatusOK, model.Callback{Code: eventList[0].Code, Info: eventList[0].Info})
			}
			return c.JSON(http.StatusOK, model.Callback{Code: 500, Info: "Internal Server Error"})
		}
		return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func ChangeDockerStatus(c echo.Context) error {
	user := CheckAuth(c)
	if user.Level <= 1 {
		action, _ := strconv.ParseInt(c.Param("action"), 10, 64)
		containterID := c.Param("id")
		serverID, _ := strconv.ParseInt(c.Param("serverid"), 10, 64)
		if len(database.GetServer(model.Server{ID: serverID, UID: database.GetUser(model.User{Mail: user.Mail})[0].ID})) == 1 {
			if database.AddEvent(1, serverID, action, containterID) == false {
				log.Print("AddEvent Fail:" + c.Path())
				database.AddLog("Event", c.Path()+"|From:"+c.RealIP(), 2)
				return c.JSON(http.StatusOK, model.Callback{Code: 500, Info: "Internal Server Error"})
			}
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}

func AddClientEvent(c echo.Context) error {
	user := CheckAuth(c)
	if user.Level <= 1 {
		action, _ := strconv.ParseInt(c.Param("action"), 10, 64)
		serverID, _ := strconv.ParseInt(c.Param("serverid"), 10, 64)
		//fmt.Println(database.AddEvent(1, serverID, action, "OK"))
		if len(database.GetServer(model.Server{ID: serverID, UID: database.GetUser(model.User{Mail: user.Mail})[0].ID})) == 1 {
			if database.AddEvent(1, serverID, action, "OK") == false {
				log.Print("AddEvent Fail:" + c.Path())
				database.AddLog("Event", c.Path()+"|From:"+c.RealIP(), 2)
				return c.JSON(http.StatusOK, model.Callback{Code: 500, Info: "Internal Server Error"})
			}
			return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
		}
	}
	return c.JSON(http.StatusUnauthorized, model.Callback{Code: 0, Info: "Unauthorized"})
}
