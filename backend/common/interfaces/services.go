package interfaces

import (
	"context"
	"time"
)

// ProfanityServiceInterface defines the contract for the profanity detection service
type ProfanityServiceInterface interface {
	IsProfane(s string) bool
}

// RedisServiceInterface defines the contract for interacting with Redis.
type RedisServiceInterface interface {
	Store(ctx context.Context, key string, value string, expiry time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
	Health() map[string]string
	CheckRedisHealth(ctx context.Context, stats map[string]string) map[string]string
}
