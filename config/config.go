package config

import (
	"ControlCenter/pkg/apiMiddleware/session"
	"ControlCenter/pkg/influxDB"
	"ControlCenter/pkg/rabbitmq"
	jsoniter "github.com/json-iterator/go"
	"io"
	"os"
	"time"
)

var Config = new(ServiceConfig)

type ServiceConfig struct {
	configPath string `json:"-"`

	// 服务端配置
	HttpServerListen string                 `json:"http_server_listen"` // API 监听
	TcpServerListen  string                 `json:"tcp_server_listen"`  // tcp 服务器监听
	Environment      string                 `json:"environment"`        //
	MongoDBConfig    *MongoDBConfig         `json:"mongo_db_config"`    // mongo 配置
	RedisConfig      *RedisConfig           `json:"redis_config"`       // redis 配置
	TaskQueue        *rabbitmq.Config       `json:"task_producer"`      // 任务队列
	Session          *session.SessionConfig `json:"session"`            // session 配置
	GrpcConfigMap    map[string]*GrpcConfig `json:"grpc_config_map"`    // grpc 配置
	Salt             string                 `json:"salt"`               // 加密盐
	AesKey           string                 `json:"aes_key"`            // Aes key
	URL              string                 `json:"url"`                // 服务提供网址
	InfluxDB         *influxDB.Config       `json:"influx_db"`          // 时序数据库
	PerformanceMQ    *rabbitmq.Config       `json:"performance_mq"`     // 性能采集队列
	TcpServerMQ      *rabbitmq.Config       `json:"tcp_server_mq"`      // tcp 服务器消息队列

	// 客户端 agent 配置
	ServerID             string        `json:"server_id"`               // 服务器ID
	Token                string        `json:"token"`                   // 服务器Token
	RemoteAddress        string        `json:"remote_address"`          // 远端服务器地址
	CollectionInterval   int64         `json:"collection_interval"`     // 性能采集间隔时间(秒)
	HeartBeatDuration    time.Duration `json:"heart_beat_duration"`     // 心跳间隔
	HeartBeatFixInterval int           `json:"heart_beat_fix_interval"` // 心跳修正次数间隔
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
