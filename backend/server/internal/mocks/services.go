package mocks

import (
	"common/constants"
	"common/ent"
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

func (m *MockAccessService) GetAccessibleExamsForUser(ctx context.Context, exams []*ent.Exam, userId string) ([]*ent.Exam, error) {
	args := m.Called(ctx, exams, userId)
	return args.Get(0).([]*ent.Exam), args.Error(1)
}

// Mock for ExamGenerationServiceInterface
type MockExamGenerationService struct {
	mock.Mock
}

func (m *MockExamGenerationService) GetGeneratedExamById(ctx context.Context, generatedExamId int, userId string, isOpen bool) (*models.GeneratedExamOverview, error) {
	args := m.Called(ctx, generatedExamId, userId, isOpen)
	return args.Get(0).(*models.GeneratedExamOverview), args.Error(1)
}

// MockExamAssesmentService is a mock implementation of ExamAssesmentServiceInterface
type MockExamAssesmentService struct {
	mock.Mock
}

func (m *MockExamAssesmentService) StartNewDescriptiveAssesment(ctx context.Context, generatedExamId int, attempt *ent.ExamAttempt, request *models.DescriptiveExamAssesmentRequest, userId string, isOpen bool) (*models.AssessmentDetails, error) {
	args := m.Called(ctx, generatedExamId, attempt, request, userId, isOpen)
	return args.Get(0).(*models.AssessmentDetails), args.Error(1)
}

func (m *MockExamAssesmentService) GetAssesmentById(ctx context.Context, assessmentId int, userId string) (*models.AssessmentDetails, error) {
	args := m.Called(ctx, assessmentId, userId)
	return args.Get(0).(*models.AssessmentDetails), args.Error(1)
}

func (m *MockExamAssesmentService) GetExamAssessments(ctx context.Context, generatedExamId int, userId string) ([]models.AssessmentDetails, error) {
	args := m.Called(ctx, generatedExamId, userId)
	return args.Get(0).([]models.AssessmentDetails), args.Error(1)
}

func (m *MockExamAssesmentService) AssessDescriptiveExam(ctx context.Context, generatedExamId int, assessmentId int, content string, userId string, isOpen bool) {
	m.Called(ctx, generatedExamId, assessmentId, content, userId, isOpen)
}
