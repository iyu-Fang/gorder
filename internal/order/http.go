package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct{}

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
