package cmd

import (
	"ControlCenter/initHelper"
	"ControlCenter/initHelper/depend"
	"context"
	"github.com/spf13/cobra"
)

var httpServerCommand = &cobra.Command{
	Use:   "api",
	Short: "Start http server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		i := initHelper.Helper{}
		i.AddDepend(&depend.MongoDB{})
		i.AddDepend(&depend.Redis{})
		i.AddDepend(&depend.Api{})
		err := i.Init(ctx)
		if err != nil {
			panic(err)
			return
		}

		forever := make(chan struct{})
		<-forever
	},
}
