package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/amarantec/nobar/internal/models"
	"github.com/amarantec/nobar/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func welcomeCustomer(c *gin.Context) {
	customerId := uuid.New().String()
	newCustomer := models.Customer{
		CustomerID: customerId,
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "could not parse this request",
				"error": err.Error()})
		return
	}

	if res, err := service.WelcomeCustomer(ctxTimeout, newCustomer); err != nil || !res {
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "could not register costumer",
				"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(utils.CustomerTokenType, customerId)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "could not generate token",
				"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, token)
}
