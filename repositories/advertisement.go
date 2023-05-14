package repositories

import (
	"swetelove/models"

	"gorm.io/gorm"
)

type AdvertisementRepository struct {
	db *gorm.DB
}

func NewAdvertisementRepository(db *gorm.DB) *AdvertisementRepository {
	return &AdvertisementRepository{db}
}

func (r *AdvertisementRepository) GetAdvertisementsByPosition(position string) ([]models.Advertisement, error) {
	var advertisements []models.Advertisement
	result := r.db.Where("position = ?", position).Find(&advertisements)
	if result.Error != nil {
		return nil, result.Error
	}
	return advertisements, nil
}
