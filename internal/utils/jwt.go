package utils

import (
	"fmt"
	"os"
	"time"
    "net/http"
	"github.com/golang-jwt/jwt/v5"
)

const (
    CustomerTokenType = "customer"
    AdminTokenType = "admin"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateToken(userType string, id interface{}) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "type": userType,
        "id":   id,
        "exp":  time.Now().Add(time.Hour * 6).Unix(), // Token expira em 6 horas
    }) 

    return token.SignedString([]byte(privateKey))
}

func ValidateToken(tokenString string) (string, interface{}, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, http.ErrNotSupported
        }
        return []byte(privateKey), nil
    })

    if err != nil {
        return "", nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims["type"].(string), claims["id"], nil // Retorna o tipo de usuário e ID
    }

    return "", nil, fmt.Errorf("invalid token: %v", err)
}


