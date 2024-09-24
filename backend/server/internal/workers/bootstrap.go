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
	examService *services.ExamGenerationService
}

func InitWorkers(redisClient *redis.Client, dbClient *ent.Client) *cron.Cron {
	c := cron.New()

	examService := services.NewExamGenerationService(redisClient, dbClient)

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
		// w.cronHandler.AddFunc("0 4 * * *", func() {
		log.Println("Starting Worker Job for Adding Descriptive Question in Database")

		// err := w.AddDescriptiveQuestionsInDatabase()
		// if err != nil {
		// 	log.Printf("Failed to Add Descriptive Question in Database: %v", err)
		// }

		err := w.AddMcqExamsInDatabase()
		if err != nil {
			log.Printf("Failed to Add MCQ Exam in Database: %v", err)
		}

		log.Println("Finished Worker Job for Adding Descriptive Question in Database")
	})

	// w.cronHandler.AddFunc("*/1 * * * *", func() {
	w.cronHandler.AddFunc("0 0 * * 0", func() {
		log.Println("Starting Worker Job for Creating Descriptive Open Questions")

		err := w.MarkDescriptiveQuestionsAsOpenInDatabase()
		if err != nil {
			log.Printf("Failed to Create Descriptive Open Questions: %v", err)
			return
		}

		log.Println("Finished Worker Job for Creating Descriptive Open Questions")
	})
}
