package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/vertexai/genai"
	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"

	"ai-service/internal/services"
)

type Server struct {
	port            int
	questionService *services.QuestionService
	redisService    *services.RedisService
}

func InitServer(genAiClient *genai.Client, redisClient *redis.Client) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	questionService := services.NewQuestionService(genAiClient)
	redisService := services.NewRedisService(redisClient)

	NewServer := &Server{
		port:            port,
		questionService: questionService,
		redisService:    redisService,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
