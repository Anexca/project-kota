package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	commonConfig "common/config"
	"common/ent"
	commonService "common/services"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nedpals/supabase-go"
	"github.com/redis/go-redis/v9"
	"server/internal/services"
)

type Server struct {
	port int

	authService           *services.AuthService
	userService           *services.UserService
	redisService          *commonService.RedisService
	paymentService        *services.PaymentService
	examCategoryService   *services.ExamCategoryService
	subscriptionService   *services.SubscriptionService
	examAttemptService    *services.ExamAttemptService
	examAssesmentService  *services.ExamAssesmentService
	examGenerationService *services.ExamGenerationService
}

func InitServer(redisClient *redis.Client, dbClient *ent.Client, supabaseClient *supabase.Client) *http.Server {
	logger := commonConfig.SetupLogger()
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	authService := services.NewAuthService(supabaseClient)
	redisService := commonService.NewRedisService(redisClient)
	paymentService := services.NewPaymentService()
	examAttemptService := services.NewExamAttemptService(dbClient)
	userService := services.NewUserService(dbClient)
	examCategoryService := services.NewExamCategoryService(dbClient)
	examAssesmentService := services.NewExamAssesmentService(redisClient, dbClient)
	subscriptionService := services.NewSubscriptionService(dbClient)
	examGenerationService := services.NewExamGenerationService(redisClient, dbClient)

	NewServer := &Server{
		port:                  port,
		userService:           userService,
		authService:           authService,
		redisService:          redisService,
		paymentService:        paymentService,
		examAttemptService:    examAttemptService,
		examAssesmentService:  examAssesmentService,
		examGenerationService: examGenerationService,
		subscriptionService:   subscriptionService,
		examCategoryService:   examCategoryService,
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
