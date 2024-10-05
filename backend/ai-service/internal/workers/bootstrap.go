package workers

import (
	"common/ent"
	"common/repositories"

	"cloud.google.com/go/vertexai/genai"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"

	"ai-service/internal/services"
	commonService "common/services"
)

// Worker handles cron jobs and the associated services
type Worker struct {
	cronHandler *cron.Cron
	examService *services.ExamService
}

// InitWorkers initializes the cron jobs and starts the workers
func InitWorkers(genAiClient *genai.Client, redisClient *redis.Client, dbClient *ent.Client) *cron.Cron {
	c := cron.New()

	// Inject real implementations into the ExamService
	examService := services.NewExamService(
		services.NewGenAIService(genAiClient),
		commonService.NewRedisService(redisClient),
		repositories.NewExamRepository(dbClient),
		repositories.NewExamCategoryRepository(dbClient),
		repositories.NewExamSettingRepository(dbClient),
		repositories.NewCachedExamRepository(dbClient),
	)

	// Create the worker with the initialized examService
	worker := &Worker{
		cronHandler: c,
		examService: examService,
	}

	// Register the cron jobs
	worker.RegisterWorkers()

	// Start the cron scheduler
	c.Start()

	return c
}

// RegisterWorkers registers the cron jobs and their associated tasks
func (w *Worker) RegisterWorkers() {
	// Schedule the cron job to run daily at 3 AM
	// _, err := w.cronHandler.AddFunc("0 3 * * *", func() {
	// 	log.Println("Starting Worker Job for Populating Exam Question Cache")
	// 	ctx := context.Background()
	// 	err := w.examService.PopulateExamQuestionCache(ctx)
	// 	if err != nil {
	// 		log.Printf("Failed to generate questions: %v", err)
	// 	}
	// })

	// if err != nil {
	// 	log.Printf("Error registering worker: %v", err)
	// }
}
