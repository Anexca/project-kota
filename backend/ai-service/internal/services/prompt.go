package services

import (
	"context"
	"strings"

	"cloud.google.com/go/vertexai/genai"

	commonConstants "common/constants"

	"ai-service/pkg/models"
)

// Define an interface for GenAIService to make it mockable
type GenAIServiceInterface interface {
	GetContentStream(ctx context.Context, prompt string, model commonConstants.GenAiModel) (string, error)
}

type PromptService struct {
	genAIService GenAIServiceInterface
}

// Factory method for production use, using *genai.Client
func NewPromptService(genAiClient *genai.Client) *PromptService {
	genAIService := NewGenAIService(genAiClient)

	return &PromptService{
		genAIService: genAIService,
	}
}

func (p *PromptService) GetPromptResults(ctx context.Context, request *models.GetPromptResultsRequest) (string, error) {
	model := commonConstants.FLASH_15

	if strings.Contains(request.Model, "gemini-1.5-pro") {
		model = commonConstants.PRO_15
	}

	return p.genAIService.GetContentStream(ctx, request.Prompt, model)
}
