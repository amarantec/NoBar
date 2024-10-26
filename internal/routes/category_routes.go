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

func insertCategory(c *gin.Context) {
    adminId := c.MustGet("userType").(string)
    if adminId != utils.AdminTokenType {
        c.JSON(http.StatusForbidden,
            gin.H{"message": "Access denied"})
        return
    }

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newCategory := models.Categories{}

	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "could not parse this request",
				"error": err.Error()})
	}

	response, err := service.InsertCategory(ctxTimeout, newCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could not insert this category",
				"error": err.Error()})
	}

	c.JSON(http.StatusCreated, response)
}

func listCategories(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := service.ListCategories(ctxTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could not get list of categories",
				"error": err.Error()})
	}

	c.JSON(http.StatusOK, response)
}

func getCategory(c *gin.Context) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "invalid parameter",
				"error": err.Error()})
	}

	response, err := service.GetCategory(ctxTimeout, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could not get this category",
				"error": err.Error()})
	}

	c.JSON(http.StatusOK, response)
}


func updateCategory(c *gin.Context) {
    adminId := c.MustGet("userType").(string)
    if adminId != utils.AdminTokenType {
        c.JSON(http.StatusForbidden,
            gin.H{"message": "Access denied"})
        return
    }

    ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    category := models.Categories{}

    if err :=
        c.ShouldBindJSON(&category); err != nil {
            c.JSON(http.StatusBadRequest,
                gin.H{"message": "could not decode this request",
                    "error": err.Error()})
            return
        }
    
    response, err := service.UpdateCategory(ctxTimeout, category)
    if err != nil {
        c.JSON(http.StatusInternalServerError,
            gin.H{"message": "could not update this category",
                "error": err.Error()})
        return
    }

    c.JSON(http.StatusNoContent, response)
}

func deleteCategory(c *gin.Context) {
    adminId := c.MustGet("userType").(string)
    if adminId != utils.AdminTokenType {
        c.JSON(http.StatusForbidden,
            gin.H{"message": "Access denied"})
        return
    }

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "invalid parameter",
				"error": err.Error()})
	}

	response, err := service.DeleteCategory(ctxTimeout, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could not delete this category",
				"error": err.Error()})
	}

	c.JSON(http.StatusNoContent, response)
}
