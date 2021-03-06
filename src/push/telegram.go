package push

import (
	"ControlCenter-Server/src/database"
	"ControlCenter-Server/src/model"
	"fmt"
	"net/http"
	"net/url"
)

func Telegram_Push(servers []model.Server, status string) bool {
	for i := 0; i < len(servers); i++ {
		TelegramBotToken := database.GetConfig(model.SysConfig{UID: servers[i].UID, Name: "TELEGRAM_BOT_TOKEN"})
		TelegramChatId := database.GetConfig(model.SysConfig{UID: servers[i].UID, Name: "TELEGRAM_CHAT_ID"})
		TelegramNotification := database.GetConfig(model.SysConfig{UID: servers[i].UID, Name: "TELEGRAM_NOTIFICATION"})
		if len(TelegramBotToken) == 0 || len(TelegramChatId) == 0 || len(TelegramNotification) == 0 {
			continue
		}
		botToken := TelegramBotToken[0].Value
		userID := TelegramChatId[0].Value
		if TelegramNotification[0].Value != "true" {
			continue
		}
		name := servers[i].Hostname
		ip := servers[i].Ipv4
		text := "[Server] " + name + " (" + ip + ") " + status + ""
		text = url.QueryEscape(text)
		_, err := http.Get("https://api.telegram.org/bot" + botToken + "/sendMessage?chat_id=" + userID + "&text=" + text)
		if err != nil {
			fmt.Print(err)
		}
	}
	return true
}
