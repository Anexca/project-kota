package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/vertexai/genai"
	_ "github.com/joho/godotenv/autoload"

	"ai-service/internal/database"
	"ai-service/internal/services"
)

type Server struct {
	port            int
	questionService *services.QuestionService
	db              database.Service
}

func InitServer(genAiClient *genai.Client) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	questionService := services.NewQuestionService(genAiClient)

	NewServer := &Server{
		port:            port,
		questionService: questionService,
		db:              database.New(),
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
