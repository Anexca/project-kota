package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/vertexai/genai"
	"github.com/redis/go-redis/v9"

	_ "github.com/joho/godotenv/autoload"

	commonConfig "common/config"
	"common/ent"
	commonService "common/services"

	"ai-service/internal/services"
)

type Server struct {
	port          int
	promptService *services.PromptService
	redisService  *commonService.RedisService
	examService   *services.ExamService
}

func InitServer(genAiClient *genai.Client, redisClient *redis.Client, dbClient *ent.Client) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	logger := commonConfig.SetupLogger()

	promptService := services.NewPromptService(genAiClient)
	examService := services.InitExamService(genAiClient, redisClient, dbClient)
	redisService := commonService.NewRedisService(redisClient)

	NewServer := &Server{
		port:          port,
		redisService:  redisService,
		promptService: promptService,
		examService:   examService,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  100 * time.Minute,
		WriteTimeout: 100 * time.Minute,
		ErrorLog:     logger,
	}

	return server
}
