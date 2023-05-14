package services

import (
	"errors"
	"swetelove/models"
	"swetelove/repositories"
	"swetelove/utils"

	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type AdvertisementService struct {
	advertisementRepo *repositories.AdvertisementRepository
	cacheConfig       *utils.CacheConfig
}

func NewAdvertisementService(db *gorm.DB, redisClient *redis.Client) (*AdvertisementService, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}
	if redisClient == nil {
		return nil, errors.New("redisClient cannot be nil")
	}

	cacheExpiration := time.Hour * 1 // 设置缓存过期时间，例如 1 小时
	cacheTimeout := time.Second * 5  // 设置缓存超时时间，例如 5 秒

	return &AdvertisementService{
		advertisementRepo: repositories.NewAdvertisementRepository(db),
		cacheConfig:       utils.NewCacheConfig(redisClient, cacheExpiration, cacheTimeout),
	}, nil
}

const advertisementCacheKey = "advertisements"

func (s *AdvertisementService) GetAdvertisementsByPosition(position string) ([]models.Advertisement, error) {
	var advertisements []models.Advertisement

	err := s.cacheConfig.GetWithCache(advertisementCacheKey+":"+position, &advertisements, func() (interface{}, error) {
		return s.advertisementRepo.GetAdvertisementsByPosition(position)
	})

	if err != nil {
		return nil, err
	}

	return advertisements, nil
}
