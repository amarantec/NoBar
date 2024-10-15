package services

import (
	"context"
	"errors"
    "strings"
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
			Preload("Categories").
			Find(&products).Error; err != nil {
		return []models.Products{}, err
	}

	return products, nil
}

func (s *ServicePostgres) GetProduct(ctx context.Context, id uint) (models.Products, error) {
	product := models.Products{}

	if err :=
		s.Db.WithContext(ctx).
			Preload("Categories").
			Where("id = ?", id).
			First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Products{}, nil
		}
		return models.Products{}, err
	}

	return product, nil
}

func (s *ServicePostgres) ListProductsByCategory(ctx context.Context, categoryUrl string) ([]models.Products, error) {
	products := []models.Products{}

	if err := s.Db.WithContext(ctx).
		Model(models.Products{}).
		Joins("JOIN categories ON categories.id = products.categories_id").
		Where("categories.url = ? ", categoryUrl).
		Find(&products).Error; err != nil {
		return []models.Products{}, err
	}

	return products, nil
}

func (s *ServicePostgres) SearchProducts(ctx context.Context, searchText string) ([]models.Products, error) {
    products := []models.Products{}

    if err :=
        s.Db.WithContext(ctx).
           Preload("Categories").
           Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ?", strings.ToLower("%"+searchText+"%"),
				strings.ToLower("%"+searchText+"%")).
           Find(&products).Error; err != nil {
                return []models.Products{}, err
                }

    return products, nil
}
