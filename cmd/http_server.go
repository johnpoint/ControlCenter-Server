package cmd

import (
	"ControlCenter-Server/init"
	"context"
	"github.com/spf13/cobra"
)

var httpServerCommand = &cobra.Command{
	Use:   "api",
	Short: "Start http server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		i := init.Helper{}
		err := i.Init(ctx)
		if err != nil {
			panic(err)
			return
		}
	},
}
