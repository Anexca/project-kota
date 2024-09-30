package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisServiceInterface defines the contract for interacting with Redis.
type RedisServiceInterface interface {
	Store(ctx context.Context, key string, value string, expiry time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
	Health() map[string]string
	CheckRedisHealth(ctx context.Context, stats map[string]string) map[string]string
}

// RedisService is an implementation of RedisServiceInterface, which interacts with a Redis server.
type RedisService struct {
	client *redis.Client
}

// NewRedisService initializes a new RedisService instance.
func NewRedisService(redisClient *redis.Client) *RedisService {
	return &RedisService{
		client: redisClient,
	}
}

// Store stores a key-value pair in Redis with a specified expiration time.
func (r *RedisService) Store(ctx context.Context, key string, value string, expiry time.Duration) error {
	status := r.client.Set(ctx, key, value, expiry)
	if err := status.Err(); err != nil {
		return fmt.Errorf("failed to store key %s in Redis: %w", key, err)
	}
	return nil
}

// Get retrieves a value from Redis by key.
func (r *RedisService) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("key %s does not exist in Redis", key)
		}
		return "", fmt.Errorf("failed to get key %s from Redis: %w", key, err)
	}
	return val, nil
}

// Delete removes a key from Redis.
func (r *RedisService) Delete(ctx context.Context, key string) error {
	if err := r.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete key %s from Redis: %w", key, err)
	}
	return nil
}

// Health returns Redis health information, including a Redis ping.
func (r *RedisService) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Check Redis health and populate the stats map.
	stats = r.CheckRedisHealth(ctx, stats)

	return stats
}

// CheckRedisHealth pings the Redis server and checks whether it's healthy.
func (r *RedisService) CheckRedisHealth(ctx context.Context, stats map[string]string) map[string]string {
	// Ping the Redis server.
	pong, err := r.client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Redis is down: %v", err)
		stats["redis_status"] = "down"
	} else {
		stats["redis_status"] = "up"
		stats["redis_ping"] = pong
	}
	return stats
}
