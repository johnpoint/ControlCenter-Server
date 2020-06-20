package push

import (
	"github.com/johnpoint/ControlCenter-Server/src/model"
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
