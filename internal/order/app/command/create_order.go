package command

import (
	"context"
	"errors"
	"github.com/iyu-Fang/gorder/common/decorator"
	"github.com/iyu-Fang/gorder/common/genproto/orderpb"
	"github.com/iyu-Fang/gorder/order/app/query"
	domain "github.com/iyu-Fang/gorder/order/domain/order"
	"github.com/sirupsen/logrus"
)

type CreateOrder struct {
	CustomerID string
	Items      []*orderpb.ItemWithQuantity
}

type CreateOrderResult struct {
	OrderID string
}

type CreateOrderHandler decorator.CommandHandler[CreateOrder, *CreateOrderResult]

type createOrderHandler struct {
	orderRepo domain.Repository
	stockGRPC query.StockService
}

func NewCreateOrderHandler(
	orderRepo domain.Repository,
	stockGRPC query.StockService,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) CreateOrderHandler {
	if orderRepo == nil {
		panic("nil orderRepo")
	}
	return decorator.ApplyCommandDecorators[CreateOrder, *CreateOrderResult](
		createOrderHandler{orderRepo: orderRepo, stockGRPC: stockGRPC},
		logger,
		metricsClient)
}

func (c createOrderHandler) Handle(ctx context.Context, cmd CreateOrder) (*CreateOrderResult, error) {
	//err := c.stockGRPC.CheckItemsInStock(ctx, cmd.Items)
	//resp, err := c.stockGRPC.GetItems(ctx, []string{"123"})
	//logrus.Info("createOrderHandler || resp from stockGRPC.GetItems ", resp)
	//var stockResponse []*orderpb.Item
	//for _, item := range cmd.Items {
	//	stockResponse = append(stockResponse, &orderpb.Item{
	//		ID:       item.ID,
	//		Quantity: item.Quantity,
	//	})
	//}
	validItems, err := c.validate(ctx, cmd.Items)
	if err != nil {
		return nil, err
	}
	o, err := c.orderRepo.Create(ctx, &domain.Order{
		CustomerID: cmd.CustomerID,
		Items:      validItems,
	})
	if err != nil {
		return nil, err
	}
	return &CreateOrderResult{OrderID: o.ID}, nil
}

func (c createOrderHandler) validate(ctx context.Context, items []*orderpb.ItemWithQuantity) ([]*orderpb.Item, error) {
	if len(items) == 0 {
		return nil, errors.New("must have at least 1 item")
	}
	items = packItems(items)
	resp, err := c.stockGRPC.CheckItemsInStock(ctx, items)
	if err != nil {
		return nil, err
	}
	//var ids []string
	//for _, item := range items {
	//	ids = append(ids, item.ID)
	//}
	//return c.stockGRPC.GetItems(ctx, ids)
	return resp.Items, nil
}

func packItems(items []*orderpb.ItemWithQuantity) []*orderpb.ItemWithQuantity {
	merged := make(map[string]int32)
	for _, item := range items {
		merged[item.ID] += item.Quantity
	}
	var res []*orderpb.ItemWithQuantity
	for id, quantity := range merged {
		res = append(res, &orderpb.ItemWithQuantity{
			ID:       id,
			Quantity: quantity,
		})
	}
	return res
}
