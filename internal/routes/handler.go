package routes

import (
	"github.com/amarantec/nobar/internal/services"
	"gorm.io/gorm"
)

var service services.Service

func ConfigureHandler(db *gorm.DB) {
	service = &services.ServicePostgres{
		Db: db,
	}
}
