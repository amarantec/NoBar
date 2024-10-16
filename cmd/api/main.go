package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/amarantec/nobar/internal/db"
	"github.com/amarantec/nobar/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	ctx := context.Background()

	host := os.Getenv("DB_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("DB_PORT")
	server := os.Getenv("SERVER_PORT")
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", host, username, password, databaseName, port)

	db, err := db.DatabaseConnection(ctx, connectionString)
	if err != nil {
		panic(err)
	}

	if db == nil {
		log.Printf("nil database connection")
	}

	routes.ConfigureHandler(db)
	r := gin.Default()
	r.Static("/css", "../../web/css")
	r.LoadHTMLGlob("../../web/templates/*.html")
	routes.SetRoutes(r)

	r.Run(":" + server)
}

func loadEnv() {
	err := godotenv.Load("../../config/.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
