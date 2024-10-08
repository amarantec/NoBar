package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateToken(customerId string) (string, error) {
	if customerId == "" {
		return "", fmt.Errorf("customer id is empty")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customerId": customerId,
		"exp":        time.Now().Add(time.Hour * 6).Unix(),
	})

	return token.SignedString([]byte(privateKey))
}

func VerifyToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(privateKey), nil
	})

	if err != nil {
		log.Printf("token err: %v", err)
		return "", errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		log.Printf("parse token err: %v", err)
		return "", errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	customerId, ok := claims["customerId"].(string)
	if !ok {
		return "", errors.New("customer id is not a valid float64")
	}

	return customerId, nil
}
