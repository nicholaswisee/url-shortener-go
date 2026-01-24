package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"url-shortener-go/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis(cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       0, // use default DB
	})

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	RedisClient = client
	log.Println("✅ Connected to Redis")
	return client, nil
}

// CacheURL stores a URL mapping in Redis with expiration
func CacheURL(ctx context.Context, shortCode, originalURL string, expiration time.Duration) error {
	return RedisClient.Set(ctx, shortCode, originalURL, expiration).Err()
}

// GetCachedURL retrieves a URL from Redis cache
func GetCachedURL(ctx context.Context, shortCode string) (string, error) {
	return RedisClient.Get(ctx, shortCode).Result()
}

// DeleteCachedURL removes a URL from Redis cache
func DeleteCachedURL(ctx context.Context, shortCode string) error {
	return RedisClient.Del(ctx, shortCode).Err()
}
