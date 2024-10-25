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

func (s *ServicePostgres) UpdateProduct(ctx context.Context, product models.Products) (bool, error) {
    result :=
        s.Db.WithContext(ctx).
            Where("id = ?", product.ID).
            Updates(models.Products{
                Name: product.Name,
                Description: product.Description,
                ImageURL: product.ImageURL,
                CategoriesID: product.CategoriesID,
                Price: product.Price,
            })

    if result.Error != nil {
        return false, result.Error
    } else if result.RowsAffected == 0 {
        return false, nil
    }

    return true, nil
}

func (s *ServicePostgres) DeleteProduct(ctx context.Context, id uint) (bool, error) {
    if err :=
        s.Db.WithContext(ctx).
            Delete(&models.Products{}, id).Error; err != nil {
                if errors.Is(err, gorm.ErrRecordNotFound) {
                    return false, nil
                }
                return false, err
            }

    return true, nil
}
