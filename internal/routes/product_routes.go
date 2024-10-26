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

func insertProduct(c *gin.Context) {
    adminId := c.MustGet("userType").(string)
    if adminId != utils.AdminTokenType {
        c.JSON(http.StatusForbidden,
            gin.H{"message": "Access denied"})
        return
    }

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
		c.HTML(http.StatusBadRequest,
			"bad_request.html",
			gin.H{"message": "invalid parameter",
				"error": err.Error()})
	}

	response, err := service.GetProduct(ctxTimeout, uint(id))
	if err != nil {
		c.HTML(http.StatusInternalServerError,
			"internal_server_error.html",
			gin.H{"message": "could not get this product",
				"error": err.Error()})
		return
	}

	if response.Name == "" {
		c.HTML(http.StatusNotFound,
			"not_found.html",
			gin.H{"message": "product not found"})
		return
	}

	c.HTML(http.StatusOK,
		"product.html",
		gin.H{"Products": response},
	)
}

func listProductsByCategory(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	categoryUrl := c.Param("categoryUrl")

	response, err := service.ListProductsByCategory(ctxTimeout, categoryUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could not list products",
				"error": err.Error()})
        return
	}

	c.JSON(http.StatusOK, response)
}

func searchProducts(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

    searchText := c.Param("searchText")

    response, err := service.SearchProducts(ctxTimeout, searchText)
    if err != nil {
        c.JSON(http.StatusInternalServerError,
            gin.H{"message": "could no search products",
                "error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, response)
}

func updateProduct(c *gin.Context) {
    adminId := c.MustGet("userType").(string)
    if adminId != utils.AdminTokenType {
        c.JSON(http.StatusForbidden,
            gin.H{"message": "Access denied"})
        return
    }

    ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

    product := models.Products{}

    if err :=
        c.ShouldBindJSON(&product); err != nil {
            c.JSON(http.StatusBadRequest,
                gin.H{"message": "could not decode this request",
                    "error": err.Error()})
             return
        }

    response, err := service.UpdateProduct(ctxTimeout, product)
    if err != nil {
        c.JSON(http.StatusInternalServerError,
            gin.H{"message": "could not update this product",
                "error": err.Error()})
        return
    }

    c.JSON(http.StatusNoContent, response)
}

func deleteProduct(c *gin.Context) {
    adminId := c.MustGet("userType").(string)
    if adminId != utils.AdminTokenType {
        c.JSON(http.StatusForbidden,
            gin.H{"message": "Access denied"})
        return
    }

    ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

    id, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "invalid parameter",
				"error": err.Error()})
	}

    response, err := service.DeleteProduct(ctxTimeout, uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError,
            gin.H{"message": "could not delete this product",
                "error": err.Error()})
        return
    }

    c.JSON(http.StatusNoContent, response)
}

