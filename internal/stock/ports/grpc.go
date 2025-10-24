package ports

import (
	"context"
	"github.com/iyu-Fang/gorder/common/genproto/stockpb"
	"github.com/iyu-Fang/gorder/stock/app"
	"github.com/iyu-Fang/gorder/stock/app/query"
)

type GRPCServer struct {
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (G GRPCServer) GetItems(ctx context.Context, request *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	//logrus.Info("rpc_request_in, stock.GetItems")
	//defer func() { logrus.Info("rpc_request_out, stock.GetItems") }()
	//fake := []*orderpb.Item{
	//	{
	//		ID: "fake-item-from-GetItems",
	//	},
	//}
	//return &stockpb.GetItemsResponse{Items: fake}, nil
	items, err := G.app.Queries.GetItems.Handle(ctx, query.GetItems{ItemIDs: request.ItemIDs})
	if err != nil {
		return nil, err
	}
	return &stockpb.GetItemsResponse{Items: items}, nil
}

func (G GRPCServer) CheckIfItemsInStock(ctx context.Context, request *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error) {
	//logrus.Info("rpc_request_in, stock.CheckIfItemsInStock")
	//defer func() { logrus.Info("rpc_request_out, stock.CheckIfItemsInStock") }()
	//return nil, nil
	items, err := G.app.Queries.CheckIfItemsInStock.Handle(ctx, query.CheckIfItemsInStock{Items: request.Items})
	if err != nil {
		return nil, err
	}
	return &stockpb.CheckIfItemsInStockResponse{
		Instock: 1,
		Items:   items,
	}, nil
}
