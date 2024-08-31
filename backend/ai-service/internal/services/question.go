package services

import (
	"ai-service/internal/repositories"
	"common/ent"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/vertexai/genai"
	"github.com/redis/go-redis/v9"
)

type QuestionService struct {
	genAIService           *GenAIService
	redisService           *RedisService
	examRepository         *repositories.ExamRepository
	examCategoryRepository *repositories.ExamCategoryRepository
}

type QuestionWithExplanation struct {
	Type        string   `json:"type"` // Added field to differentiate MCQ and NVQ
	Question    string   `json:"question"`
	Options     []string `json:"options,omitempty"` // Optional for NVQs
	Answer      string   `json:"answer"`
	Explanation string   `json:"explanation"`
}

type DescriptiveQuestion struct {
	Type  string   `json:"type"`
	Topic string   `json:"topic"`
	Hint  []string `json:"hint,omitempty"`
}

func NewQuestionService(genAIClient *genai.Client, redisClient *redis.Client, dbClient *ent.Client) *QuestionService {
	genAIService := NewGenAIService(genAIClient)
	redisService := NewRedisService(redisClient)
	examRepository := repositories.NewExamRespository(dbClient)
	examCategoryRepository := repositories.NewExamCategoryRepository(dbClient)

	return &QuestionService{
		genAIService:           genAIService,
		redisService:           redisService,
		examRepository:         examRepository,
		examCategoryRepository: examCategoryRepository,
	}
}

const GEN_AI_MODEL = "gemini-1.5-pro"
const REDIS_CACHE_PREFIX = "QUESTIONS"

func (q *QuestionService) GenerateQuestions(ctx context.Context, questionType, examName, subject string, numberOfQuestions int) ([]QuestionWithExplanation, error) {
	cachedQuestionKey := generateCacheKey(questionType, examName, subject, numberOfQuestions)
	var formattedQuestions []QuestionWithExplanation

	// Check cache for existing questions
	cachedQuestions, err := q.redisService.Get(ctx, cachedQuestionKey)
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if err != redis.Nil {
		err = json.Unmarshal([]byte(cachedQuestions), &formattedQuestions)

		if err != nil {
			return nil, err
		}

		return formattedQuestions, nil
	}

	// Generate questions using GenAI if not found in cache
	prompt := fmt.Sprintf(`Generate a JSON array containing %d %s questions for the %s subject in %s exam. 
                            Each question should have a "type", "question", "answer", "explanation" and "options", field.
                            The JSON output should be formatted as a single-line string without any extra whitespace or formatting.`,
		numberOfQuestions, questionType, subject, examName)

	questions, err := q.genAIService.GetContentStream(ctx, prompt, GEN_AI_MODEL)
	if err != nil {
		return nil, err
	}

	// Unmarshall and store questions in cache
	err = json.Unmarshal([]byte(questions), &formattedQuestions)
	if err != nil {
		return nil, err
	}

	q.redisService.Store(ctx, cachedQuestionKey, questions)

	return formattedQuestions, nil
}

func (q *QuestionService) GenerateDescriptiveQuestions(ctx context.Context, examName string, numberOfQuestions int) (any, error) {
	var formattedQuestions []DescriptiveQuestion

	examCategories, err := q.examCategoryRepository.Get(ctx)
	if err != nil {
		return nil, err
	}
	for _, cat := range examCategories {
		exam, err := q.examRepository.GetByExamCategory(ctx, cat)
		if err != nil {
			return nil, err
		}

		log.Println(exam)
	}

	// prompt := fmt.Sprintf(`Generate a JSON array containing %d Descriptive questions for the %s exam.
	// 						Essay should be a one sentence topic, letter writing should be formal.
	//                         Each question should have a "type" should be either formal letter or essay, "topic" should be the question itself, "hint" should be an array of hints for topic.
	//                         The JSON output should be a single-line string without any extra formatting.`,
	// 	numberOfQuestions, examName)

	// questions, err := q.genAIService.GetContentStream(ctx, prompt, GEN_AI_MODEL)

	// if err != nil {
	// 	return nil, err
	// }
	// // Unmarshall and store questions in cache
	// err = json.Unmarshal([]byte(questions), &formattedQuestions)
	// if err != nil {
	// 	return nil, err
	// }

	// uid := util.GenerateUUID()
	// q.redisService.Store(ctx, uid, questions)
	// exams, _ := q.examRepository.GetByExamCategory(ctx, )

	return formattedQuestions, nil
}

func generateCacheKey(questionType, examName, subject string, numberOfQuestions int) string {
	return fmt.Sprintf("%s_%s_%s_%s_%d", REDIS_CACHE_PREFIX, questionType, examName, subject, numberOfQuestions)
}
