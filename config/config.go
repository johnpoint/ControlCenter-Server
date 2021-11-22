package config

import (
	"ControlCenter/pkg/rabbitmq"
	"github.com/spf13/viper"
	"time"
)

var Config = &ServiceConfig{}

type ServiceConfig struct {
	ConfigFile       string           `json:"config_file"`
	HttpServerListen string           `json:"http_server_listen"`
	TcpServerListen  string           `json:"tcp_server_listen"`
	Environment      string           `json:"environment"`
	MongoDBConfig    *MongoDBConfig   `json:"mongo_db_config"`
	RedisConfig      *RedisConfig     `json:"redis_config"`
	TaskQueue        *rabbitmq.Config `json:"task_producer"`
	GrpcClientServer string           `json:"grpc_client_server"`
}

func (c *ServiceConfig) ReadConfig() error {
	viper.SetConfigFile(c.ConfigFile)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(Config)
	if err != nil {
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
