package ports

import (
	"context"

	"github.com/iyu-Fang/gorder/common/genproto/orderpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (G GRPCServer) CreateOrder(ctx context.Context, request *orderpb.CreateOrderRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateOrder not Implemented")
}

func (G GRPCServer) GetOrder(ctx context.Context, request *orderpb.GetOrderRequest) (*orderpb.Order, error) {
	return nil, status.Error(codes.Unimplemented, "method GetOrder not Implemented")
}

func (G GRPCServer) UpdateOrder(ctx context.Context, order *orderpb.Order) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "method UpdateOrder not Implemented")
}
