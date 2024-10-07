package services

import (
	"context"

	"github.com/amarantec/nobar/internal/models"
	"gorm.io/gorm"
)

type Service interface {
	InsertCategory(ctx context.Context, category models.Categories) (models.Categories, error)
	ListCategories(ctx context.Context) ([]models.Categories, error)
	GetCategory(ctx context.Context, id uint) (models.Categories, error)
	DeleteCategory(ctx context.Context, id uint) (bool, error)

	InsertProduct(ctx context.Context, product models.Products) (models.Products, error)
	ListProducts(ctx context.Context) ([]models.Products, error)
	GetProduct(ctx context.Context, id uint) (models.Products, error)
}

type ServicePostgres struct {
	Db *gorm.DB
}
