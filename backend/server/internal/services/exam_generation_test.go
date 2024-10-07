package services_test

import (
	"context"
	"testing"

	"common/ent"
	commonMocks "common/mocks"
	"server/internal/mocks"
	"server/internal/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGenerateExams(t *testing.T) {
	mockAccessService := new(mocks.MockAccessService)
	mockRedisService := new(commonMocks.MockRedisService)
	mockExamRepository := new(commonMocks.MockExamRepository)
	mockExamGroupRepository := new(commonMocks.MockExamGroupRepository)
	mockGeneratedExamRepository := new(commonMocks.MockGeneratedExamRepository)
	mockExamCategoryRepository := new(commonMocks.MockExamCategoryRepository)
	mockExamSettingRepository := new(commonMocks.MockExamSettingRepository)
	mockExamAttemptRepository := new(commonMocks.MockExamAttemptRepository)
	mockCachedExamRepository := new(commonMocks.MockCachedExamRepository)

	service := services.NewExamGenerationService(
		mockAccessService,
		mockRedisService,
		mockExamRepository,
		mockExamGroupRepository,
		mockExamCategoryRepository,
		mockCachedExamRepository,
		mockGeneratedExamRepository,
		mockExamSettingRepository,
		mockExamAttemptRepository,
	)

	ctx := context.Background()

	t.Run("Mark Questions as Open", func(t *testing.T) {
		examName := "Math Exam"
		mockExamRepository.On("GetByName", ctx, examName).Return(&ent.Exam{ID: 1}, nil)
		mockGeneratedExamRepository.On("GetByOpenFlag", ctx, 1).Return([]*ent.GeneratedExam{
			{IsOpen: true},
		}, nil)
		mockGeneratedExamRepository.On("UpdateMany", ctx, mock.Anything).Return(nil)
		mockGeneratedExamRepository.On("GetByWeekOffset", ctx, mock.Anything, 1, 2).Return([]*ent.GeneratedExam{
			{IsOpen: false},
		}, nil)

		err := service.MarkQuestionsAsOpen(ctx, examName)
		assert.NoError(t, err)

		mockExamRepository.AssertExpectations(t)
		mockGeneratedExamRepository.AssertExpectations(t)
	})

	t.Run("Fetch Cached Exam Data - No Cache", func(t *testing.T) {
		exam := &ent.Exam{ID: 1}
		mockCachedExamRepository.On("GetByExam", ctx, exam).Return([]*ent.CachedExam{}, nil) // Return empty slice

		data, err := service.FetchCachedExamData(ctx, exam)
		assert.Error(t, err)
		assert.Empty(t, data) // Ensure data is empty

		mockCachedExamRepository.AssertExpectations(t)
	})

	t.Run("Fetch Cached Exam Data - No Cache", func(t *testing.T) {
		exam := &ent.Exam{ID: 1}
		mockCachedExamRepository.On("GetByExam", ctx, exam).Return([]*ent.CachedExam{}, nil) // Return an empty slice

		data, err := service.FetchCachedExamData(ctx, exam)
		assert.Error(t, err)  // Expect an error
		assert.Empty(t, data) // Ensure data is empty

		mockCachedExamRepository.AssertExpectations(t)
	})

}
