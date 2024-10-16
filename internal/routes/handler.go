package routes

import (
	"html/template"

	"github.com/amarantec/nobar/internal/services"
	"gorm.io/gorm"
)

var tmpl *template.Template

var service services.Service

func ConfigureHandler(db *gorm.DB) {
	service = &services.ServicePostgres{
		Db: db,
	}
}
