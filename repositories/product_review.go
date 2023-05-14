package repositories

import (
	"swetelove/models"

	"gorm.io/gorm"
)

type ProductReviewRepository struct {
	db *gorm.DB
}

func NewProductReviewRepository(db *gorm.DB) *ProductReviewRepository {
	return &ProductReviewRepository{db}
}

func (r *ProductReviewRepository) GetReviewsByProductID(productID uint, page, perPage int) ([]models.ProductReview, error) {
	var reviews []models.ProductReview

	err := r.db.Preload("Images", "imageable_type = ?", "product_review").
		Where("product_id = ?", productID).
		Offset((page - 1) * perPage).Limit(perPage).
		Find(&reviews).Error

	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *ProductReviewRepository) GetReviewsWithImages(page, perPage int) ([]models.ProductReview, error) {
	var reviews []models.ProductReview

	subQuery := r.db.Table("images").
		Select("imageable_id").
		Where("imageable_type = ?", "product_review").
		Group("imageable_id")

	err := r.db.Preload("Images", "imageable_type = ?", "product_review").
		Where("id IN (?)", subQuery).
		Offset((page - 1) * perPage).Limit(perPage).
		Find(&reviews).Error

	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (repo *ProductReviewRepository) CreateProductReview(review *models.ProductReview) (*models.ProductReview, error) {
	if err := repo.db.Create(review).Error; err != nil {
		return nil, err
	}
	return review, nil
}
