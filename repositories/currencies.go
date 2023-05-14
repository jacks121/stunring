package repositories

import (
	"swetelove/models"

	"gorm.io/gorm"
)

type CurrencyRepository struct {
	db *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) *CurrencyRepository {
	return &CurrencyRepository{db}
}

func (repo *CurrencyRepository) GetAllCurrencies() ([]models.Currency, error) {
	var currencies []models.Currency

	err := repo.db.Find(&currencies).Error
	if err != nil {
		return nil, err
	}

	return currencies, nil
}
