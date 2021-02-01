package apis

import (
	"ControlCenter-Server/src/database"
	"ControlCenter-Server/src/model"
	"encoding/json"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
	"net/http"
	"strconv"
	"strings"
)

func PushToken(c echo.Context) error {
	u := CheckAuth(c)
	if u.Level <= 1 {
		return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
	}
	return c.JSON(http.StatusOK, model.Callback{Code: 200, Info: "OK"})
}

func APIv2(c echo.Context) error {
	flag := 0
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		msg := ""
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			ws.Close()
			c.Logger().Error(err)
		}
		u := model.User{Token: msg}
		users := database.GetUser(u)
		if len(users) != 1 {
			flag = 1
			ws.Close()
			return
		}
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
				returnData := ""
				if msg == "serverList" {
					list := GetServerList(users[0], 0)
					data, err := json.Marshal(list)
					if err != nil {
						returnData = "[]"
					} else {
						returnData = string(data)
					}
				} else if msg == "overView" {
					list := GetOverView(users[0])
					data, err := json.Marshal(list)
					if err != nil {
						returnData = "[]"
					} else {
						returnData = string(data)
					}
				} else if strings.Contains(msg, "serverStatus") {
					msgSlice := strings.Split(msg, "#")
					if len(msgSlice) != 2 {
						returnData = "{}"
					} else {
						id, err := strconv.ParseInt(msgSlice[1], 10, 64)
						if err == nil {
							list := GetServerStatus(users[0], id)
							data, err := json.Marshal(list)
							if err != nil {
								returnData = "[]"
							} else {
								returnData = "serverStatus" + string(data)
							}
						}
					}
				}
				err = websocket.Message.Send(ws, returnData)
				if err != nil {
					c.Logger().Error(err)
				}
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func GetServerStatus(u model.User, id int64) interface{} {
	type data struct {
		Server        model.Server
		Tasks         []model.Event
		Configuration []struct {
			ID     int64
			Type   string
			Name   string
			Path   string
			Enable bool
		}
		Certificate []struct {
			ID        int64
			Issuer    string
			DNSNames  string
			NotBefore int64
			NotAfter  int64
			Enable    bool
		}
	}
	d := data{}
	servers := GetServerList(u, id)
	if len(servers) != 0 {
		d.Server = servers[0]
	}
	d.Tasks = GetServerEventsV2(u, 0)
	confList := GetConfiguration(u, 0)
	linked := database.GetServerLinkedItem(model.ServerLink{ServerID: id})
	for _, i := range confList {
		enable := false
		for _, j := range linked {
			if i.ID == j.ItemID && j.Type == "File" {
				enable = true
			}
		}
		d.Configuration = append(d.Configuration, struct {
			ID     int64
			Type   string
			Name   string
			Path   string
			Enable bool
		}{ID: i.ID, Type: i.Type, Name: i.Name, Path: i.Path, Enable: enable})
	}
	cerList := database.GetCer(model.Certificate{UID: u.ID})
	for _, i := range cerList {
		enable := false
		for _, j := range linked {
			if i.ID == j.ItemID && j.Type == "Certificate" {
				enable = true
			}
		}
		d.Certificate = append(d.Certificate, struct {
			ID        int64
			Issuer    string
			DNSNames  string
			NotBefore int64
			NotAfter  int64
			Enable    bool
		}{ID: i.ID, Issuer: i.Issuer, DNSNames: i.DNSNames, NotBefore: i.NotBefore, NotAfter: i.NotAfter, Enable: enable})
	}
	return d
}

func GetOverView(user model.User) interface{} {
	type overView struct {
		Server        int
		Certificate   int
		Configuration int
	}
	i := overView{}
	if user.Level <= 1 {
		server := model.Server{UID: user.ID}
		i.Server = len(database.GetServer(server))
		cer := model.Certificate{UID: user.ID}
		i.Certificate = len(database.GetCer(cer))
		conf := model.Configuration{UID: user.ID}
		i.Configuration = len(database.GetConfiguration(conf))
	}
	return i
}

func GetServerList(u model.User, id int64) []model.Server {
	if u.Level <= 1 {
		server := model.Server{UID: u.ID, ID: id}
		return database.GetServer(server)
	}
	return []model.Server{}
}

func GetServerEventsV2(u model.User, id int64) []model.Event {
	if u.Level <= 1 {
		return database.GetEvent(0, id, 0, "", 0)
	}
	return []model.Event{}
}

func GetConfiguration(u model.User, id int64) []model.Configuration {
	conf := model.Configuration{UID: u.ID, ID: id}
	if u.Level <= 1 {
		return database.GetConfiguration(conf)
	}
	return []model.Configuration{}
}
