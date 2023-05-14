package services

import (
	"errors"
	"fmt"
	"swetelove/models"
	"swetelove/repositories"
	"swetelove/utils"

	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type CategoryService struct {
	categoryRepo *repositories.CategoryRepository
	cacheConfig  *utils.CacheConfig
}

func NewCategoryService(db *gorm.DB, redisClient *redis.Client) (*CategoryService, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}
	if redisClient == nil {
		return nil, errors.New("redisClient cannot be nil")
	}

	cacheExpiration := time.Hour * 1 // 设置缓存过期时间，例如 1 小时
	cacheTimeout := time.Second * 5  // 设置缓存超时时间，例如 5 秒

	return &CategoryService{
		categoryRepo: repositories.NewCategoryRepository(db),
		cacheConfig:  utils.NewCacheConfig(redisClient, cacheExpiration, cacheTimeout),
	}, nil
}

const categoryCacheKey = "categories"

// GetAllCategories returns all categories, along with their associated products, from the database.
func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category

	err := s.cacheConfig.GetWithCache(categoryCacheKey, &categories, func() (interface{}, error) {
		return s.categoryRepo.GetTree()
	})

	if err != nil {
		return nil, err
	}

	return categories, nil
}

// GetProductsByCategoryID returns all products associated with the specified category ID.
func (s *CategoryService) GetProductsByCategoryID(categoryID uint, sortType string, isDesc bool) ([]models.Product, error) {
	var products []models.Product

	err := s.cacheConfig.GetWithCache(categoryIDCacheKey(categoryID, sortType, isDesc), &products, func() (interface{}, error) {
		return s.categoryRepo.GetCategoryByIDWithSorting(categoryID, sortType, isDesc)
	})

	if err != nil {
		return nil, err
	}

	return products, nil
}

// categoryIDCacheKey returns the cache key for the specified category ID, sort type, and sort order.
func categoryIDCacheKey(categoryID uint, sortType string, isDesc bool) string {
	return fmt.Sprintf("category-%d-products-%s-%v", categoryID, sortType, isDesc)
}
