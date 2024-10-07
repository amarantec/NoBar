package models

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	Name string `json:"name" binding:"required" gorm:"not null; unique"`
	URL  string `json:"url" binding:"required" gorm:"not null; unique"`
}
