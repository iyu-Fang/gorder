package query

import (
	"context"
	"github.com/iyu-Fang/gorder/common/genproto/orderpb"
	"github.com/iyu-Fang/gorder/common/genproto/stockpb"
	//"github.com/iyu-Fang/gorder/common/genproto/stockpb"
)

type StockService interface {
	CheckItemsInStock(ctx context.Context, items []*orderpb.ItemWithQuantity) (*stockpb.CheckIfItemsInStockResponse, error)
	GetItems(ctx context.Context, itemIDs []string) ([]*orderpb.Item, error)
}
