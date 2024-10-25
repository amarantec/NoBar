package routes

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"github.com/amarantec/nobar/internal/models"
    "github.com/amarantec/nobar/internal/utils"
	"github.com/gin-gonic/gin"
)

func getCartProducts(c *gin.Context) {
	customerId := c.MustGet("userType").(string)
    if customerId != utils.CustomerTokenType {
        c.JSON(http.StatusForbidden,
            gin.H{"message": "Access denied"})
        return
    }

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := service.GetCartProducts(ctxTimeout, customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could not get cart products",
				"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

func getCartItemsCount(c *gin.Context) {
    customerId := c.MustGet("userType").(string)
    if customerId != utils.CustomerTokenType {
        c.JSON(http.StatusForbidden,
            gin.H{"message": "Access denied"})
        return
    }

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := service.GetCartItemsCount(ctxTimeout, customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could not get cart items count",
				"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func addToCart(c *gin.Context) {
    customerId := c.MustGet("userType").(string)
    if customerId != utils.CustomerTokenType {
        c.JSON(http.StatusForbidden,
            gin.H{"message": "Access denied"})
        return
    }

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cartItem := models.Carts{
		CustomerID: customerId,
	}

	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "could not parse this request",
				"error": err.Error()})
		return
	}

	res, err := service.AddToCart(ctxTimeout, cartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could not add this item to cart",
				"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func updateQuantity(c *gin.Context) {
    customerId := c.MustGet("userType").(string)
    if customerId != utils.CustomerTokenType {
        c.JSON(http.StatusForbidden,
            gin.H{"message": "Access denied"})
        return
    }

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	productsId, err1 := strconv.Atoi(c.Param("productsId"))
	quantity, err2 := strconv.ParseInt(c.Param("quantity"), 10, 64)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "invalid parameters"})
		return
	}

	res, err := service.UpdateQuantity(ctxTimeout, customerId, uint(productsId), quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could not update this item quantity",
				"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, res)
}

func removeItemFromCart(c *gin.Context) {
    customerId := c.MustGet("userType").(string)
    if customerId != utils.CustomerTokenType {
        c.JSON(http.StatusForbidden,
            gin.H{"message": "Access denied"})
        return
    }

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	productsId, err := strconv.Atoi(c.Param("productsId"))
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "invalid parameter"})
		return
	}

	res, err := service.RemoveItemFromCart(ctxTimeout, customerId, uint(productsId))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could not remove this item from cart",
				"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, res)
}
