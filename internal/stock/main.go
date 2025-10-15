package main

import (
	"context"
	"github.com/iyu-Fang/gorder/common/config"
	"github.com/iyu-Fang/gorder/common/genproto/stockpb"
	"github.com/iyu-Fang/gorder/common/server"
	"github.com/iyu-Fang/gorder/stock/ports"
	"github.com/iyu-Fang/gorder/stock/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("stock.service-name")
	serverType := viper.GetString("stock.server-to-run")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := service.NewApplication(ctx)
	switch serverType {
	case "grpc":
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer(application)
			stockpb.RegisterStockServiceServer(server, svc)
		})
	case "http":
		// Not in use now
	default:
		panic("unexpect server type")
	}
}
