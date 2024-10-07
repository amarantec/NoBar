package routes

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/amarantec/nobar/internal/models"
	"github.com/gin-gonic/gin"
)

func insertProduct(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newProduct := models.Products{}

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "could not parse this request",
				"error": err.Error()})
	}

	response, err := service.InsertProduct(ctxTimeout, newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could not insert this product",
				"error": err.Error()})
	}

	c.JSON(http.StatusCreated, response)
}

func listProducts(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := service.ListProducts(ctxTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could not get list of products",
				"error": err.Error()})
	}

	c.JSON(http.StatusOK, response)
}

func getProduct(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "invalid parameter",
				"error": err.Error()})
	}

	response, err := service.GetProduct(ctxTimeout, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could not get this product",
				"error": err.Error()})
	}

	c.JSON(http.StatusOK, response)
}
