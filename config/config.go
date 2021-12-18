package config

import (
	"ControlCenter/pkg/apiMiddleware/session"
	"ControlCenter/pkg/rabbitmq"
	jsoniter "github.com/json-iterator/go"
	"io"
	"os"
	"time"
)

var Config = new(ServiceConfig)

type ServiceConfig struct {
	configPath       string
	HttpServerListen string                 `json:"http_server_listen"`
	TcpServerListen  string                 `json:"tcp_server_listen"`
	Environment      string                 `json:"environment"`
	MongoDBConfig    *MongoDBConfig         `json:"mongo_db_config"`
	RedisConfig      *RedisConfig           `json:"redis_config"`
	TaskQueue        *rabbitmq.Config       `json:"task_producer"`
	Session          *session.SessionConfig `json:"session"`
	GrpcConfigMap    map[string]*GrpcConfig `json:"grpc_config_map"`
	Salt             string                 `json:"salt"`
	URL              string                 `json:"url"`
}

type GrpcConfig struct {
	ServerListen  string `json:"server_listen"`
	ClientAddress string `json:"client_address"`
}

func (c *ServiceConfig) SetPath(path string) *ServiceConfig {
	c.configPath = path
	return c
}

func (c *ServiceConfig) ReadConfig() error {
	f, err := os.Open(c.configPath)
	if err != nil {
		return err
	}
	defer f.Close()
	cfgByte, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	if err = jsoniter.Unmarshal(cfgByte, c); err != nil {
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
