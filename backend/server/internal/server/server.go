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
	"github.com/nedpals/supabase-go"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	port int

	examService  *services.ExamService
	authService  *services.AuthService
	redisService *commonService.RedisService
}

func InitServer(redisClient *redis.Client, dbClient *ent.Client, supabaseClient *supabase.Client) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	logger := commonConfig.SetupLogger()

	examService := services.NewExamService(redisClient, dbClient)
	authService := services.NewAuthService(supabaseClient)
	redisService := commonService.NewRedisService(redisClient)

	NewServer := &Server{
		port:         port,
		examService:  examService,
		redisService: redisService,
		authService:  authService,
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
