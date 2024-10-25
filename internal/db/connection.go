package db

import (
	"context"

	"github.com/amarantec/nobar/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DatabaseConnection(ctx context.Context, connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect to database")
	}

	if err := db.AutoMigrate(
		&models.Products{},
		&models.Categories{},
		&models.Carts{},
		&models.Customer{},
        &models.Orders{},
        &models.OrderItems{},
        &models.Users{},
	); err != nil {
		panic("Failed to create tables")
	}

	return db, nil
}
