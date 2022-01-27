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
	configPath  string `yaml:"-"`
	ServiceName string `yaml:"service_name"`

	// 服务端配置
	HttpServerListen string                 `yaml:"http_server_listen"` // API 监听
	TcpServerListen  string                 `yaml:"tcp_server_listen"`  // tcp 服务器监听
	Environment      string                 `yaml:"environment"`        //
	MongoDBConfig    *MongoDBConfig         `yaml:"mongo_db_config"`    // mongo 配置
	RedisConfig      *RedisConfig           `yaml:"redis_config"`       // redis 配置
	TaskQueue        *rabbitmq.Config       `yaml:"task_producer"`      // 任务队列
	Session          *session.SessionConfig `yaml:"session"`            // session 配置
	GrpcConfigMap    map[string]*GrpcConfig `yaml:"grpc_config_map"`    // grpc 配置
	Salt             string                 `yaml:"salt"`               // 加密盐
	AesKey           string                 `yaml:"aes_key"`            // Aes key
	URL              string                 `yaml:"url"`                // 服务提供网址
	InfluxDB         *influxDB.Config       `yaml:"influx_db"`          // 时序数据库
	PerformanceMQ    *rabbitmq.Config       `yaml:"performance_mq"`     // 性能采集队列
	TcpServerMQ      *rabbitmq.Config       `yaml:"tcp_server_mq"`      // tcp 服务器消息队列

	// 客户端 agent 配置
	ServerID             string        `yaml:"server_id"`               // 服务器ID
	Token                string        `yaml:"token"`                   // 服务器Token
	RemoteAddress        string        `yaml:"remote_address"`          // 远端服务器地址
	CollectionInterval   int64         `yaml:"collection_interval"`     // 性能采集间隔时间(秒)
	HeartBeatDuration    time.Duration `yaml:"heart_beat_duration"`     // 心跳间隔
	HeartBeatFixInterval int           `yaml:"heart_beat_fix_interval"` // 心跳修正次数间隔
}

type GrpcConfig struct {
	ServerListen  string `yaml:"server_listen"`
	ClientAddress string `yaml:"client_address"`
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
	URL      string `yaml:"url"`
	Database string `yaml:"database"`
}

type RedisConfig struct {
	Network            string        `yaml:"network"`
	Addr               string        `yaml:"addr"`
	Username           string        `yaml:"username"`
	Password           string        `yaml:"password"`
	DB                 int           `yaml:"db"`
	MaxRetries         int           `yaml:"max_retries"`
	MinRetryBackoff    time.Duration `yaml:"min_retry_backoff"`
	MaxRetryBackoff    time.Duration `yaml:"max_retry_backoff"`
	DialTimeout        time.Duration `yaml:"dial_timeout"`
	ReadTimeout        time.Duration `yaml:"read_timeout"`
	WriteTimeout       time.Duration `yaml:"write_timeout"`
	PoolFIFO           bool          `yaml:"pool_fifo"`
	PoolSize           int           `yaml:"pool_size"`
	MinIdleConns       int           `yaml:"min_idle_conns"`
	MaxConnAge         time.Duration `yaml:"max_conn_age"`
	PoolTimeout        time.Duration `yaml:"pool_timeout"`
	IdleTimeout        time.Duration `yaml:"idle_timeout"`
	IdleCheckFrequency time.Duration `yaml:"idle_check_frequency"`
	readOnly           bool          `yaml:"read_only"`
}
