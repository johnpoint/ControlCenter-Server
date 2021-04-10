package cmd

import (
	"ControlCenter-Server/app/router"
	"github.com/spf13/cobra"
)

var (
	newServerCmd = &cobra.Command{
		Use:   "start",
		Short: "start ControlCenter-Server http api server",
		Long: `ControlCenter-Server 中心控制服务器
https://github.com/johnpoint/ControlCenter-Server`,
		Run: func(cmd *cobra.Command, args []string) {
			router.Run()
		},
	}
)
