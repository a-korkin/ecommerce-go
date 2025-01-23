package rpc

import (
	"context"
	pb "github.com/a-korkin/ecommerce/internal/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type BillRPCServer struct {
	pb.UnimplementedBillServiceServer
	GRPCServer *grpc.Server
}

func NewBillRPCServer() *BillRPCServer {
	return &BillRPCServer{}
}

func (r *BillRPCServer) Run(ctx context.Context, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed create listener: %s", err)
	}
	s := grpc.NewServer()
	pb.RegisterBillServiceServer(s, r)
	r.GRPCServer = s
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to run grpc server: %s", err)
	}
}
