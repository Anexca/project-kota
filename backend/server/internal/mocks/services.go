package mocks

import (
	"common/constants"
	"context"
	"server/pkg/models"

	"github.com/stretchr/testify/mock"
)

// Mock for AccessServiceInterface
type MockAccessService struct {
	mock.Mock
}

func (m *MockAccessService) UserHasAccessToExam(ctx context.Context, examId int, userId string) (bool, error) {
	args := m.Called(ctx, examId, userId)
	return args.Bool(0), args.Error(1)
}

// Mock for PromptServiceInterface
type MockPromptService struct {
	mock.Mock
}

func (m *MockPromptService) GetPromptResult(ctx context.Context, prompt string, model constants.GenAiModel) (string, error) {
	args := m.Called(ctx, prompt, model)
	return args.String(0), args.Error(1)
}

// Mock for ExamGenerationServiceInterface
type MockExamGenerationService struct {
	mock.Mock
}

func (m *MockExamGenerationService) GetGeneratedExamById(ctx context.Context, generatedExamId int, userId string, isOpen bool) (*models.GeneratedExamOverview, error) {
	args := m.Called(ctx, generatedExamId, userId, isOpen)
	return args.Get(0).(*models.GeneratedExamOverview), args.Error(1)
}
