package push

import (
	"ControlCenter-Server/app/model"
)

func PushNotification(servers []model.Server, status int64) bool {
	if status == -1 {
		if !Telegram_Push(servers, " × ") {
			return false
		}
	} else {
		if !Telegram_Push(servers, " ✓ ") {
			return false
		}
	}
	return true
}
