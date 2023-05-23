package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"swetelove/models"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
	ES    *elastic.Client
	Ctx   context.Context
}

func NewProductRepository(db *gorm.DB, redis *redis.Client, es *elastic.Client, ctx context.Context) *ProductRepository {
	return &ProductRepository{
		DB:    db,
		Redis: redis,
		ES:    es,
		Ctx:   ctx,
	}
}

func (r *ProductRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product

	// if true {
	// 	if err := r.DB.Preload(clause.Associations).Preload("Reviews", func(db *gorm.DB) *gorm.DB {
	// 		return db.Order("rating DESC").Limit(5)
	// 	}).Preload("Reviews.Images").First(&product, id).Error; err != nil {
	// 		return nil, fmt.Errorf("failed to get product from DB: %w", err)
	// 	}
	// 	return &product, nil
	// }

	// 1. Try to fetch from Redis
	redisKey := fmt.Sprintf("product:%d", id)
	val, err := r.Redis.Get(r.Ctx, redisKey).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(val), &product); err != nil {
			return nil, fmt.Errorf("failed to unmarshal product from Redis: %w", err)
		}
		return &product, nil
	}

	// 2. Try to fetch from Elasticsearch
	esRes, err := r.ES.Get().Index("products").Id(strconv.Itoa(int(id))).Do(r.Ctx)
	if err == nil && esRes.Found {
		if err = json.Unmarshal(esRes.Source, &product); err != nil {
			return nil, fmt.Errorf("failed to unmarshal product from Elasticsearch: %w", err)
		}
		// Update Redis with a random expiration time between 600 to 1200 minutes
		expiration := time.Duration(rand.Intn(601)+600) * time.Minute
		redisVal, err := json.Marshal(product)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal product for Redis: %w", err)
		}
		if err = r.Redis.Set(r.Ctx, redisKey, redisVal, expiration).Err(); err != nil {
			return nil, fmt.Errorf("failed to set product in Redis: %w", err)
		}
		return &product, nil
	}

	// 3. Fetch from MySQL
	if err = r.DB.Preload(clause.Associations).Preload("ProductAttributes.Value").
		Preload("ProductAttributes.Images").Preload("Reviews", func(db *gorm.DB) *gorm.DB {
		return db.Order("rating DESC").Limit(5)
	}).Preload("Reviews.Images").First(&product, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get product from DB: %w", err)
	}

	// Update Redis and Elasticsearch
	expiration := time.Duration(rand.Intn(601)+600) * time.Minute
	redisVal, err := json.Marshal(product)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal product for Redis: %w", err)
	}
	if err = r.Redis.Set(r.Ctx, redisKey, redisVal, expiration).Err(); err != nil {
		return nil, fmt.Errorf("failed to set product in Redis: %w", err)
	}
	_, err = r.ES.Index().Index("products").Id(strconv.Itoa(int(id))).BodyJson(product).Do(r.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to set product in Elasticsearch: %w", err)
	}

	return &product, nil
}
