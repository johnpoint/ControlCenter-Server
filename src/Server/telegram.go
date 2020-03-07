package main

import (
	"net/http"
	"net/url"
)

func pushNotification(servers []Server, status string) {
	for i := 0; i < len(servers); i++ {
		Tconfig := SysConfig{UID: servers[i].UID, Name: "TELEGRAM_BOT_TOKEN"}
		Iconfig := SysConfig{UID: servers[i].UID, Name: "TELEGRAM_CHAT_ID"}
		Econfig := SysConfig{UID: servers[i].UID, Name: "TELEGRAM_NOTIFICATION"}
		Tdata := getConfig(Tconfig)
		Idata := getConfig(Iconfig)
		Edata := getConfig(Econfig)
		if len(Tdata) == 0 || len(Idata) == 0 || len(Edata) == 0 {
			continue
		}
		botToken := Tdata[0].Value
		userID := Idata[0].Value
		if Edata[0].Value != "true" {
			continue
		}
		name := servers[i].Hostname
		ip := servers[i].Ipv4
		text := "[WARN] " + name + "(" + ip + ") " + status + ""
		text = url.QueryEscape(text)
		http.Get("https://api.telegram.org/bot" + botToken + "/sendMessage?chat_id=" + userID + "&text=" + text)
	}
}
