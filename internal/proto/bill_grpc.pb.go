// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: bill.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BillService_CreateBill_FullMethodName = "/BillService/CreateBill"
)

// BillServiceClient is the client API for BillService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BillServiceClient interface {
	CreateBill(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Bill, error)
}

type billServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBillServiceClient(cc grpc.ClientConnInterface) BillServiceClient {
	return &billServiceClient{cc}
}

func (c *billServiceClient) CreateBill(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Bill, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Bill)
	err := c.cc.Invoke(ctx, BillService_CreateBill_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillServiceServer is the server API for BillService service.
// All implementations must embed UnimplementedBillServiceServer
// for forward compatibility.
type BillServiceServer interface {
	CreateBill(context.Context, *UserID) (*Bill, error)
	mustEmbedUnimplementedBillServiceServer()
}

// UnimplementedBillServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBillServiceServer struct{}

func (UnimplementedBillServiceServer) CreateBill(context.Context, *UserID) (*Bill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBill not implemented")
}
func (UnimplementedBillServiceServer) mustEmbedUnimplementedBillServiceServer() {}
func (UnimplementedBillServiceServer) testEmbeddedByValue()                     {}

// UnsafeBillServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillServiceServer will
// result in compilation errors.
type UnsafeBillServiceServer interface {
	mustEmbedUnimplementedBillServiceServer()
}

func RegisterBillServiceServer(s grpc.ServiceRegistrar, srv BillServiceServer) {
	// If the following call pancis, it indicates UnimplementedBillServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BillService_ServiceDesc, srv)
}

func _BillService_CreateBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillServiceServer).CreateBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillService_CreateBill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillServiceServer).CreateBill(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

// BillService_ServiceDesc is the grpc.ServiceDesc for BillService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BillService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BillService",
	HandlerType: (*BillServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBill",
			Handler:    _BillService_CreateBill_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bill.proto",
}
