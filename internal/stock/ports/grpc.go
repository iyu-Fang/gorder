package ports

import (
	"context"

	"github.com/iyu-Fang/gorder/common/genproto/stockpb"
	"github.com/iyu-Fang/gorder/stock/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (G GRPCServer) GetItems(ctx context.Context, request *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetItems not Implemented")
}

func (G GRPCServer) CheckIfItemsInStock(ctx context.Context, request *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method CheckItemsInStock not Implemented")
}
