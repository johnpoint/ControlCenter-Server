package cmd

import (
	"ControlCenter-Server/initHelper"
	"context"
	"github.com/spf13/cobra"
)

var httpServerCommand = &cobra.Command{
	Use:   "api",
	Short: "Start http server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		i := initHelper.Helper{}
		err := i.Init(ctx)
		if err != nil {
			panic(err)
			return
		}
	},
}
