package services

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/vertexai/genai"
)

type QuestionService struct {
	genAIService *GenAIService
}

type MultipleChoiceQuestion struct {
	Question    string   `json:"question"`
	Options     []string `json:"options"`
	Answer      string   `json:"answer"`
	Explanation string   `json:"explanation"`
}

func NewQuestionService(client *genai.Client) *QuestionService {
	genAIService := NewGenAIService(client)

	return &QuestionService{
		genAIService: genAIService,
	}
}

const GEN_AI_MODEL = "gemini-1.5-flash"

func (q *QuestionService) GenerateMCQsForSubject(ctx context.Context, examName, subject string, numberOfQuestions int) (*[]MultipleChoiceQuestion, error) {
	prompt := fmt.Sprintf(`Generate a JSON array containing %d multiple-choice questions (MCQs) for the %s subject in %s exam. 
            Each MCQ should have a "question", "options", "answer", and "explanation" field. 
            The JSON output should be formatted as a single-line string without any extra whitespace or formatting`, numberOfQuestions, subject, examName)

	questions, err := q.genAIService.GetContentStream(ctx, prompt, GEN_AI_MODEL)
	if err != nil {
		return nil, err
	}

	var formattedQuestions []MultipleChoiceQuestion
	err = json.Unmarshal([]byte(questions), &formattedQuestions)
	if err != nil {
		return nil, err
	}

	return &formattedQuestions, nil
}
