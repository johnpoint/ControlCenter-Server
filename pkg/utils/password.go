package utils

import (
	"ControlCenter/config"
	"fmt"
)

func EncodePassword(password string) string {
	hashedPwd := Md5(fmt.Sprintf("%s%s%s", config.Config.Salt, password, config.Config.Salt))
	encrypt, err := AesEncrypt([]byte(hashedPwd), []byte(config.Config.AesKey))
	if err != nil {
		return ""
	}
	return string(encrypt)
}

func EqualPassword(password, standardPwd string) bool {
	return standardPwd == EncodePassword(password)
}
