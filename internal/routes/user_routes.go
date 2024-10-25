package routes

import (
    "context"
    "net/http"
    "time"
    "github.com/amarantec/nobar/internal/models"
    "github.com/gin-gonic/gin"
    "github.com/amarantec/nobar/internal/utils"
)

func register(c *gin.Context) {
   user := models.Users{} 
   ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
   defer cancel()

   if err :=
       c.ShouldBindJSON(&user); err != nil {
           c.JSON(http.StatusBadRequest,
                gin.H{"message": "could not parse this request",
                    "error": err.Error()})
       return
   }

    response, err := service.Register(ctxTimeout, user)
    if err != nil {
        c.JSON(http.StatusInternalServerError,
            gin.H{"message": "could not register this user",
                "error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, response)
}

func login(c *gin.Context) {
    user := models.Users{}

    ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err :=
        c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest,
                gin.H{"message": "could not parse this request",
                    "error": err.Error()})
         return
    }
    
    response, err := service.Login(ctxTimeout, user)
    if err != nil {
        c.JSON(http.StatusInternalServerError,
            gin.H{"message": "could not validate this user",
                "error": err.Error()})
        return
    }

    
    token, err := utils.GenerateToken(utils.AdminTokenType, response)
    if err != nil {
        c.JSON(http.StatusBadRequest,
            gin.H{"message": "could not generate token",
                "error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, token)
}
