package services

import (
	"context"
	"errors"

	"github.com/amarantec/nobar/internal/models"
	"gorm.io/gorm"
)

func (s *ServicePostgres) InsertCategory(ctx context.Context, category models.Categories) (models.Categories, error) {
	if s.Db == nil {
		panic("nil database connection")
	}

	if err :=
		s.Db.WithContext(ctx).
			Create(&category).Error; err != nil {
		return models.Categories{}, err
	}

	return category, nil
}

func (s *ServicePostgres) ListCategories(ctx context.Context) ([]models.Categories, error) {
	categories := []models.Categories{}

	if err :=
		s.Db.WithContext(ctx).
			Find(&categories).Error; err != nil {
		return []models.Categories{}, err
	}

	return categories, nil
}

func (s *ServicePostgres) GetCategory(ctx context.Context, id uint) (models.Categories, error) {
	category := models.Categories{}

	if err :=
		s.Db.WithContext(ctx).
			Where("id = ?", id).
			First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Categories{}, nil
		}
		return models.Categories{}, err
	}

	return category, nil
}

func (s *ServicePostgres) DeleteCategory(ctx context.Context, id uint) (bool, error) {
	if err :=
		s.Db.WithContext(ctx).
			Where("id = ?", id).
			Unscoped().Delete(&models.Categories{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
