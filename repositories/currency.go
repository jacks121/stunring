package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"swetelove/models"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type CurrencyRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
	Ctx   context.Context
}

func NewCurrencyRepository(db *gorm.DB, redis *redis.Client, ctx context.Context) *CurrencyRepository {
	return &CurrencyRepository{
		DB:    db,
		Redis: redis,
		Ctx:   ctx,
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
