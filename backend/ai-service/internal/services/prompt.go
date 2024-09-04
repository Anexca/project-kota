package services

import (
	commonConstants "common/constants"
	"strings"

	"ai-service/pkg/models"
	"context"

	"cloud.google.com/go/vertexai/genai"
)

type PromptService struct {
	genAIService *GenAIService
}

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
