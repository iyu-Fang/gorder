package service

import (
	"context"
	"github.com/iyu-Fang/gorder/order/app"
)

func NewApplication(ctx context.Context) app.Application {
	//orderRepo := adapters.NewMemoryOrderRepository()
	return app.Application{}
}
