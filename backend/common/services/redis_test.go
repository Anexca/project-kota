package services_test

import (
	"context"
	"testing"
	"time"

	"common/services"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMiniredis() (*miniredis.Miniredis, *redis.Client) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err) // Handle error appropriately in production code
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(), // Use the address of the miniredis server
	})

	return mr, client
}

func TestRedisService_StoreAndGet(t *testing.T) {
	mr, client := setupMiniredis()
	defer mr.Close() // Close miniredis server after tests

	redisService := services.NewRedisService(client)
	ctx := context.Background()

	key := "test_key"
	value := "test_value"
	expiry := time.Minute

	err := redisService.Store(ctx, key, value, expiry)
	require.NoError(t, err)

	storedValue, err := redisService.Get(ctx, key)
	require.NoError(t, err)
	assert.Equal(t, value, storedValue)
}

func TestRedisService_Get_NonExistentKey(t *testing.T) {
	mr, client := setupMiniredis()
	defer mr.Close()

	redisService := services.NewRedisService(client)
	ctx := context.Background()

	_, err := redisService.Get(ctx, "non_existent_key")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "does not exist")
}

func TestRedisService_Delete(t *testing.T) {
	mr, client := setupMiniredis()
	defer mr.Close()

	redisService := services.NewRedisService(client)
	ctx := context.Background()

	key := "test_delete_key"
	value := "value_to_delete"
	_ = redisService.Store(ctx, key, value, time.Minute)

	err := redisService.Delete(ctx, key)
	require.NoError(t, err)

	_, err = redisService.Get(ctx, key)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "does not exist")
}

func TestRedisService_Health(t *testing.T) {
	mr, client := setupMiniredis()
	defer mr.Close()

	redisService := services.NewRedisService(client)
	stats := redisService.Health()

	assert.Equal(t, "up", stats["redis_status"])
	assert.Contains(t, stats, "redis_ping")
}
