package server

import (
	commonConfig "common/config"
	commonService "common/services"

	"common/ent"
	"fmt"
	"net/http"
	"os"
	"server/internal/services"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	port int

	examService  *services.ExamService
	redisService *commonService.RedisService
}

func InitServer(redisClient *redis.Client, dbClient *ent.Client) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	logger := commonConfig.SetupLogger()

	redisService := commonService.NewRedisService(redisClient)
	examService := services.NewExamService(redisClient, dbClient)

	NewServer := &Server{
		port:         port,
		examService:  examService,
		redisService: redisService,
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
