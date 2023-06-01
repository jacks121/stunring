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
	redisKey := "categories"
	if err := r.FetchFromRedis(redisKey, &categories); err != nil {
		// 如果从Redis获取数据失败，就从数据库中获取数据
		if err := r.FetchFromDB(&categories); err != nil {
			return nil, err
		}
		// 将获取到的数据写入Redis
		if err := r.StoreInRedis(redisKey, &categories); err != nil {
			return nil, err
		}
	}
	return categories, nil
}

func (r *CategoryRepository) FetchFromDB(categories *[]*models.Category) error {
	// 在这里，我们可以添加特定的查询逻辑，比如预加载Images
	return r.DB.Preload("Images").Find(categories).Error
}

// GetCategoryTree 获取分类树
func (r *CategoryRepository) GetCategoryTree() ([]*models.Category, error) {
	var categoryTree []*models.Category
	redisKey := "category_tree"
	if err := r.FetchFromRedis(redisKey, &categoryTree); err != nil {
		// 如果从Redis获取数据失败，就从数据库中获取数据
		if err := r.FetchTreeFromDB(&categoryTree); err != nil {
			return nil, err
		}
		// 将获取到的数据写入Redis
		if err := r.StoreInRedis(redisKey, &categoryTree); err != nil {
			return nil, err
		}
	}
	return categoryTree, nil
}

// FetchTreeFromDB 从数据库中获取分类树
func (r *CategoryRepository) FetchTreeFromDB(categoryTree *[]*models.Category) error {
	rootCategories, err := r.getRootCategories()
	if err != nil {
		return err
	}
	*categoryTree = rootCategories
	return nil
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
