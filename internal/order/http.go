package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iyu-Fang/gorder/order/app"
)

type HTTPServer struct {
	app app.Application
}

func (H HTTPServer) PostCustomerCustomerIdOrders(c *gin.Context, customerId string) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemented",
	})
}

func (H HTTPServer) GetCustomerCustomerIdOrdersOrderId(c *gin.Context, customerId string, orderId string) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "not implemented",
	})
}
