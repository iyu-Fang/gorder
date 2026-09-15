package ports

import (
	"context"

	"github.com/iyu-Fang/gorder/common/genproto/stockpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (G GRPCServer) GetItems(ctx context.Context, request *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetItems not Implemented")
}

func (G GRPCServer) CheckIfItemsInStock(ctx context.Context, request *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method CheckItemsInStock not Implemented")
}
