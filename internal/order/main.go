package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iyu-Fang/gorder/common/config"
	"github.com/iyu-Fang/gorder/common/genproto/orderpb"
	"github.com/iyu-Fang/gorder/common/server"
	"github.com/iyu-Fang/gorder/order/ports"
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

	go server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		svc := ports.NewGRPCServer()
		orderpb.RegisterOrderServiceServer(server, svc)
	})

	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, HTTPServer{}, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil})
	})

}
