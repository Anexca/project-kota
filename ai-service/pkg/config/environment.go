package config

import (
	"errors"
	"os"
	"strconv"
)

type Environment struct {
	ServerPort    string
	RedisPort     string
	RedisAddress  string
	RedisPassword string
	RedisDatabase int
}

func LoadEnvironment() (*Environment, error) {
	redisDbFromEnv := os.Getenv("REDIS_DATABASE")
	redisDatabase, err := strconv.Atoi(redisDbFromEnv)
	if err != nil {
		return nil, err
	}

	env := &Environment{
		ServerPort:    os.Getenv("PORT"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		RedisAddress:  os.Getenv("REDIS_ADDRESS"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDatabase: redisDatabase,
	}

	if env.ServerPort == "" || env.RedisPort == "" || env.RedisAddress == "" || env.RedisPassword == "" {
		return nil, errors.New("missing required environment variables")
	}

	return env, nil
}
