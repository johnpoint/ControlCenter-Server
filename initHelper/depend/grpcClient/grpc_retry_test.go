package grpcClient

import (
	"context"
	grpcConn "heywoods_golang_robot_server/app/dao/grpc"
	"heywoods_golang_robot_server/proto/customize"
	"testing"
	"time"
)

func TestTest(t *testing.T) {
	err := AddClient("crm", "127.0.0.1:9999")
	if err != nil {
		return
	}
	for i := 0; i < 20; i++ {
		conn, _ := grpcConn.GetCrmGrpcConn(context.Background())
		client := customize.NewManagementClient(conn)
		_, _ = client.GetShopID(context.TODO(), &customize.ClientRequest{})
		time.Sleep(time.Second)
	}

}
