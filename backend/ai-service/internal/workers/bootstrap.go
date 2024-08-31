package workers

import (
	"ai-service/internal/services"
	"common/ent"
	"log"

	"cloud.google.com/go/vertexai/genai"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
)

type Worker struct {
	cronHandler     *cron.Cron
	questionService *services.QuestionService
}

func InitWorkers(genAiClient *genai.Client, redisClient *redis.Client, dbClient *ent.Client) *cron.Cron {
	c := cron.New()

	questionService := services.NewQuestionService(genAiClient, redisClient, dbClient)

	worker := &Worker{
		cronHandler:     c,
		questionService: questionService,
	}

	worker.RegisterWorkers()
	c.Start()
	return c
}

func (w *Worker) RegisterWorkers() {
	w.cronHandler.AddFunc("*/1 * * * *", func() {
		// ctx := context.Background()
		log.Println("Every Minute")
		// _, err := w.questionService.GenerateDescriptiveQuestions(ctx, "IBPS PO Mains", 10)
		// if err != nil {
		// 	log.Printf("Failed to generate questions: %v", err)
		// }
	})
}
