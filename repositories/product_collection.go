package repositories

import (
	"swetelove/models"

	"gorm.io/gorm"
)

type ProductCollectionRepository struct {
	db *gorm.DB
}

func NewProductCollectionRepository(db *gorm.DB) *ProductCollectionRepository {
	return &ProductCollectionRepository{db}
}

func (r *ProductCollectionRepository) GetProductCollectionByID(id uint) (*models.ProductCollection, error) {
	var collection models.ProductCollection
	result := r.db.First(&collection, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &collection, nil
}
