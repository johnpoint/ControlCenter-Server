package config

import (
	"ControlCenter-Server/app/model"
	"encoding/json"
	"fmt"
	"os"
)

var Cfg model.Config

func LoadConfig(cfgFile string) error {
	file, _ := os.Open(cfgFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Cfg)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

func InitServer() {
	return
}
