package workers

import (
	"ai-service/internal/services"
	"common/ent"
	"context"
	"log"

	"cloud.google.com/go/vertexai/genai"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
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
	w.cronHandler.AddFunc("*/1 * * * *", func() {
		log.Println("Starting Worker Job for Populating Exam Question Cache")
		ctx := context.Background()
		err := w.examService.PopulateExamQuestionCache(ctx)
		if err != nil {
			log.Printf("Failed to generate questions: %v", err)
		}
	})
}
