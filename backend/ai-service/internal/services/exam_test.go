package services_test

import (
	"ai-service/internal/services"
	"context"
	"errors"
	"testing"

	"common/constants"
	"common/ent"
	"common/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPopulateExamQuestionCache_Success(t *testing.T) {
	mockGenAIService := new(mocks.MockGenAIService)
	mockRedisService := new(mocks.MockRedisService)
	mockExamRepository := new(mocks.MockExamRepository)
	mockExamCategoryRepository := new(mocks.MockExamCategoryRepository)
	mockExamSettingRepository := new(mocks.MockExamSettingRepository)
	mockCachedExamRepository := new(mocks.MockCachedExamRepository)

	examService := services.NewExamService(
		mockGenAIService,
		mockRedisService,
		mockExamRepository,
		mockExamCategoryRepository,
		mockExamSettingRepository,
		mockCachedExamRepository,
	)

	ctx := context.Background()

	// Mocking the exam categories
	category := &ent.ExamCategory{ID: 1, Name: "Banking"}
	mockExamCategoryRepository.On("Get", ctx).Return([]*ent.ExamCategory{category}, nil)

	// Mocking the exams
	exam := &ent.Exam{ID: 1, Name: "Banking Exam", IsActive: true}
	mockExamRepository.On("GetByExamCategory", ctx, category).Return([]*ent.Exam{exam}, nil)

	// Mocking the exam settings
	examSetting := &ent.ExamSetting{AiPrompt: "Generate questions for banking exam."}
	mockExamSettingRepository.On("GetByExam", ctx, exam.ID).Return(examSetting, nil)

	// Mocking AI service
	mockGenAIService.On("GetContentStream", ctx, examSetting.AiPrompt, constants.PRO_15).Return("AI generated content", nil)

	// Mocking Redis service
	mockRedisService.On("Store", ctx, mock.Anything, "AI generated content", services.DEFAULT_CACHE_EXPIRY).Return(nil)

	// Mocking cached exam repository
	cachedExam := &ent.CachedExam{ID: 1}
	mockCachedExamRepository.On("Create", ctx, mock.Anything, services.DEFAULT_CACHE_EXPIRY, exam).Return(cachedExam, nil)

	// Execute the method
	err := examService.PopulateExamQuestionCache(ctx)
	assert.NoError(t, err)

	// Assert expectations
	mockExamCategoryRepository.AssertExpectations(t)
	mockExamRepository.AssertExpectations(t)
	mockExamSettingRepository.AssertExpectations(t)
	mockGenAIService.AssertExpectations(t)
	mockRedisService.AssertExpectations(t)
	mockCachedExamRepository.AssertExpectations(t)
}

func TestPopulateExamQuestionCache_ErrorFetchingCategories(t *testing.T) {
	mockGenAIService := new(mocks.MockGenAIService)
	mockRedisService := new(mocks.MockRedisService)
	mockExamRepository := new(mocks.MockExamRepository)
	mockExamCategoryRepository := new(mocks.MockExamCategoryRepository)
	mockExamSettingRepository := new(mocks.MockExamSettingRepository)
	mockCachedExamRepository := new(mocks.MockCachedExamRepository)

	examService := services.NewExamService(
		mockGenAIService,
		mockRedisService,
		mockExamRepository,
		mockExamCategoryRepository,
		mockExamSettingRepository,
		mockCachedExamRepository,
	)

	ctx := context.Background()

	// Simulate an error when fetching exam categories
	mockExamCategoryRepository.On("Get", ctx).Return([]*ent.ExamCategory{}, errors.New("category fetch error"))

	err := examService.PopulateExamQuestionCache(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "category fetch error")
}

func TestPopulateExamQuestionCache_ErrorFetchingExams(t *testing.T) {
	mockGenAIService := new(mocks.MockGenAIService)
	mockRedisService := new(mocks.MockRedisService)
	mockExamRepository := new(mocks.MockExamRepository)
	mockExamCategoryRepository := new(mocks.MockExamCategoryRepository)
	mockExamSettingRepository := new(mocks.MockExamSettingRepository)
	mockCachedExamRepository := new(mocks.MockCachedExamRepository)

	examService := services.NewExamService(
		mockGenAIService,
		mockRedisService,
		mockExamRepository,
		mockExamCategoryRepository,
		mockExamSettingRepository,
		mockCachedExamRepository,
	)

	ctx := context.Background()

	category := &ent.ExamCategory{ID: 1, Name: "Banking"}
	mockExamCategoryRepository.On("Get", ctx).Return([]*ent.ExamCategory{category}, nil)

	// Simulate an error when fetching exams for the category
	mockExamRepository.On("GetByExamCategory", ctx, category).Return([]*ent.Exam{}, errors.New("exam fetch error"))

	err := examService.PopulateExamQuestionCache(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "exam fetch error")
}

func TestPopulateExamQuestionCache_ErrorFetchingExamSetting(t *testing.T) {
	mockGenAIService := new(mocks.MockGenAIService)
	mockRedisService := new(mocks.MockRedisService)
	mockExamRepository := new(mocks.MockExamRepository)
	mockExamCategoryRepository := new(mocks.MockExamCategoryRepository)
	mockExamSettingRepository := new(mocks.MockExamSettingRepository)
	mockCachedExamRepository := new(mocks.MockCachedExamRepository)

	examService := services.NewExamService(
		mockGenAIService,
		mockRedisService,
		mockExamRepository,
		mockExamCategoryRepository,
		mockExamSettingRepository,
		mockCachedExamRepository,
	)

	ctx := context.Background()

	category := &ent.ExamCategory{ID: 1, Name: "Banking"}
	mockExamCategoryRepository.On("Get", ctx).Return([]*ent.ExamCategory{category}, nil)

	exam := &ent.Exam{ID: 1, Name: "Banking Exam", IsActive: true}
	mockExamRepository.On("GetByExamCategory", ctx, category).Return([]*ent.Exam{exam}, nil)

	// Simulate an error when fetching exam settings
	mockExamSettingRepository.On("GetByExam", ctx, exam.ID).Return(&ent.ExamSetting{}, errors.New("exam setting fetch error"))

	err := examService.PopulateExamQuestionCache(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "exam setting fetch error")
}

func TestPopulateExamQuestionCache_ErrorGeneratingContent(t *testing.T) {
	mockGenAIService := new(mocks.MockGenAIService)
	mockRedisService := new(mocks.MockRedisService)
	mockExamRepository := new(mocks.MockExamRepository)
	mockExamCategoryRepository := new(mocks.MockExamCategoryRepository)
	mockExamSettingRepository := new(mocks.MockExamSettingRepository)
	mockCachedExamRepository := new(mocks.MockCachedExamRepository)

	examService := services.NewExamService(
		mockGenAIService,
		mockRedisService,
		mockExamRepository,
		mockExamCategoryRepository,
		mockExamSettingRepository,
		mockCachedExamRepository,
	)

	ctx := context.Background()

	category := &ent.ExamCategory{ID: 1, Name: "Banking"}
	mockExamCategoryRepository.On("Get", ctx).Return([]*ent.ExamCategory{category}, nil)

	exam := &ent.Exam{ID: 1, Name: "Banking Exam", IsActive: true}
	mockExamRepository.On("GetByExamCategory", ctx, category).Return([]*ent.Exam{exam}, nil)

	examSetting := &ent.ExamSetting{AiPrompt: "Generate questions for banking exam."}
	mockExamSettingRepository.On("GetByExam", ctx, exam.ID).Return(examSetting, nil)

	// Simulate an error when generating AI content
	mockGenAIService.On("GetContentStream", ctx, examSetting.AiPrompt, constants.PRO_15).Return("", errors.New("AI generation error"))

	err := examService.PopulateExamQuestionCache(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "AI generation error")
}

func TestPopulateExamQuestionCache_ErrorStoringInRedis(t *testing.T) {
	mockGenAIService := new(mocks.MockGenAIService)
	mockRedisService := new(mocks.MockRedisService)
	mockExamRepository := new(mocks.MockExamRepository)
	mockExamCategoryRepository := new(mocks.MockExamCategoryRepository)
	mockExamSettingRepository := new(mocks.MockExamSettingRepository)
	mockCachedExamRepository := new(mocks.MockCachedExamRepository)

	examService := services.NewExamService(
		mockGenAIService,
		mockRedisService,
		mockExamRepository,
		mockExamCategoryRepository,
		mockExamSettingRepository,
		mockCachedExamRepository,
	)

	ctx := context.Background()

	category := &ent.ExamCategory{ID: 1, Name: "Banking"}
	mockExamCategoryRepository.On("Get", ctx).Return([]*ent.ExamCategory{category}, nil)

	exam := &ent.Exam{ID: 1, Name: "Banking Exam", IsActive: true}
	mockExamRepository.On("GetByExamCategory", ctx, category).Return([]*ent.Exam{exam}, nil)

	examSetting := &ent.ExamSetting{AiPrompt: "Generate questions for banking exam."}
	mockExamSettingRepository.On("GetByExam", ctx, exam.ID).Return(examSetting, nil)

	// Mock the first call to GetContentStream for AI content generation
	mockGenAIService.On("GetContentStream", ctx, examSetting.AiPrompt, constants.PRO_15).Return("AI generated content", nil)

	// Simulate an error when storing in Redis
	mockRedisService.On("Store", ctx, mock.Anything, "AI generated content", services.DEFAULT_CACHE_EXPIRY).Return(errors.New("redis store error"))

	err := examService.PopulateExamQuestionCache(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "redis store error")
}

func TestPopulateExamQuestionCache_ErrorSavingCachedMetadata(t *testing.T) {
	mockGenAIService := new(mocks.MockGenAIService)
	mockRedisService := new(mocks.MockRedisService)
	mockExamRepository := new(mocks.MockExamRepository)
	mockExamCategoryRepository := new(mocks.MockExamCategoryRepository)
	mockExamSettingRepository := new(mocks.MockExamSettingRepository)
	mockCachedExamRepository := new(mocks.MockCachedExamRepository)

	examService := services.NewExamService(
		mockGenAIService,
		mockRedisService,
		mockExamRepository,
		mockExamCategoryRepository,
		mockExamSettingRepository,
		mockCachedExamRepository,
	)

	ctx := context.Background()

	category := &ent.ExamCategory{ID: 1, Name: "Banking"}
	mockExamCategoryRepository.On("Get", ctx).Return([]*ent.ExamCategory{category}, nil)

	exam := &ent.Exam{ID: 1, Name: "Banking Exam", IsActive: true}
	mockExamRepository.On("GetByExamCategory", ctx, category).Return([]*ent.Exam{exam}, nil)

	examSetting := &ent.ExamSetting{AiPrompt: "Generate questions for banking exam."}
	mockExamSettingRepository.On("GetByExam", ctx, exam.ID).Return(examSetting, nil)

	mockGenAIService.On("GetContentStream", ctx, examSetting.AiPrompt, constants.PRO_15).Return("AI generated content", nil)

	mockRedisService.On("Store", ctx, mock.Anything, "AI generated content", services.DEFAULT_CACHE_EXPIRY).Return(nil)

	// Create a valid CachedExam instance to return from the mock
	cachedExam := &ent.CachedExam{ID: 1}
	mockCachedExamRepository.On("Create", ctx, mock.Anything, services.DEFAULT_CACHE_EXPIRY, exam).Return(cachedExam, errors.New("cache metadata error"))

	err := examService.PopulateExamQuestionCache(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cache metadata error")
}
