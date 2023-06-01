package repositories

import (
	"encoding/json"
	"fmt"
	"swetelove/models"
)

type CurrencyRepository struct {
	BaseRepository
}

func NewCurrencyRepository() *CurrencyRepository {
	return &CurrencyRepository{
		BaseRepository: *NewBaseRepository(),
	}
}

func (r *CurrencyRepository) GetAllCurrencies() ([]*models.Currency, error) {
	var currencies []*models.Currency

	// Check if data is available in Redis
	val, err := r.Redis.Get(r.Ctx, "currencies").Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &currencies); err != nil {
			return nil, fmt.Errorf("failed to unmarshal currencies from Redis: %w", err)
		}
		return currencies, nil
	}

	// Fetch data from the database
	if err := r.DB.Find(&currencies).Error; err != nil {
		return nil, fmt.Errorf("failed to get currencies from DB: %w", err)
	}

	// Store data in Redis
	currenciesJSON, err := json.Marshal(currencies)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal currencies for Redis: %w", err)
	}
	if err := r.Redis.Set(r.Ctx, "currencies", currenciesJSON, 0).Err(); err != nil {
		return nil, fmt.Errorf("failed to set currencies in Redis: %w", err)
	}

	return currencies, nil
}
