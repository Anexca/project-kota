package server

import (
	"ai-service/internal/services"
	commonConfig "common/config"
	commonService "common/services"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/vertexai/genai"
	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	port          int
	promptService *services.PromptService
	redisService  *commonService.RedisService
}

func InitServer(genAiClient *genai.Client, redisClient *redis.Client) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	logger := commonConfig.SetupLogger()

	promptService := services.NewPromptService(genAiClient)
	redisService := commonService.NewRedisService(redisClient)

	NewServer := &Server{
		port:          port,
		redisService:  redisService,
		promptService: promptService,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorLog:     logger,
	}

	return server
}
