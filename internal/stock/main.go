package main

import (
	"github.com/iyu-Fang/gorder/common/genproto/stockpb"
	"github.com/iyu-Fang/gorder/stock/ports"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	serviceName := viper.GetString("stock.service-name")
	serverType := viper.GetString("stock.server-to-run")

	switch serverType {
	case "grpc":
		ports.RunGRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer()
			stockpb.RegisterStockServiceServer(server, svc)
		})
	case "http":
		// Not in use now
	default:
		panic("unexpect server type")
	}
}
