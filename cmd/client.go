package cmd

import (
	"ControlCenter/initHelper"
	"ControlCenter/initHelper/depend"
	"context"
	"github.com/spf13/cobra"
)

var clientCommand = &cobra.Command{
	Use:   "client",
	Short: "Start client",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		i := initHelper.Helper{}
		i.AddDepend(
			&depend.GrpcClientServer{},
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
