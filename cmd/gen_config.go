package cmd

import (
	"ControlCenter/config"
	"ControlCenter/pkg/apiMiddleware/session"
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
		}
		jsonByte, _ := json.Marshal(&newConfig)
		fmt.Println(string(jsonByte))
	},
}
