package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd    = &cobra.Command{}
	configPath string
	commands   []*cobra.Command
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "config_local.json", "config file (default is ./config_local.json)")
	rootCmd.AddCommand(httpServerCommand)          // API服务
	rootCmd.AddCommand(clientCommand)              // 上报客户端
	rootCmd.AddCommand(genConfigCommand)           // 生成空配置文件
	rootCmd.AddCommand(taskConsumerCommand)        // 下发任务消费者
	rootCmd.AddCommand(tcpServerCommand)           // tcp
	rootCmd.AddCommand(performanceConsumerCommand) // 性能信息消费者
	rootCmd.AddCommand(tcpServerConsumerCommand)   // tcp信息消费者
}

func initConfig() {
	if configPath == "" {
		configPath = "config_local.json"
	}
}
