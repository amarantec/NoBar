package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name         string      `json:"name" binding:"required" gorm:"not null; unique"`
	Description  string      `json:"description" binding:"required" gorm:"not null"`
	ImageURL     string      `json:"image_url" binding:"required" gorm:"not null"`
	CategoriesID uint        `json:"categories_id" binding:"required" gorm:"not null"`
	Categories   *Categories `json:"categories,omitempty"`
	Price        float64     `json:"price" binding:"required" gorm:"not null"`
}
