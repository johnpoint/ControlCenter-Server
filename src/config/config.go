package config

import (
	"ControlCenter-Server/src/model"
	"encoding/json"
	"fmt"
	"os"
)

func LoadConfig() model.Config {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := model.Config{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return conf
}

func InitServer() {
	return
}
