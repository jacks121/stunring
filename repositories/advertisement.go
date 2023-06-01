package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"swetelove/models"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type AdvertisementRepository struct {
	BaseRepository
}

func NewAdvertisementRepository() *AdvertisementRepository {
	return &AdvertisementRepository{
		BaseRepository: *NewBaseRepository(),
	}
}

// GetAdvertisementByCode 根据广告代码从缓存或数据库中获取广告
func (r *AdvertisementRepository) GetAdvertisementByCode(code string) (*models.Advertisement, error) {
	var ad models.Advertisement
	redisKey := fmt.Sprintf("advertisement:%v", code)

	// 尝试从 Redis 中获取广告
	val, err := r.Redis.Get(r.Ctx, redisKey).Result()
	if err == redis.Nil {
		// Redis 中不存在键，从数据库中获取广告
		if err := r.fetchAdvertisementFromMySQL(&ad, code); err != nil {
			return nil, err
		}

		// 更新 Redis 中的广告
		if err := r.updateAdvertisementInRedis(redisKey, &ad); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, fmt.Errorf("failed to get advertisement from Redis: %w", err)
	} else {
		// 从 Redis 中反序列化广告
		if err = json.Unmarshal([]byte(val), &ad); err != nil {
			return nil, fmt.Errorf("failed to unmarshal advertisement from Redis: %w", err)
		}
	}

	return &ad, nil
}

// fetchAdvertisementFromMySQL 从数据库中获取广告
func (r *AdvertisementRepository) fetchAdvertisementFromMySQL(ad *models.Advertisement, code string) error {
	if err := r.DB.Preload("Images").Where("code = ?", code).First(ad).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("advertisement not found in MySQL for code: %s", code)
		}
		return fmt.Errorf("failed to fetch advertisement from MySQL: %w", err)
	}
	return nil
}

// updateAdvertisementInRedis 更新 Redis 中的广告
func (r *AdvertisementRepository) updateAdvertisementInRedis(redisKey string, ad *models.Advertisement) error {
	redisVal, err := json.Marshal(ad)
	if err != nil {
		return fmt.Errorf("failed to marshal advertisement for Redis: %w", err)
	}

	if err := r.Redis.Set(r.Ctx, redisKey, redisVal, 0).Err(); err != nil {
		return fmt.Errorf("failed to set advertisement in Redis: %w", err)
	}

	return nil
}
