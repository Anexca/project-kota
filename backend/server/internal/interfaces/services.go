package interfaces

import (
	"common/constants"
	"context"
	"server/pkg/models"
)

// PromptServiceInterface defines the contract for PromptService
type PromptServiceInterface interface {
	GetPromptResult(ctx context.Context, prompt string, model constants.GenAiModel) (string, error)
}

// ExamGenerationServiceInterface defines the contract for ExamGenerationService
type ExamGenerationServiceInterface interface {
	GetGeneratedExamById(ctx context.Context, generatedExamId int, userId string, isOpen bool) (*models.GeneratedExamOverview, error)
}
