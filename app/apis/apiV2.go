package apis

import (
	"ControlCenter-Server/app/database"
	"ControlCenter-Server/app/model"
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
			ws.Close()
			return
		}
		err = websocket.Message.Send(ws, "Certified, Welcome")
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
			returnData := parseCommand(msg, users[0])
			err = websocket.Message.Send(ws, returnData)
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func parseCommand(msg string, u model.User) string {
	returnData := "[]"
	msgSlice := strings.Split(msg, "@")
	if len(msgSlice) < 2 {
		returnData = "[]"
		return returnData
	}
	if msgSlice[0] == "get" {
		msg = msgSlice[1]
		if msg == "serverList" {
			list := GetServerList(u, 0)
			data, err := json.Marshal(list)
			if err != nil {
				returnData = "[]"
			} else {
				returnData = string(data)
			}
		} else if msg == "overView" {
			list := GetOverView(u)
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
					list := GetServerStatus(u, id)
					data, err := json.Marshal(list)
					if err != nil {
						returnData = "[]"
					} else {
						returnData = "serverStatus" + string(data)
					}
				}
			}
		} else if msg == "ConfigurationList" {
			list := GetConfigurationList(u)
			data, err := json.Marshal(list)
			if err != nil {
				returnData = "[]"
			} else {
				returnData = string(data)
			}
		} else if msg == "CertificateList" {
			list := GetCertificateList(u)
			data, err := json.Marshal(list)
			if err != nil {
				returnData = "[]"
			} else {
				returnData = string(data)
			}
		}
	}

	return returnData
}

func GetCertificateList(u model.User) interface{} {
	type data struct {
		ID        int64
		Issuer    string
		DNSNames  string
		NotBefore int64
		NotAfter  int64
	}
	d := []data{}
	cerList := database.GetCer(model.Certificate{UID: u.ID})
	for _, i := range cerList {
		d = append(d, struct {
			ID        int64
			Issuer    string
			DNSNames  string
			NotBefore int64
			NotAfter  int64
		}{ID: i.ID, Issuer: i.Issuer, DNSNames: i.DNSNames, NotBefore: i.NotBefore, NotAfter: i.NotAfter})
	}
	return d
}

func GetConfigurationList(u model.User) interface{} {
	type data struct {
		ID   int64
		Type string
		Name string
		Path string
	}
	d := []data{}
	list := database.GetConfiguration(model.Configuration{UID: u.ID})
	for _, i := range list {
		d = append(d, struct {
			ID   int64
			Type string
			Name string
			Path string
		}{ID: i.ID, Type: i.Type, Name: i.Name, Path: i.Path})
	}
	return d
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
