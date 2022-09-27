package service

import (
	"context"
	"donasi/model"

	"donasi/repository"
)

type DonasiService interface {
	GetAllDonasi(ctx context.Context) (result model.ResponseKampanyeDonasi, err error)
}

type donasiService struct {
	dbRepo repository.DatabaseRepository
}

func NewDonasiService(dbRepo repository.DatabaseRepository) DonasiService {
	return &donasiService{dbRepo}
}

func (s *donasiService) GetAllDonasi(ctx context.Context) (result model.ResponseKampanyeDonasi, err error) {
	return result, nil
}