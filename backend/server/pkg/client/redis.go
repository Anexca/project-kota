package client

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"server/pkg/config"
)

func NewRedisClient(ctx context.Context) (*redis.Client, error) {
	env, err := config.LoadEnvironment()
	if err != nil {
		return nil, err
	}

	fullAddress := fmt.Sprintf("%s:%s", env.RedisAddress, env.RedisPort)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fullAddress,
		Password: env.RedisPassword,
		DB:       env.RedisDatabase,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("redis did not respond to Ping request %v", err)
	}

	log.Println("redis connected to client", env.RedisAddress, ",PING ->", pong)

	return redisClient, nil
}
