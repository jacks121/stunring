package repositories

import (
	"context"
	"encoding/json"
	"swetelove/database"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type BaseRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
	Ctx   context.Context
}

func NewBaseRepository() *BaseRepository {
	return &BaseRepository{
		DB:    database.MysqlDB,
		Redis: database.RedisClient,
		Ctx:   database.GetContext(),
	}
}

func (r *BaseRepository) FetchFromRedis(key string, value interface{}) error {
	val, err := r.Redis.Get(r.Ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), value)
}

func (r *BaseRepository) FetchFromDB(value interface{}) error {
	return r.DB.Find(value).Error
}

func (r *BaseRepository) StoreInRedis(key string, value interface{}) error {
	dataJSON, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Redis.Set(r.Ctx, key, dataJSON, 0).Err()
}
