package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheConfig struct {
	redisClient *redis.Client
	expiration  time.Duration
	timeout     time.Duration
	mutex       sync.Mutex
}

func NewCacheConfig(redisClient *redis.Client, expiration time.Duration, timeout time.Duration) *CacheConfig {
	return &CacheConfig{
		redisClient: redisClient,
		expiration:  expiration,
		timeout:     timeout,
	}
}

func (cc *CacheConfig) GetWithCache(key string, target interface{}, fetchFunc func() (interface{}, error)) error {
	if cc.redisClient == nil {
		return errors.New("Redis client is nil")
	}

	// 加锁保证并发安全
	cc.mutex.Lock()
	defer cc.mutex.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), cc.timeout)
	defer cancel()

	val, err := cc.redisClient.Get(ctx, key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			LogError.Printf("redis get error: %v", err)
			return fmt.Errorf("redis get error: %w", err)
		}

		// 缓存未命中，从数据库获取数据
		data, err := fetchFunc()
		if err != nil {
			LogError.Printf("fetch data error: %v", err)
			return fmt.Errorf("fetch data error: %w", err)
		}

		// 将数据存入缓存
		jsonData, err := json.Marshal(data)
		if err != nil {
			LogError.Printf("json marshal error: %v", err)
			return fmt.Errorf("json marshal error: %w", err)
		}

		if cc.expiration > 0 {
			err = cc.redisClient.Set(ctx, key, jsonData, cc.expiration).Err()
		} else {
			err = cc.redisClient.Set(ctx, key, jsonData, 0).Err()
		}
		if err != nil {
			LogError.Printf("redis set error: %v", err)
			return fmt.Errorf("redis set error: %w", err)
		}

		// 将数据传递给目标变量
		if err = json.Unmarshal(jsonData, target); err != nil {
			LogError.Printf("json unmarshal error: %v", err)
			return fmt.Errorf("json unmarshal error: %w", err)
		}

		LogInfo.Printf("get with cache: cache miss, key=%s", key)
		return nil
	}

	// 缓存命中，将数据传递给目标变量
	if err = json.Unmarshal([]byte(val), target); err != nil {
		LogError.Printf("json unmarshal error: %v", err)
		return fmt.Errorf("json unmarshal error: %w", err)
	}

	return nil
}
