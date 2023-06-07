package repositories

import (
	"swetelove/models"
)

type ReviewRepository struct {
	BaseRepository
}

func NewReviewRepository() *ReviewRepository {
	return &ReviewRepository{
		BaseRepository: *NewBaseRepository(),
	}
}

// GetReviewsWithImagesByPageAndCount retrieves reviews with images based on the page number and page size,
// and returns the total count of reviews.
func (r *ReviewRepository) GetReviewsWithImagesByPageAndCount(page, pageSize int) ([]*models.Review, int, error) {
	var reviews []*models.Review
	var total int64

	offset := (page - 1) * pageSize

	// Subquery to retrieve imageable_ids of reviews with images
	subQuery := r.DB.Model(&models.Image{}).
		Select("imageable_id").
		Where("imageable_type = ?", "reviews").
		Distinct()

	err := r.DB.Model(&models.Review{}).
		Joins("JOIN (?) AS sub ON sub.imageable_id = reviews.id", subQuery).
		Count(&total).
		Order("created_at desc").
		Offset(offset).
		Limit(pageSize).
		Preload("Images").
		Find(&reviews).
		Error

	if err != nil {
		return nil, 0, err
	}

	return reviews, int(total), nil
}
