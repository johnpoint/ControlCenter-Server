package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadConfig() Config {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Config{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return conf
}

func initServer() {
	return
}
