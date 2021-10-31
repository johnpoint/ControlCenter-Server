package cmd

import (
	"ControlCenter/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd  = &cobra.Command{}
	cfgFile  string
	commands []*cobra.Command
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config_local.json", "config file (default is ./config_local.json)")
	rootCmd.AddCommand(httpServerCommand)   // API服务
	rootCmd.AddCommand(clientCommand)       // 上报客户端
	rootCmd.AddCommand(genConfigCommand)    // 生成空配置文件
	rootCmd.AddCommand(taskConsumerCommand) // 下发任务消费者
	rootCmd.AddCommand(tcpServerCommand)    // tcp
}

func initConfig() {
	if cfgFile == "" {
		cfgFile = "config_local.json"
	}
	config.Config = &config.ServiceConfig{ConfigFile: cfgFile}
}
