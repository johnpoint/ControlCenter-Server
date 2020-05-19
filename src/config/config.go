package main

import (
	"encoding/json"
	"fmt"
	"main/src/model"
	"os"
)

func loadConfig() model.Config {
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

func initServer() {
	return
}
