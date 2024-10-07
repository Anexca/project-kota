package services_test

import (
	"context"
	"encoding/json"
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
	t.Run("StartNewDescriptiveAssesment Forbidden Access", func(t *testing.T) {
		generatedExamId := 1
		attempt := &ent.ExamAttempt{ID: 1}
		request := &models.DescriptiveExamAssesmentRequest{CompletedSeconds: 120, Content: "Sample content"}
		userId := "test-user-id"
		isOpen := false

		// Mocking necessary methods to reach the forbidden access logic
		mockGeneratedExamRepository.On("GetOpenById", ctx, generatedExamId, isOpen).Return(&ent.GeneratedExam{Edges: ent.GeneratedExamEdges{Exam: &ent.Exam{ID: 1}}}, nil)
		mockAccessService.On("UserHasAccessToExam", ctx, 1, userId).Return(false, nil) // User does not have access

		// Call the method under test
		assessmentDetails, err := examAssessmentService.StartNewDescriptiveAssesment(ctx, generatedExamId, attempt, request, userId, isOpen)

		// Assert the expected error
		assert.Error(t, err)
		assert.Equal(t, "forbidden", err.Error())
		assert.Nil(t, assessmentDetails, "Expected no assessment details to be returned")

		// Assert that all expectations are met
		mockGeneratedExamRepository.AssertExpectations(t)
		mockAccessService.AssertExpectations(t)
	})

	t.Run("StartNewDescriptiveAssesment Forbidden Access", func(t *testing.T) {
		generatedExamId := 1
		attempt := &ent.ExamAttempt{ID: 1}
		request := &models.DescriptiveExamAssesmentRequest{CompletedSeconds: 120, Content: "Sample content"}
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

// Helper function to create a mock MCQ exam data
func createMockMCQExam() map[string]interface{} {
	exam := models.GeneratedMCQExam{
		Sections: map[string][]models.MCQExamQuestion{
			"ENGLISH": {
				{
					ContentReferenceId: "ref1", // Example content reference ID
					Question:           "QQ",
					QuestionNumber:     1,
					Answer:             []int{1}, // Example answer index
					Options:            []string{"Option A", "Option B", "Option C"},
					Explanation:        "This is why Option B is correct.",
				},
			},
		},
	}

	jsonData, _ := json.Marshal(exam)
	var mappedData map[string]interface{}

	_ = json.Unmarshal(jsonData, &mappedData)
	return mappedData
}

func TestAssessMCQExam(t *testing.T) {
	ctx := context.Background()
	t.Run("AssessMCQExam Success", func(t *testing.T) {
		// Mock services
		mockAccessService := new(mocks.MockAccessService)
		mockPromptService := new(mocks.MockPromptService)
		mockProfanityService := new(commonMocks.MockProfanityService)
		mockGeneratedExamRepository := new(commonMocks.MockGeneratedExamRepository)
		mockExamGenerationService := new(mocks.MockExamGenerationService)
		mockExamAttemptRepository := new(commonMocks.MockExamAttemptRepository)
		mockExamAssessmentRepository := new(commonMocks.MockExamAssessmentRepository)

		mockAccessService.ExpectedCalls = nil
		mockAccessService.Calls = nil

		// Initialize the service with the mocked dependencies
		service := services.NewExamAssesmentService(
			mockAccessService,
			mockPromptService,
			mockProfanityService,
			mockGeneratedExamRepository,
			mockExamGenerationService,
			mockExamAttemptRepository,
			mockExamAssessmentRepository,
		)

		request := &models.MCQExamAssessmentRequest{
			AttemptedQuestions: []models.MCQExamAssessmentRequestModel{
				{
					QuestionNumber:          1,
					UserSelectedOptionIndex: []int{1},
				},
			},
			CompletedSeconds: 100,
		}
		attempt := &ent.ExamAttempt{ID: 1}
		generatedExam := &ent.GeneratedExam{
			RawExamData: createMockMCQExam(),
			Edges: ent.GeneratedExamEdges{
				Exam: &ent.Exam{
					ID: 1,
					Edges: ent.ExamEdges{
						Setting: &ent.ExamSetting{
							NegativeMarking: 0.25,
						},
					},
				},
			},
		}

		mockGeneratedExamRepository.On("GetOpenById", ctx, mock.Anything, mock.Anything).Return(generatedExam, nil)
		mockAccessService.On("UserHasAccessToExam", ctx, mock.Anything, mock.Anything).Return(true, nil)
		mockExamAssessmentRepository.On("Create", ctx, mock.Anything, mock.Anything).Return(&ent.ExamAssesment{
			ID: 1,
		}, nil)

		result, err := service.AssessMCQExam(ctx, 1, attempt, request, "user1", true)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("AssessMCQExam Access Denied", func(t *testing.T) {
		// Mock services
		mockAccessService := new(mocks.MockAccessService)
		mockPromptService := new(mocks.MockPromptService)
		mockProfanityService := new(commonMocks.MockProfanityService)
		mockGeneratedExamRepository := new(commonMocks.MockGeneratedExamRepository)
		mockExamGenerationService := new(mocks.MockExamGenerationService)
		mockExamAttemptRepository := new(commonMocks.MockExamAttemptRepository)
		mockExamAssessmentRepository := new(commonMocks.MockExamAssessmentRepository)

		mockAccessService.ExpectedCalls = nil
		mockAccessService.Calls = nil

		// Initialize the service with the mocked dependencies
		service := services.NewExamAssesmentService(
			mockAccessService,
			mockPromptService,
			mockProfanityService,
			mockGeneratedExamRepository,
			mockExamGenerationService,
			mockExamAttemptRepository,
			mockExamAssessmentRepository,
		)

		request := &models.MCQExamAssessmentRequest{
			CompletedSeconds: 100,
		}
		attempt := &ent.ExamAttempt{ID: 1}
		generatedExam := &ent.GeneratedExam{
			RawExamData: createMockMCQExam(),
			Edges: ent.GeneratedExamEdges{
				Exam: &ent.Exam{
					ID: 1,
				},
			},
		}

		mockGeneratedExamRepository.On("GetOpenById", ctx, mock.Anything, mock.Anything).Return(generatedExam, nil)
		mockAccessService.On("UserHasAccessToExam", ctx, mock.Anything, mock.Anything).Return(false, nil)

		result, err := service.AssessMCQExam(ctx, 1, attempt, request, "user1", false)
		assert.Error(t, err)
		assert.Equal(t, "forbidden", err.Error())
		assert.Nil(t, result)
	})
}
