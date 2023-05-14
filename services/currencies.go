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

type CurrencyService struct {
	currencyRepo *repositories.CurrencyRepository
	cacheConfig  *utils.CacheConfig
}

func NewCurrencyService(db *gorm.DB, redisClient *redis.Client) (*CurrencyService, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}
	if redisClient == nil {
		return nil, errors.New("redisClient cannot be nil")
	}

	cacheExpiration := time.Hour * 1 // 设置缓存过期时间，例如 1 小时
	cacheTimeout := time.Second * 5  // 设置缓存超时时间，例如 5 秒

	currencyRepo := repositories.NewCurrencyRepository(db)
	return &CurrencyService{
		currencyRepo: currencyRepo,
		cacheConfig:  utils.NewCacheConfig(redisClient, cacheExpiration, cacheTimeout),
	}, nil
}

const currencyCacheKey = "currencies"

func (s *CurrencyService) GetAllCurrencies() ([]models.Currency, error) {
	var currencies []models.Currency

	err := s.cacheConfig.GetWithCache(currencyCacheKey, &currencies, func() (interface{}, error) {
		return s.currencyRepo.GetAllCurrencies()
	})

	if err != nil {
		return nil, err
	}

	return currencies, nil
}
