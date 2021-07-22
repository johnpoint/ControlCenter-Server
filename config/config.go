package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var Config = &ServiceConfig{}

type ServiceConfig struct {
	ConfigFile       string `json:"config_file"`
	HttpServerListen string `json:"http_server_listen"`
	Environment      string `json:"environment"`
}

func (c *ServiceConfig) ReadConfig() error {
	f, err := os.Open(c.ConfigFile)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, c); err != nil {
		return err
	}
	return nil
}
