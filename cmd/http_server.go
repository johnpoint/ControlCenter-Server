package cmd

import (
	"ControlCenter/depend"
	"ControlCenter/pkg/bootstrap"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var httpServerCommand = &cobra.Command{
	Use:   "api",
	Short: "Start http server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		i := bootstrap.Helper{}
		i.AddComponent(
			&depend.Config{
				Path: configPath,
			},
			&depend.Redis{},
			&depend.Session{},
			&depend.MongoDB{},
			&depend.TaskProducer{},
			&depend.Api{},
		)
		err := i.Init(ctx)
		if err != nil {
			panic(err)
			return
		}

		stopChan := make(chan os.Signal)
		signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

		select {
		case signal := <-stopChan:
			fmt.Println("[System] Catch signal:" + signal.String() + ",and wait 30 sec")
			time.Sleep(30 * time.Second)
			return
		}
	},
}
