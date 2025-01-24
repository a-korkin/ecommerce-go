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
}

func NewBillRPCServer() *BillRPCServer {
	return &BillRPCServer{}
}
func (r *BillRPCServer) CreateBill(context.Context, *pb.UserID) (*pb.Bill, error) {
	return nil, nil
}

func (r *BillRPCServer) Run(ctx context.Context, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed create listener: %s", err)
	}
	s := grpc.NewServer()
	pb.RegisterBillServiceServer(s, r)

	go func() {
		log.Printf("grpc server started")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to run grpc server: %s", err)
		}
	}()
	<-ctx.Done()
	log.Printf("grpc server shutting down")
	s.GracefulStop()
	log.Printf("grpc server stoped")
}
