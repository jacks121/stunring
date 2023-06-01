package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"swetelove/models"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type AdvertisementRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
	Ctx   context.Context
}

func NewAdvertisementRepository(db *gorm.DB, redis *redis.Client, ctx context.Context) *AdvertisementRepository {
	return &AdvertisementRepository{
		DB:    db,
		Redis: redis,
		Ctx:   ctx,
	}
}

func (r *AdvertisementRepository) GetAdvertisementByCode(code string) (*models.Advertisement, error) {
	var ad models.Advertisement
	// Generate Redis key based on code
	redisKey := fmt.Sprintf("advertisement:%v", code)
	// val, err := r.Redis.Get(r.Ctx, redisKey).Result()
	// if err == nil {
	// 	if err = json.Unmarshal([]byte(val), &ad); err != nil {
	// 		return nil, fmt.Errorf("failed to unmarshal advertisement from Redis: %w", err)
	// 	}
	// 	return &ad, nil
	// }
	// Fetch from MySQL
	if err := r.DB.Preload("Images").Where("code = ?", code).First(&ad).Error; err != nil {
		return nil, err
	}
	// Update Redis
	redisVal, err := json.Marshal(ad)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal advertisement for Redis: %w", err)
	}
	if err = r.Redis.Set(r.Ctx, redisKey, redisVal, 0).Err(); err != nil {
		return nil, fmt.Errorf("failed to set advertisement in Redis: %w", err)
	}
	return &ad, nil
}
