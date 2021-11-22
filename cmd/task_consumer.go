package cmd

import (
	"ControlCenter/depend"
	"ControlCenter/pkg/initHelper"
	"context"
	"github.com/spf13/cobra"
)

var taskConsumerCommand = &cobra.Command{
	Use:   "task",
	Short: "Start task consumer",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		i := initHelper.Helper{}
		i.AddDepend(
			&depend.TaskConsumer{},
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
