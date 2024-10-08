package middlewares

import (
	"net/http"

	"github.com/amarantec/nobar/internal/utils"
	"github.com/gin-gonic/gin"
)

const CUSTOMERID = "customerId"

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"message": "token is empty"})
		return
	}

	customerId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate session",
			"error": err.Error()})
		return
	}

	c.Set(CUSTOMERID, customerId)
	c.Next()
}
