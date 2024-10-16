package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {

	c.HTML(http.StatusOK, "home.html",
		nil)
}
