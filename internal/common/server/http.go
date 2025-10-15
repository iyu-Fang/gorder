package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RunHTTPServer(serviceName string, wrapper func(router *gin.Engine)) {
	addr := viper.Sub(serviceName).GetString("http-addr")
	if addr == "" {
		//TODO WARNING LOG
	}
	RunHTTPServeronAddr(addr, wrapper)
}

func RunHTTPServeronAddr(addr string, wrapper func(router *gin.Engine)) {
	apiRouter := gin.New()
	wrapper(apiRouter)
	apiRouter.Group("/api")

	if err := apiRouter.Run(addr); err != nil {
		panic(err)
	}
}
