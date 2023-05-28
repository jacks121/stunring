package service

import (
	"context"
	"fmt"
	"swetelove/models"
	"swetelove/repositories"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type HeaderInfo struct {
	CategoryTree []*models.Category
	// Other fields here...
}

type CommonService struct {
	CategoryRepository *repositories.CategoryRepository
	// Other repositories here...
}

func NewCommonService(db *gorm.DB, redis *redis.Client, ctx context.Context) *CommonService {
	return &CommonService{
		CategoryRepository: repositories.NewCategoryRepository(db, redis, ctx),
		// Initialize other repositories here...
	}
}

func (s *CommonService) GetHeaderInfo() (*HeaderInfo, error) {
	info := &HeaderInfo{}

	// Get category tree
	categoryTree, err := s.CategoryRepository.GetCategoryTree()

	if err != nil {
		return nil, fmt.Errorf("failed to get category tree: %w", err)
	}
	info.CategoryTree = categoryTree

	// Get other info...
	// ...

	return info, nil
}
