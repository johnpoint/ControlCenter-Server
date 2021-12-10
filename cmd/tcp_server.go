package cmd

import (
	"ControlCenter/app/service/grpcService"
	"ControlCenter/depend"
	"ControlCenter/pkg/bootstrap"
	serverInfo "ControlCenter/proto/server_info"
	"context"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
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
			&depend.GrpcServer{
				Name: "push_task",
				AddFunc: func(grpcServer *grpc.Server) {
					serverInfo.RegisterPushToServerServer(grpcServer, &grpcService.PushToServerService{})
				},
			},
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
