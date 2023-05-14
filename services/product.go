package services

import (
	"errors"
	"fmt"
	"math/rand"
	"swetelove/models"
	"swetelove/repositories"
	"swetelove/utils"

	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
	cacheConfig *utils.CacheConfig
}

func NewProductService(db *gorm.DB, redisClient *redis.Client) (*ProductService, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}
	if redisClient == nil {
		return nil, errors.New("redisClient cannot be nil")
	}

	cacheExpiration := time.Duration(rand.Intn(240)+60) * time.Minute
	cacheTimeout := time.Second * 5

	return &ProductService{
		productRepo: repositories.NewProductRepository(db),
		cacheConfig: utils.NewCacheConfig(redisClient, cacheExpiration, cacheTimeout),
	}, nil
}

const productCacheKey = "product"

func (s *ProductService) GetProductByID(productID uint) (*models.Product, error) {
	var product models.Product

	cacheKey := fmt.Sprintf("%s:%d", productCacheKey, productID)

	err := s.cacheConfig.GetWithCache(cacheKey, &product, func() (interface{}, error) {
		return s.productRepo.GetProductByID(productID)
	})

	if err != nil {
		return nil, err
	}

	return &product, nil
}
