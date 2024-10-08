package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustomerID string
	Name       string `json:"name" binding:"required" gorm:"not null"`
}
