package config

import (
	"encoding/json"
	"io"
	"os"
)

func UpdateConfig() {
	conf := LoadConfig()
	file, _ := os.Create("config.json")
	defer file.Close()
	conf.RedisConfig.Enable = false
	databy, _ := json.Marshal(conf)
	io.WriteString(file, string(databy))
}
