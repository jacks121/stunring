package repositories

import (
	"swetelove/models"

	"gorm.io/gorm"
)

type ReviewRepository struct {
	BaseRepository
}

func NewReviewRepository() *ReviewRepository {
	return &ReviewRepository{
		BaseRepository: *NewBaseRepository(),
	}
}

func (r *ReviewRepository) GetReviewsWithImagesByCreatedAtDescending(limit int) ([]*models.Review, error) {
	var reviews []*models.Review

	err := r.DB.
		Order("created_at desc").
		Limit(limit).
		Preload("Images", func(db *gorm.DB) *gorm.DB {
			return db.Select("id")
		}).
		Find(&reviews).
		Error

	if err != nil {
		return nil, err
	}

	return reviews, nil
}
