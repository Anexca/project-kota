package services

import (
	"ai-service/pkg/constants"
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
	model := constants.FLASH_15

	if request.Model == "pro" {
		model = constants.PRO_15
	}

	return p.genAIService.GetContentStream(ctx, request.Prompt, model)
}
