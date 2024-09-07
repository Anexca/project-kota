package config

import (
	"errors"
	"os"
	"strconv"
)

type Environment struct {
	IsProduction       bool
	CorsAllowedOrigin  string
	ServerPort         string
	DatabaseHost       string
	DatabasePort       string
	DatabaseName       string
	DatabaseUser       string
	DatabasePassword   string
	RedisPort          string
	RedisAddress       string
	RedisPassword      string
	RedisDatabase      int
	SupabaseUrl        string
	SupabaseKey        string
	AIServiceAccessKey string
	AIServiceUrl       string
}

func LoadEnvironment() (*Environment, error) {
	redisDbFromEnv := os.Getenv("REDIS_DATABASE")
	redisDatabase, err := strconv.Atoi(redisDbFromEnv)
	if err != nil {
		return nil, err
	}

	env := &Environment{
		ServerPort:         os.Getenv("PORT"),
		RedisPort:          os.Getenv("REDIS_PORT"),
		DatabaseHost:       os.Getenv("DB_HOST"),
		DatabasePort:       os.Getenv("DB_PORT"),
		DatabaseName:       os.Getenv("DB_NAME"),
		DatabaseUser:       os.Getenv("DB_USER"),
		DatabasePassword:   os.Getenv("DB_PASSWORD"),
		RedisAddress:       os.Getenv("REDIS_ADDRESS"),
		RedisPassword:      os.Getenv("REDIS_PASSWORD"),
		RedisDatabase:      redisDatabase,
		SupabaseUrl:        os.Getenv("SUPABASE_URL"),
		SupabaseKey:        os.Getenv("SUPABASE_KEY"),
		AIServiceAccessKey: os.Getenv("AI_SERVICE_ACCESS_KEY"),
		AIServiceUrl:       os.Getenv("AI_SERVICE_URL"),
		IsProduction:       os.Getenv("ENV") == "production",
		CorsAllowedOrigin:  os.Getenv("CORS_ALLOWED_ORIGIN"),
	}

	if env.ServerPort == "" || env.CorsAllowedOrigin == "" {
		return nil, errors.New("missing SERVER_PORT or Allowed CORS environment variable")
	}

	if env.RedisPort == "" || env.RedisAddress == "" || env.RedisPassword == "" {
		return nil, errors.New("missing Redis environment variables")
	}

	if env.DatabaseHost == "" || env.DatabasePort == "" || env.DatabaseName == "" || env.DatabaseUser == "" || env.DatabasePassword == "" {
		return nil, errors.New("missing Database environment variables")
	}

	if env.SupabaseKey == "" || env.SupabaseUrl == "" {
		return nil, errors.New("missing Supabase environment variables")
	}

	if env.AIServiceAccessKey == "" || env.AIServiceUrl == "" {
		return nil, errors.New("missing AI Service environment variables")
	}

	return env, nil
}
