package repositories

import "swetelove/models"

type CollectionRepository struct {
	BaseRepository
}

func NewCollectionRepository() *CollectionRepository {
	return &CollectionRepository{
		BaseRepository: *NewBaseRepository(),
	}
}

func (r *CollectionRepository) GetCollectionRule(code string) (*models.Collection, error) {
	var collection models.Collection
	err := r.DB.Where("code = ?", code).First(&collection).Error
	return &collection, err
}
