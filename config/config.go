package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

var Config = &ServiceConfig{}

type ServiceConfig struct {
	ConfigFile       string         `json:"config_file"`
	HttpServerListen string         `json:"http_server_listen"`
	Environment      string         `json:"environment"`
	MongoDBConfig    *MongoDBConfig `json:"mongo_db_config"`
	RedisConfig      *RedisConfig   `json:"redis_config"`
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

type RedisConfig struct {
	Network            string        `json:"network"`
	Addr               string        `json:"addr"`
	Username           string        `json:"username"`
	Password           string        `json:"password"`
	DB                 int           `json:"db"`
	MaxRetries         int           `json:"max_retries"`
	MinRetryBackoff    time.Duration `json:"min_retry_backoff"`
	MaxRetryBackoff    time.Duration `json:"max_retry_backoff"`
	DialTimeout        time.Duration `json:"dial_timeout"`
	ReadTimeout        time.Duration `json:"read_timeout"`
	WriteTimeout       time.Duration `json:"write_timeout"`
	PoolFIFO           bool          `json:"pool_fifo"`
	PoolSize           int           `json:"pool_size"`
	MinIdleConns       int           `json:"min_idle_conns"`
	MaxConnAge         time.Duration `json:"max_conn_age"`
	PoolTimeout        time.Duration `json:"pool_timeout"`
	IdleTimeout        time.Duration `json:"idle_timeout"`
	IdleCheckFrequency time.Duration `json:"idle_check_frequency"`
	readOnly           bool          `json:"read_only"`
}
