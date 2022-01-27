package cmd

import (
	"ControlCenter/depend"
	"ControlCenter/pkg/bootstrap"
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
	cobra.OnInitialize(func() {
		if configPath == "" {
			configPath = "config_local.yaml"
		}
		bootstrap.AddGlobalComponent(
			&depend.Config{
				Path: configPath,
			},
			&depend.Logger{},
		)
	})
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "config_local.yaml", "config file (default is ./config_local.yaml)")
	rootCmd.AddCommand(
		httpServerCommand,          // API服务
		clientCommand,              // 上报客户端
		genConfigCommand,           // 生成空配置文件
		taskConsumerCommand,        // 下发任务消费者
		tcpServerCommand,           // tcp
		performanceConsumerCommand, // 性能信息消费者
		tcpServerConsumerCommand,   // tcp信息消费者
	)
}
