package service

import (
	"context"
	"github.com/iyu-Fang/gorder/common/metrics"
	"github.com/iyu-Fang/gorder/order/adapters"
	"github.com/iyu-Fang/gorder/order/app"
	"github.com/iyu-Fang/gorder/order/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	//orderRepo := adapters.NewMemoryOrderRepository()
	orderRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			GetCustomerOrder: query.NewGetCustomerOrderHandler(orderRepo, logger, metricClient),
		},
	}
}
