package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var Config = &ServiceConfig{}

type ServiceConfig struct {
	ConfigFile       string        `json:"config_file"`
	HttpServerListen string        `json:"http_server_listen"`
	Environment      string        `json:"environment"`
	MongoDBConfig    MongoDBConfig `json:"mongo_db_config"`
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

type MongoDBConfig struct {
	URL      string `json:"url"`
	Database string `json:"database"`
}
