package repositories

import (
	"encoding/json"
	"fmt"
	"swetelove/models"
)

type CategoryRepository struct {
	BaseRepository
}

// NewCategoryRepository 创建 CategoryRepository 实例
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		BaseRepository: *NewBaseRepository(),
	}
}

// GetAllCategories 获取所有分类
func (r *CategoryRepository) GetAllCategories() ([]*models.Category, error) {
	var categories []*models.Category

	// Try to fetch from Redis
	redisKey := "categories"
	val, err := r.Redis.Get(r.Ctx, redisKey).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(val), &categories); err != nil {
			return nil, fmt.Errorf("failed to unmarshal categories from Redis: %w", err)
		}
		return categories, nil
	}

	// Fetch from MySQL
	if err = r.DB.Preload("Images").Find(&categories).Error; err != nil {
		return nil, fmt.Errorf("failed to get categories from DB: %w", err)
	}

	// Update Redis
	redisVal, err := json.Marshal(categories)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal categories for Redis: %w", err)
	}
	if err = r.Redis.Set(r.Ctx, redisKey, redisVal, 0).Err(); err != nil {
		return nil, fmt.Errorf("failed to set categories in Redis: %w", err)
	}

	return categories, nil
}

// GetCategoryTree 获取分类树
func (r *CategoryRepository) GetCategoryTree() ([]*models.Category, error) {
	// 尝试从 Redis 获取分类树
	redisKey := "category_tree"
	val, err := r.Redis.Get(r.Ctx, redisKey).Result()
	if err == nil {
		var categoryTree []*models.Category
		if err = json.Unmarshal([]byte(val), &categoryTree); err != nil {
			return nil, fmt.Errorf("failed to unmarshal category tree from Redis: %w", err)
		}
		return categoryTree, nil
	}

	// 从数据库获取分类树
	rootCategories, err := r.getRootCategories()
	if err != nil {
		return nil, err
	}

	// 更新 Redis 中的分类树
	redisVal, err := json.Marshal(rootCategories)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal category tree for Redis: %w", err)
	}
	if err = r.Redis.Set(r.Ctx, redisKey, redisVal, 0).Err(); err != nil {
		return nil, fmt.Errorf("failed to set category tree in Redis: %w", err)
	}

	return rootCategories, nil
}

// getRootCategories 获取根分类
func (r *CategoryRepository) getRootCategories() ([]*models.Category, error) {
	var rootCategories []*models.Category

	if err := r.DB.Preload("Images").Where("parent_id = 0").Find(&rootCategories).Error; err != nil {
		return nil, fmt.Errorf("failed to get root categories from DB: %w", err)
	}

	for _, category := range rootCategories {
		if err := r.findSubcategories(category); err != nil {
			return nil, fmt.Errorf("failed to get subcategories for category ID %d: %w", category.ID, err)
		}
	}

	return rootCategories, nil
}

// findSubcategories 递归查找子分类
func (r *CategoryRepository) findSubcategories(category *models.Category) error {
	var subcategories []*models.Category

	if err := r.DB.Preload("Images").Where("parent_id = ?", category.ID).Find(&subcategories).Error; err != nil {
		return fmt.Errorf("failed to get subcategories from DB: %w", err)
	}

	category.Subcategories = subcategories

	for _, subcategory := range subcategories {
		if err := r.findSubcategories(subcategory); err != nil {
			return fmt.Errorf("failed to get subcategories for category ID %d: %w", subcategory.ID, err)
		}
	}

	return nil
}

// GetCategoriesByIds 根据分类 ID 获取分类列表
func (r *CategoryRepository) GetCategoriesByIds(ids []int) ([]*models.Category, error) {
	var categories []*models.Category

	// Generate Redis key based on ids
	redisKey := fmt.Sprintf("categories:%v", ids)
	val, err := r.Redis.Get(r.Ctx, redisKey).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(val), &categories); err != nil {
			return nil, fmt.Errorf("failed to unmarshal categories from Redis: %w", err)
		}
		return categories, nil
	}

	// Fetch from MySQL
	if err = r.DB.Preload("Images").Where("id IN ?", ids).Find(&categories).Error; err != nil {
		return nil, fmt.Errorf("failed to get categories from DB: %w", err)
	}

	// Update Redis
	redisVal, err := json.Marshal(categories)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal categories for Redis: %w", err)
	}
	if err = r.Redis.Set(r.Ctx, redisKey, redisVal, 0).Err(); err != nil {
		return nil, fmt.Errorf("failed to set categories in Redis: %w", err)
	}

	return categories, nil
}
