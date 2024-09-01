package workers

import (
	"common/ent"
	"log"
	"server/internal/services"

	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
)

type Worker struct {
	cronHandler *cron.Cron
	examService *services.ExamService
}

func InitWorkers(redisClient *redis.Client, dbClient *ent.Client) *cron.Cron {
	c := cron.New()

	examService := services.NewExamService(redisClient, dbClient)

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
		err := w.AddDescriptiveQuestionsInDatabase()
		if err != nil {
			log.Printf("Failed to generate questions: %v", err)
		}
	})
}
