package services

import (
	"context"

	"github.com/amarantec/nobar/internal/models"
)

func (s *ServicePostgres) WelcomeCustomer(ctx context.Context, customer models.Customer) (bool, error) {
	if err :=
		s.Db.WithContext(ctx).
			Create(&customer).Error; err != nil {
		return false, err
	}

	return true, nil
}

