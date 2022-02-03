package grpcservice

import (
	"ControlCenter/pkg/errorhelper"
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

func RunGrpcServer(listen string, addFunc func(grpcServer *grpc.Server)) error {
	lis, err := net.Listen("tcp", listen)
	if err != nil {
		return err
	}
	grpcPanicRecover := func(p interface{}) (err error) {
		stack := string(debug.Stack())
		err = &errorhelper.Err{
			Code:    errorhelper.Unknown.Code,
			Message: fmt.Sprintf("panic triggered: %v\nstack: %+v", p, stack),
		}
		// 对外有限暴露
		return fmt.Errorf("%+v", p)
	}

	opts := []grpcRecovery.Option{
		grpcRecovery.WithRecoveryHandler(grpcPanicRecover),
	}

	var servOpts = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * 11),
		grpcMiddleware.WithStreamServerChain(
			grpcRecovery.StreamServerInterceptor(opts...),
		),
		grpcMiddleware.WithUnaryServerChain(
			grpcRecovery.UnaryServerInterceptor(opts...),
		),
	}

	grpcServer := grpc.NewServer(servOpts...)

	addFunc(grpcServer)

	reflection.Register(grpcServer)
	go func() {
		fmt.Println("[grpcServer] start")
		if err := grpcServer.Serve(lis); err != nil {
			fmt.Println(err.Error())
		}
	}()
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-stopChan
	grpcServer.GracefulStop()
	return nil
}
