package service

import (
	"context"
	"swetelove/models"
	"swetelove/repositories"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type AdvertisementService struct {
	AdvertisementRepository *repositories.AdvertisementRepository
}

func NewAdvertisementService(db *gorm.DB, redis *redis.Client, ctx context.Context) *AdvertisementService {
	return &AdvertisementService{
		AdvertisementRepository: repositories.NewAdvertisementRepository(db, redis, ctx),
	}
}

func (s *AdvertisementService) GetAdvertisementByCode(code string) (*models.Advertisement, error) {
	return s.AdvertisementRepository.GetAdvertisementByCode(code)
}
