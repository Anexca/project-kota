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
	"github.com/razorpay/razorpay-go"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	port int

	authService           *services.AuthService
	userService           *services.UserService
	redisService          *commonService.RedisService
	paymentService        *services.PaymentService
	subscriptionService   *services.SubscriptionService
	examAttemptService    *services.ExamAttemptService
	examAssesmentService  *services.ExamAssesmentService
	examGenerationService *services.ExamGenerationService
}

func InitServer(redisClient *redis.Client, dbClient *ent.Client, supabaseClient *supabase.Client, paymentClient *razorpay.Client) *http.Server {
	logger := commonConfig.SetupLogger()
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	userService := services.NewUserService(dbClient)
	authService := services.NewAuthService(supabaseClient)
	redisService := commonService.NewRedisService(redisClient)
	paymentService := services.NewPaymentService(paymentClient)
	subscriptionService := services.NewSubscriptionService(dbClient, paymentClient)
	examAttemptService := services.NewExamAttemptService(dbClient)
	examAssesmentService := services.NewExamAssesmentService(redisClient, dbClient)
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
