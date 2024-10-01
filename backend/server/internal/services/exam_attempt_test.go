package services_test

import (
	"context"
	"testing"

	"common/ent"
	commonMocks "common/mocks"
	"server/internal/mocks"
	"server/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestExamAttemptService_CheckAndAddAttempt(t *testing.T) {
	mockAccessService := new(mocks.MockAccessService)
	mockExamRepository := new(commonMocks.MockExamRepository)
	mockExamAttemptRepository := new(commonMocks.MockExamAttemptRepository)
	mockExamSettingRepository := new(commonMocks.MockExamSettingRepository)
	mockGeneratedExamRepository := new(commonMocks.MockGeneratedExamRepository)

	service := services.NewExamAttemptService(
		mockAccessService,
		mockExamRepository,
		mockExamAttemptRepository,
		mockExamSettingRepository,
		mockGeneratedExamRepository,
	)

	t.Run("Success - First Attempt", func(t *testing.T) {
		ctx := context.Background()
		generatedExamId := 1
		userId := "test-user-id"
		isOpen := false

		mockExamAttemptRepository.On("GetByExam", ctx, generatedExamId, userId).Return([]*ent.ExamAttempt{}, nil)
		mockGeneratedExamRepository.On("GetOpenById", ctx, generatedExamId, isOpen).Return(&ent.GeneratedExam{Edges: ent.GeneratedExamEdges{Exam: &ent.Exam{ID: 1}}}, nil)
		mockAccessService.On("UserHasAccessToExam", ctx, 1, userId).Return(true, nil)
		mockExamSettingRepository.On("GetByExam", ctx, 1).Return(&ent.ExamSetting{MaxAttempts: 3}, nil)
		mockExamAttemptRepository.On("Create", ctx, 0, generatedExamId, userId).Return(&ent.ExamAttempt{ID: 1}, nil)

		attempt, err := service.CheckAndAddAttempt(ctx, generatedExamId, userId, isOpen)

		assert.NoError(t, err)
		assert.NotNil(t, attempt)
		assert.Equal(t, 1, attempt.ID)

		mockExamAttemptRepository.AssertExpectations(t)
		mockGeneratedExamRepository.AssertExpectations(t)
		mockAccessService.AssertExpectations(t)
		mockExamSettingRepository.AssertExpectations(t)
	})

	t.Run("Error - Forbidden Access", func(t *testing.T) {
		ctx := context.Background()
		generatedExamId := 1
		userId := "user-id"
		isOpen := false

		// Mock the expected call to GetByExam with the specific parameters used in the test
		mockExamAttemptRepository.On("GetByExam", ctx, generatedExamId, userId).Return([]*ent.ExamAttempt{}, nil)

		// Mock the call to GetOpenById
		mockGeneratedExamRepository.On("GetOpenById", ctx, generatedExamId, isOpen).Return(&ent.GeneratedExam{Edges: ent.GeneratedExamEdges{Exam: &ent.Exam{ID: 1}}}, nil)

		// Mock the call to UserHasAccessToExam to return false, simulating forbidden access
		mockAccessService.On("UserHasAccessToExam", ctx, 1, userId).Return(false, nil)

		// Perform the operation
		attempt, err := service.CheckAndAddAttempt(ctx, generatedExamId, userId, isOpen)

		// Verify the results
		assert.Error(t, err)
		assert.Nil(t, attempt)
		assert.Equal(t, "forbidden", err.Error())

		// Assert that all expectations are met
		mockExamAttemptRepository.AssertExpectations(t)
		mockGeneratedExamRepository.AssertExpectations(t)
		mockAccessService.AssertExpectations(t)
	})

}
