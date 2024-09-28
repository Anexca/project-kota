package workers

import (
	"context"
	"log"

	"common/ent"

	"cloud.google.com/go/vertexai/genai"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"

	"ai-service/internal/services"
)

type Worker struct {
	cronHandler *cron.Cron
	examService *services.ExamService
}

func InitWorkers(genAiClient *genai.Client, redisClient *redis.Client, dbClient *ent.Client) *cron.Cron {
	c := cron.New()

	examService := services.NewExamService(genAiClient, redisClient, dbClient)

	worker := &Worker{
		cronHandler: c,
		examService: examService,
	}

	worker.RegisterWorkers()
	c.Start()
	return c
}

func (w *Worker) RegisterWorkers() {
	// _, err := w.cronHandler.AddFunc("*/5 * * * *", func() {
	_, err := w.cronHandler.AddFunc("0 3 * * *", func() {
		log.Println("Starting Worker Job for Populating Exam Question Cache")
		ctx := context.Background()
		err := w.examService.PopulateExamQuestionCache(ctx)
		if err != nil {
			log.Printf("Failed to generate questions: %v", err)
		}
	})

	if err != nil {
		log.Println(err)
	}
}
