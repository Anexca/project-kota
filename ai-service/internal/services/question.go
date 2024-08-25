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

type MultipleChoiceQuestion struct {
	Question    string   `json:"question"`
	Options     []string `json:"options"`
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

func (q *QuestionService) GenerateMCQsForSubject(ctx context.Context, examName, subject string, numberOfQuestions int) (*[]MultipleChoiceQuestion, error) {
	questionKey := "QUESTION"
	var formattedQuestions []MultipleChoiceQuestion

	cachedQuestions, err := q.redisService.Get(ctx, questionKey)
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if err != redis.Nil {
		err = json.Unmarshal([]byte(cachedQuestions), &formattedQuestions)

		if err != nil {
			return nil, err
		}

		return &formattedQuestions, nil
	}

	prompt := fmt.Sprintf(`Generate a JSON array containing %d multiple-choice questions (MCQs) for the %s subject in %s exam. 
            Each MCQ should have a "question", "options", "answer", and "explanation" field. 
            The JSON output should be formatted as a single-line string without any extra whitespace or formatting`, numberOfQuestions, subject, examName)

	questions, err := q.genAIService.GetContentStream(ctx, prompt, GEN_AI_MODEL)
	if err != nil {
		return nil, err
	}

	err = q.redisService.Store(ctx, questionKey, questions)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(questions), &formattedQuestions)
	if err != nil {
		return nil, err
	}

	return &formattedQuestions, nil
}
