package services

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/vertexai/genai"
	"github.com/redis/go-redis/v9"
)

type QuestionService struct {
	genAIService *GenAIService
	redisService *RedisService
}

type Question struct {
	Type        string   `json:"type"` // Added field to differentiate MCQ and NVQ
	Question    string   `json:"question"`
	Options     []string `json:"options,omitempty"` // Optional for NVQs
	Answer      string   `json:"answer"`
	Explanation string   `json:"explanation"`
}

func NewQuestionService(genAIClient *genai.Client, redisClient *redis.Client) *QuestionService {
	genAIService := NewGenAIService(genAIClient)
	redisService := NewRedisService(redisClient)

	return &QuestionService{
		genAIService: genAIService,
		redisService: redisService,
	}
}

const GEN_AI_MODEL = "gemini-1.5-flash"
const REDIS_CACHE_PREFIX = "QUESTIONS"

func (q *QuestionService) GenerateQuestions(ctx context.Context, questionType, examName, subject string, numberOfQuestions int) ([]Question, error) {
	cachedQuestionKey := generateCacheKey(questionType, examName, subject, numberOfQuestions)
	var formattedQuestions []Question

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

func generateCacheKey(questionType, examName, subject string, numberOfQuestions int) string {
	return fmt.Sprintf("%s_%s_%s_%s_%d", REDIS_CACHE_PREFIX, questionType, examName, subject, numberOfQuestions)
}
