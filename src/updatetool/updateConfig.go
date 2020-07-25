package main

import (
	"encoding/json"
	"github.com/johnpoint/ControlCenter-Server/src/config"
	"io"
	"os"
)

func main() {
	conf := config.LoadConfig()
	file, _ := os.Create("config.json")
	defer file.Close()
	databy, _ := json.Marshal(conf)
	io.WriteString(file, string(databy))
}
