package cmd

import (
	"ControlCenter/config"
	"ControlCenter/pkg/apiMiddleware/session"
	"ControlCenter/pkg/influxDB"
	"ControlCenter/pkg/rabbitmq"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

var genConfigCommand = &cobra.Command{
	Use:   "genConfig",
	Short: "Generate config",
	Run: func(cmd *cobra.Command, args []string) {
		newConfig := config.ServiceConfig{
			MongoDBConfig: &config.MongoDBConfig{},
			RedisConfig:   &config.RedisConfig{},
			TaskQueue:     &rabbitmq.Config{},
			Session:       &session.SessionConfig{},
			PerformanceMQ: &rabbitmq.Config{},
			GrpcConfigMap: make(map[string]*config.GrpcConfig),
			InfluxDB:      &influxDB.Config{},
		}
		jsonByte, _ := json.Marshal(&newConfig)
		fmt.Println(string(jsonByte))
	},
}
