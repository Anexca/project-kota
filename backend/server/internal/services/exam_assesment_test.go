package services_test

import (
	"context"
	"testing"

	"common/constants"
	"common/ent"
	"common/ent/examassesment"
	"common/repositories"
	"server/internal/services"
	"server/pkg/models"

	"github.com/stretchr/testify/assert"
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

// Mock for ProfanityServiceInterface
type MockProfanityService struct {
	mock.Mock
}

func (m *MockProfanityService) IsProfane(s string) bool {
	args := m.Called(s)
	return args.Bool(0)
}

// Mock for GeneratedExamRepositoryInterface
type MockGeneratedExamRepository struct {
	mock.Mock
}

func (m *MockGeneratedExamRepository) GetOpenById(ctx context.Context, id int, isOpen bool) (*ent.GeneratedExam, error) {
	args := m.Called(ctx, id, isOpen)
	return args.Get(0).(*ent.GeneratedExam), args.Error(1)
}

// Mock for ExamAttemptRepositoryInterface
type MockExamAttemptRepository struct {
	mock.Mock
}

func (m *MockExamAttemptRepository) GetById(ctx context.Context, attemptId int, userId string) (*ent.ExamAttempt, error) {
	args := m.Called(ctx, attemptId, userId)
	return args.Get(0).(*ent.ExamAttempt), args.Error(1)
}

func (m *MockExamAttemptRepository) GetByUserId(ctx context.Context, userId string) ([]*ent.ExamAttempt, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).([]*ent.ExamAttempt), args.Error(1)
}

func (m *MockExamAttemptRepository) GetByExam(ctx context.Context, generatedExamId int, userId string) ([]*ent.ExamAttempt, error) {
	args := m.Called(ctx, generatedExamId, userId)
	return args.Get(0).([]*ent.ExamAttempt), args.Error(1)
}

func (m *MockExamAttemptRepository) Create(ctx context.Context, currentAttempt int, generatedExamId int, userId string) (*ent.ExamAttempt, error) {
	args := m.Called(ctx, currentAttempt, generatedExamId, userId)
	return args.Get(0).(*ent.ExamAttempt), args.Error(1)
}

// Mock for ExamAssessmentRepositoryInterface
type MockExamAssessmentRepository struct {
	mock.Mock
}

func (m *MockExamAssessmentRepository) Create(ctx context.Context, attemptId int, model repositories.AssessmentModel) (*ent.ExamAssesment, error) {
	args := m.Called(ctx, attemptId, model)
	return args.Get(0).(*ent.ExamAssesment), args.Error(1)
}

func (m *MockExamAssessmentRepository) Update(ctx context.Context, assessmentId int, model repositories.AssessmentModel) error {
	args := m.Called(ctx, assessmentId, model)
	return args.Error(0)
}

func (m *MockExamAssessmentRepository) GetById(ctx context.Context, assessmentId int, userId string) (*ent.ExamAssesment, error) {
	args := m.Called(ctx, assessmentId, userId)
	return args.Get(0).(*ent.ExamAssesment), args.Error(1)
}

func (m *MockExamAssessmentRepository) GetByExam(ctx context.Context, generatedExamId int, userId string) ([]*ent.ExamAssesment, error) {
	args := m.Called(ctx, generatedExamId, userId)
	return args.Get(0).([]*ent.ExamAssesment), args.Error(1)
}

// Test suite for ExamAssesmentService
func TestExamAssesmentService(t *testing.T) {
	mockAccessService := new(MockAccessService)
	mockPromptService := new(MockPromptService)
	mockProfanityService := new(MockProfanityService)
	mockGeneratedExamRepository := new(MockGeneratedExamRepository)
	mockExamGenerationService := new(MockExamGenerationService)
	mockExamAttemptRepository := new(MockExamAttemptRepository)
	mockExamAssessmentRepository := new(MockExamAssessmentRepository)

	examAssessmentService := services.NewExamAssesmentService(
		mockAccessService,
		mockPromptService,
		mockProfanityService,
		mockGeneratedExamRepository,
		mockExamGenerationService,
		mockExamAttemptRepository,
		mockExamAssessmentRepository,
	)

	ctx := context.Background()

	// Test for StartNewDescriptiveAssesment
	t.Run("StartNewDescriptiveAssesment Success", func(t *testing.T) {
		generatedExamId := 1
		attempt := &ent.ExamAttempt{ID: 1}
		request := &services.DescriptiveExamAssesmentRequest{CompletedSeconds: 120, Content: "Sample content"}
		userId := "test-user-id"
		isOpen := false

		// Mocking generated exam retrieval
		mockGeneratedExamRepository.On("GetOpenById", ctx, generatedExamId, isOpen).Return(&ent.GeneratedExam{Edges: ent.GeneratedExamEdges{Exam: &ent.Exam{ID: 1}}}, nil)
		mockAccessService.On("UserHasAccessToExam", ctx, 1, userId).Return(true, nil)

		// Mocking assessment creation
		mockExamAssessmentRepository.On("Create", ctx, attempt.ID, mock.Anything).Return(&ent.ExamAssesment{ID: 1}, nil)

		assessmentDetails, err := examAssessmentService.StartNewDescriptiveAssesment(ctx, generatedExamId, attempt, request, userId, isOpen)
		assert.NoError(t, err)
		assert.NotNil(t, assessmentDetails)

		// Assert expectations
		mockGeneratedExamRepository.AssertExpectations(t)
		mockAccessService.AssertExpectations(t)
		mockExamAssessmentRepository.AssertExpectations(t)
	})

	t.Run("StartNewDescriptiveAssesment Forbidden Access", func(t *testing.T) {
		generatedExamId := 1
		attempt := &ent.ExamAttempt{ID: 1}
		request := &services.DescriptiveExamAssesmentRequest{CompletedSeconds: 120, Content: "Sample content"}
		userId := "test-user-id"
		isOpen := false

		mockGeneratedExamRepository.On("GetOpenById", ctx, generatedExamId, isOpen).Return(&ent.GeneratedExam{Edges: ent.GeneratedExamEdges{Exam: &ent.Exam{ID: 1}}}, nil)
		mockAccessService.On("UserHasAccessToExam", ctx, 1, userId).Return(false, nil)

		assessmentDetails, err := examAssessmentService.StartNewDescriptiveAssesment(ctx, generatedExamId, attempt, request, userId, isOpen)
		assert.Error(t, err)
		assert.Nil(t, assessmentDetails)
		assert.Equal(t, "forbidden", err.Error())

		// Assert expectations
		mockGeneratedExamRepository.AssertExpectations(t)
		mockAccessService.AssertExpectations(t)
	})

	// Test for GetAssesmentById
	t.Run("GetAssesmentById Success", func(t *testing.T) {
		assessmentId := 1
		userId := "test-user-id"

		mockExamAssessmentRepository.On("GetById", ctx, assessmentId, userId).Return(&ent.ExamAssesment{
			ID:               assessmentId,
			CompletedSeconds: 120,
			Status:           examassesment.Status(constants.ASSESSMENT_COMPLETED),
		}, nil)

		assessmentDetails, err := examAssessmentService.GetAssesmentById(ctx, assessmentId, userId)
		assert.NoError(t, err)
		assert.NotNil(t, assessmentDetails)
		assert.Equal(t, assessmentId, assessmentDetails.Id)

		// Assert expectations
		mockExamAssessmentRepository.AssertExpectations(t)
	})

	// Test for GetExamAssessments
	t.Run("GetExamAssessments Success", func(t *testing.T) {
		generatedExamId := 1
		userId := "test-user-id"
		assessments := []*ent.ExamAssesment{
			{ID: 1, CompletedSeconds: 120, Status: examassesment.Status(constants.ASSESSMENT_COMPLETED)},
			{ID: 2, CompletedSeconds: 150, Status: examassesment.Status(constants.ASSESSMENT_PENDING)},
		}

		mockExamAssessmentRepository.On("GetByExam", ctx, generatedExamId, userId).Return(assessments, nil)

		result, err := examAssessmentService.GetExamAssessments(ctx, generatedExamId, userId)
		assert.NoError(t, err)
		assert.Len(t, result, len(assessments))

		// Assert expectations
		mockExamAssessmentRepository.AssertExpectations(t)
	})

}
