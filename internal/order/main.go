package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/iyu-Fang/gorder/common/config"
	"github.com/iyu-Fang/gorder/common/genproto/orderpb"
	"github.com/iyu-Fang/gorder/common/server"
	"github.com/iyu-Fang/gorder/order/ports"
	"github.com/iyu-Fang/gorder/order/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("order.service-name")
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	applicaiton := service.NewApplication(ctx)
	go server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		svc := ports.NewGRPCServer(applicaiton)
		orderpb.RegisterOrderServiceServer(server, svc)
	})

	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, HTTPServer{
			app: applicaiton,
		}, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil})
	})

}
