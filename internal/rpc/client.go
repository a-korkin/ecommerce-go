package rpc

import (
	pb "github.com/a-korkin/ecommerce/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClient(addr string) (pb.BillServiceClient, error) {
	conn, err := grpc.NewClient(
		addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewBillServiceClient(conn), nil
}
