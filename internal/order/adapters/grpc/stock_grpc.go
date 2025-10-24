package grpc

import (
	"context"
	"github.com/iyu-Fang/gorder/common/genproto/orderpb"
	"github.com/iyu-Fang/gorder/common/genproto/stockpb"
	"github.com/sirupsen/logrus"
)

type StockGRPC struct {
	client stockpb.StockServiceClient
}

func NewStockGRPC(client stockpb.StockServiceClient) *StockGRPC {
	return &StockGRPC{client: client}
}

func (s StockGRPC) CheckItemsInStock(ctx context.Context, items []*orderpb.ItemWithQuantity) (*stockpb.CheckIfItemsInStockResponse, error) {
	resp, err := s.client.CheckIfItemsInStock(ctx, &stockpb.CheckIfItemsInStockRequest{Items: items})
	logrus.Info("stock grpc response ", resp)
	return resp, err

}

func (s StockGRPC) GetItems(ctx context.Context, itemIDs []string) ([]*orderpb.Item, error) {
	resp, err := s.client.GetItems(ctx, &stockpb.GetItemsRequest{ItemIDs: itemIDs})
	if err != nil {
		return nil, err
	}
	return resp.Items, nil
}
