package services_test

import (
	"context"
	"testing"

	"common/constants"
	"common/ent"
	"common/ent/examassesment"
	commonMocks "common/mocks"
	"server/internal/mocks"
	"server/internal/services"
	"server/pkg/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Test suite for ExamAssesmentService
func TestExamAssesmentService(t *testing.T) {
	mockAccessService := new(mocks.MockAccessService)
	mockPromptService := new(mocks.MockPromptService)
	mockProfanityService := new(commonMocks.MockProfanityService)
	mockGeneratedExamRepository := new(commonMocks.MockGeneratedExamRepository)
	mockExamGenerationService := new(mocks.MockExamGenerationService)
	mockExamAttemptRepository := new(commonMocks.MockExamAttemptRepository)
	mockExamAssessmentRepository := new(commonMocks.MockExamAssessmentRepository)

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
		mockAccessService.On("UserHasAccessToExam", ctx, generatedExamId, userId).Return(true, nil)

		// Mocking the behavior for GetGeneratedExamById
		mockExamGenerationService.On("GetGeneratedExamById", ctx, generatedExamId, userId, isOpen).Return(&models.GeneratedExamOverview{}, nil)

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
		mockAccessService.On("UserHasAccessToExam", ctx, generatedExamId, userId).Return(false, nil)

		assessmentDetails, err := examAssessmentService.StartNewDescriptiveAssesment(ctx, generatedExamId, attempt, request, userId, isOpen)
		assert.Error(t, err)
		assert.Nil(t, assessmentDetails)
		assert.Equal(t, "forbidden", err.Error())

		// Assert expectations
		mockGeneratedExamRepository.AssertExpectations(t)
		mockAccessService.AssertExpectations(t)
	})

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
