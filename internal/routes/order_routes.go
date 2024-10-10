package routes

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/amarantec/nobar/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func placeOrder(c *gin.Context) {
	customerId := c.GetString(middlewares.CUSTOMERID)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := service.PlaceOrder(ctxTimeout, customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could not place order",
				"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func getOrders(c *gin.Context) {
	customerId := c.GetString(middlewares.CUSTOMERID)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := service.GetOrders(ctxTimeout, customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could get orders",
				"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func getOrderDetails(c *gin.Context) {
	customerId := c.GetString(middlewares.CUSTOMERID)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	orderId, err := strconv.Atoi(c.Param("orderId"))
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "invalid parameter",
				"error": err.Error()})
		return
	}

	res, err := service.GetOrderDetails(ctxTimeout, customerId, uint(orderId))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could get order details",
				"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
