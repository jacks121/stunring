package repositories

import (
	"context"
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
