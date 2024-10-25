package middlewares

import (
	"net/http"

	"github.com/amarantec/nobar/internal/utils"
	"github.com/gin-gonic/gin"
)


func Auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"message": "token is empty"})
		return
	}

	userType, id, err := utils.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate session",
			"error": err.Error()})
		return
	}

    c.Set("userType", userType)
	c.Set("id", id)
	c.Next()
}
