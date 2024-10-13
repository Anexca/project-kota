package config

import (
	"errors"
	"os"
	"strconv"
)

type Environment struct {
	IsProduction             bool
	ServerPort               string
	DatabaseHost             string
	DatabasePort             string
	DatabaseName             string
	DatabaseUser             string
	DatabasePassword         string
	RedisPort                string
	RedisAddress             string
	RedisPassword            string
	RedisDatabase            int
	GoogleCloudProjectId     string
	GoogleCloudProjectRegion string
	ServerAccessKey          string
	ServerAdminKey           string
}

func LoadEnvironment() (*Environment, error) {
	redisDbFromEnv := os.Getenv("REDIS_DATABASE")
	redisDatabase, err := strconv.Atoi(redisDbFromEnv)
	if err != nil {
		return nil, err
	}

	env := &Environment{
		ServerPort:               os.Getenv("PORT"),
		RedisPort:                os.Getenv("REDIS_PORT"),
		DatabaseHost:             os.Getenv("DB_HOST"),
		DatabasePort:             os.Getenv("DB_PORT"),
		DatabaseName:             os.Getenv("DB_NAME"),
		DatabaseUser:             os.Getenv("DB_USER"),
		DatabasePassword:         os.Getenv("DB_PASSWORD"),
		RedisAddress:             os.Getenv("REDIS_ADDRESS"),
		RedisPassword:            os.Getenv("REDIS_PASSWORD"),
		RedisDatabase:            redisDatabase,
		GoogleCloudProjectId:     os.Getenv("GCLOUD_PROJECT_ID"),
		GoogleCloudProjectRegion: os.Getenv("GCLOUD_PROJECT_REGION"),
		ServerAccessKey:          os.Getenv("SERVER_ACCESS_KEY"),
		ServerAdminKey:           os.Getenv("SERVER_ADMIN_KEY"),
		IsProduction:             os.Getenv("ENV") == "production",
	}

	if env.ServerPort == "" {
		return nil, errors.New("missing SERVER_PORT environment variable")
	}

	if env.RedisPort == "" || env.RedisAddress == "" || env.RedisPassword == "" {
		return nil, errors.New("missing Redis environment variables")
	}

	if env.DatabaseHost == "" || env.DatabasePort == "" || env.DatabaseName == "" || env.DatabaseUser == "" || env.DatabasePassword == "" {
		return nil, errors.New("missing Database environment variables")
	}

	if env.GoogleCloudProjectId == "" || env.GoogleCloudProjectRegion == "" {
		return nil, errors.New("missing Google Cloud environment variables")
	}

	if env.ServerAccessKey == "" {
		return nil, errors.New("missing Server Access Key environment variable")
	}

	if env.ServerAdminKey == "" {
		return nil, errors.New("missing Server Admin Key environment variable")
	}

	return env, nil
}
