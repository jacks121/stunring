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

type ProductReviewService struct {
	productReviewRepo *repositories.ProductReviewRepository
	cacheConfig       *utils.CacheConfig
}

func NewProductReviewService(db *gorm.DB, redisClient *redis.Client) (*ProductReviewService, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}
	if redisClient == nil {
		return nil, errors.New("redisClient cannot be nil")
	}

	cacheExpiration := time.Hour * 1 // 设置缓存过期时间，例如 1 小时
	cacheTimeout := time.Second * 5  // 设置缓存超时时间，例如 5 秒

	return &ProductReviewService{
		productReviewRepo: repositories.NewProductReviewRepository(db),
		cacheConfig:       utils.NewCacheConfig(redisClient, cacheExpiration, cacheTimeout),
	}, nil
}

func (s *ProductReviewService) GetReviewsByProductID(productID, page, perPage int) ([]models.ProductReview, error) {
	reviews, err := s.productReviewRepo.GetReviewsByProductID(uint(productID), page, perPage)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (s *ProductReviewService) GetReviewsWithImages(page, perPage int) ([]models.ProductReview, error) {
	reviews, err := s.productReviewRepo.GetReviewsWithImages(page, perPage)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (s *ProductReviewService) CreateProductReview(review *models.ProductReview) (*models.ProductReview, error) {
	createdReview, err := s.productReviewRepo.CreateProductReview(review)
	if err != nil {
		return nil, err
	}
	return createdReview, nil
}
