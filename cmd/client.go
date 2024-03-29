package cmd

import (
	"ControlCenter/depend"
	"ControlCenter/pkg/bootstrap"
	"ControlCenter/pkg/log"
	"context"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var clientCommand = &cobra.Command{
	Use:   "client",
	Short: "Start client",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		err := bootstrap.NewBoot(ctx,
			&depend.TcpClient{},
			&depend.TcpClientAuth{},
			&depend.PerformanceCollector{},
			&depend.HeartBeat{},
		).WithLogger(log.GetLogger()).Init()
		if err != nil {
			panic(err)
			return
		}

		stopChan := make(chan os.Signal)
		signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

		select {
		case signal := <-stopChan:
			log.Info("signal", log.String("info", "Catch signal:"+signal.String()+",and wait 5 sec"))
			time.Sleep(5 * time.Second)
			return
		}
	},
}
