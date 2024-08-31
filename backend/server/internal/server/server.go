package server

import (
	commonConfig "common/config"
	"fmt"
	"net/http"
	"os"
	"server/internal/services"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int

	examService *services.ExamService
}

func InitServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	logger := commonConfig.SetupLogger()

	// examService := services.NewExamService()

	NewServer := &Server{
		port: port,
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
