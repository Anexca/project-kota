package services

import (
	"ai-service/internal/repositories"
	"common/ent"
	util "common/utils"
	"context"
	"log"
	"sync"

	"cloud.google.com/go/vertexai/genai"
	"github.com/redis/go-redis/v9"
)

type ExamService struct {
	genAIService           *GenAIService
	redisService           *RedisService
	examRepository         *repositories.ExamRepository
	examCategoryRepository *repositories.ExamCategoryRepository
	examSettingRepository  *repositories.ExamSettingRepository
}

func NewExamService(genAIClient *genai.Client, redisClient *redis.Client, dbClient *ent.Client) *ExamService {
	genAIService := NewGenAIService(genAIClient)
	redisService := NewRedisService(redisClient)
	examRepository := repositories.NewExamRespository(dbClient)
	examCategoryRepository := repositories.NewExamCategoryRepository(dbClient)
	examSettingRepository := repositories.NewExamSettingRepository(dbClient)

	return &ExamService{
		genAIService:           genAIService,
		redisService:           redisService,
		examRepository:         examRepository,
		examCategoryRepository: examCategoryRepository,
		examSettingRepository:  examSettingRepository,
	}
}

const GEN_AI_MODEL = "gemini-1.5-pro"

func (q *ExamService) PopulateExamQuestionCache(ctx context.Context) error {
	wg := &sync.WaitGroup{}

	examCategories, err := q.examCategoryRepository.Get(ctx)
	if err != nil {
		return err
	}

	for _, cat := range examCategories {
		wg.Add(1) // Increment wait group for each category
		go func(cat *ent.ExamCategory) {
			defer wg.Done() // Decrement wait group after processing category

			exams, err := q.examRepository.GetByExamCategory(ctx, cat)
			if err != nil {
				log.Printf("Error getting exams for category %s: %v", cat.Name, err)
				return
			}

			for _, exam := range exams {
				go func(exam *ent.Exam) {
					examSetting, err := q.examSettingRepository.GetByExam(ctx, exam)
					if err != nil {
						log.Printf("Error getting exam setting for exam %s: %v", exam.Name, err)
						return
					}

					response, err := q.genAIService.GetContentStream(ctx, examSetting.AiPrompt, GEN_AI_MODEL)
					if err != nil {
						log.Printf("Error generating content for exam %s: %v", exam.Name, err)
						return
					}

					uid := util.GenerateUUID()
					q.redisService.Store(ctx, uid, response)
					log.Printf("Cached response for exam %s with uid %s", exam.Name, uid)
				}(exam)
			}
		}(cat)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	return nil
}
