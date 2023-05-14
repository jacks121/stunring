package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %s", err)
	}

	log.Println("Connected to Redis server.")

	return redisClient, nil
}
