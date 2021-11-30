package cmd

import (
	"ControlCenter/depend"
	"ControlCenter/pkg/bootstrap"
	"context"
	"github.com/spf13/cobra"
)

var tcpServerCommand = &cobra.Command{
	Use:   "tcp",
	Short: "Start tcp server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		i := bootstrap.Helper{}
		i.AddComponent(
			&depend.Config{
				Path: configPath,
			},
			&depend.TcpServer{},
		)
		err := i.Init(ctx)
		if err != nil {
			panic(err)
			return
		}

		forever := make(chan struct{})
		<-forever
	},
}
