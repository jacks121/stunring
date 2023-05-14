package repositories

import (
	"swetelove/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (r *CategoryRepository) GetCategoryByIDWithSorting(categoryID uint, sortType string, isDesc bool) ([]models.Product, error) {
	var category models.Category

	query := r.db.Preload("Products").First(&category, categoryID)

	switch sortType {
	case "price":
		if isDesc {
			query = query.Order("products.price DESC")
		} else {
			query = query.Order("products.price ASC")
		}
	case "recommend":
		if isDesc {
			query = query.Order("products.recommend DESC")
		} else {
			query = query.Order("products.recommend ASC")
		}
	case "newest":
		if isDesc {
			query = query.Order("products.created_at DESC")
		} else {
			query = query.Order("products.created_at ASC")
		}
	}

	if query.Error != nil {
		return nil, query.Error
	}

	return category.Products, nil
}

func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	result := r.db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (r *CategoryRepository) getChildren(categoryID uint) ([]models.Category, error) {
	var categories []models.Category
	result := r.db.Where("parent_id = ?", categoryID).Find(&categories)

	// If there's an error, return it
	if result.Error != nil {
		return nil, result.Error
	}

	// For each category, get its children recursively
	for i, category := range categories {
		children, err := r.getChildren(category.ID)
		if err != nil {
			return nil, err
		}
		categories[i].Children = children
	}

	return categories, nil
}

func (r *CategoryRepository) GetTree() ([]models.Category, error) {
	var categories []models.Category
	result := r.db.Where("parent_id = ?", 0).Find(&categories)

	// If there's an error, return it
	if result.Error != nil {
		return nil, result.Error
	}

	// For each category, get its children recursively
	for i, category := range categories {
		children, err := r.getChildren(category.ID)
		if err != nil {
			return nil, err
		}
		categories[i].Children = children
	}

	return categories, nil
}
