package cmd

import (
	"ControlCenter-Server/app/config"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Short: "ControlCenter-Server",
		Long: `ControlCenter-Server 中心控制服务器
https://github.com/johnpoint/ControlCenter-Server`,
	}
	cfgFile string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.json)")

	rootCmd.AddCommand(newServerCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile("config.json")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		config.LoadConfig(viper.ConfigFileUsed())
	}
}
