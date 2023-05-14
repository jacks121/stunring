package repositories

import (
	"swetelove/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (repo *ProductRepository) GetProductByID(productID uint) (*models.Product, error) {
	var product models.Product

	err := repo.db.Preload("Attributes", func(db *gorm.DB) *gorm.DB {
		return db.Preload("AttributeValues")
	}).Preload("SKUs", func(db *gorm.DB) *gorm.DB {
		return db.Preload("AttributeValues")
	}).Preload("Reviews", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC").Limit(10)
	}).Preload("Images").First(&product, productID).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *ProductRepository) GetNewProducts(limit int) ([]models.Product, error) {
	var products []models.Product

	err := repo.db.Preload("Images").Order("created_at desc").Limit(limit).Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}
