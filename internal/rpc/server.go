package rpc

import (
	"context"
	"log"
	"net"

	"github.com/a-korkin/ecommerce/internal/core/adapters/db/services"
	pb "github.com/a-korkin/ecommerce/internal/proto"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type BillRPCServer struct {
	pb.UnimplementedBillServiceServer
	Service *services.BillService
}

func NewBillRPCServer(db *sqlx.DB) *BillRPCServer {
	service := services.NewBillService(db)
	return &BillRPCServer{Service: service}
}
func (r *BillRPCServer) CreateBill(
	ctx context.Context, userID *pb.UserID) (*pb.Bill, error) {
	out, err := r.Service.GetBillByUser(userID.Id)
	if err != nil {
		return nil, err
	}
	bill := pb.Bill{
		Id:         out.ID.String(),
		UserId:     userID.Id,
		TotalPrice: out.TotalSum,
	}
	return &bill, nil
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
