package services

import (
	"context"
	"errors"

	"github.com/amarantec/nobar/internal/models"
	"gorm.io/gorm"
)

func (s *ServicePostgres) InsertProduct(ctx context.Context, product models.Products) (models.Products, error) {
	if err :=
		s.Db.WithContext(ctx).
			Create(&product).Error; err != nil {
		return models.Products{}, err
	}
	return product, nil
}

func (s *ServicePostgres) ListProducts(ctx context.Context) ([]models.Products, error) {
	products := []models.Products{}

	if err :=
		s.Db.WithContext(ctx).
			Find(&products).Error; err != nil {
		return []models.Products{}, err
	}

	return products, nil
}

func (s *ServicePostgres) GetProduct(ctx context.Context, id uint) (models.Products, error) {
	product := models.Products{}

	if err :=
		s.Db.WithContext(ctx).
			Where("id = ?", id).
			First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Products{}, nil
		}
		return models.Products{}, err
	}

	return product, nil
}
