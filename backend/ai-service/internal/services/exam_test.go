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
	mockGenAIService.On("GetStructuredContentStream", ctx, examSetting.AiPrompt, constants.PRO_15).Return("AI generated content", nil)

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
	mockGenAIService.On("GetStructuredContentStream", ctx, examSetting.AiPrompt, constants.PRO_15).Return("", errors.New("AI generation error"))

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

	// Mock the first call to GetStructuredContentStream for AI content generation
	mockGenAIService.On("GetStructuredContentStream", ctx, examSetting.AiPrompt, constants.PRO_15).Return("AI generated content", nil)

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

	mockGenAIService.On("GetStructuredContentStream", ctx, examSetting.AiPrompt, constants.PRO_15).Return("AI generated content", nil)

	mockRedisService.On("Store", ctx, mock.Anything, "AI generated content", services.DEFAULT_CACHE_EXPIRY).Return(nil)

	// Create a valid CachedExam instance to return from the mock
	cachedExam := &ent.CachedExam{ID: 1}
	mockCachedExamRepository.On("Create", ctx, mock.Anything, services.DEFAULT_CACHE_EXPIRY, exam).Return(cachedExam, errors.New("cache metadata error"))

	err := examService.PopulateExamQuestionCache(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cache metadata error")
}

func TestGenerateExamQuestionAndPopulateCache_ActiveExam_Success(t *testing.T) {
	mockExamRepository := new(mocks.MockExamRepository)
	mockExamSettingRepository := new(mocks.MockExamSettingRepository)
	mockGenAIService := new(mocks.MockGenAIService)
	mockRedisService := new(mocks.MockRedisService)
	mockCachedQuestionMetaDataRepository := new(mocks.MockCachedExamRepository)

	examService := services.NewExamService(
		mockGenAIService,
		mockRedisService,
		mockExamRepository,
		nil, // Assuming exam category repository is not used here
		mockExamSettingRepository,
		mockCachedQuestionMetaDataRepository, // Assuming cached exam repository is not used here
	)

	ctx := context.Background()
	examId := 1
	exam := &ent.Exam{ID: examId, Name: "Active Exam", IsActive: true}
	examSetting := &ent.ExamSetting{AiPrompt: "Generate questions for active exam."}
	aiResponse := "AI generated question content"

	mockExamRepository.On("GetById", ctx, examId).Return(exam, nil)
	mockExamSettingRepository.On("GetByExam", ctx, examId).Return(examSetting, nil)
	mockGenAIService.On("GetStructuredContentStream", ctx, examSetting.AiPrompt, constants.PRO_15).Return(aiResponse, nil)
	var capturedUID string

	// Setup the first mock to capture the dynamically generated UID
	mockRedisService.On("Store", mock.Anything, mock.MatchedBy(func(id string) bool {
		if id != "" { // Ensure the UID is not empty
			capturedUID = id // Capture the UID for use in the next mock
			return true
		}
		return false
	}), mock.Anything, mock.Anything).Return(nil)

	// Ensure subsequent use of the captured UID in another mock
	mockCachedQuestionMetaDataRepository.On("Create", ctx, mock.MatchedBy(func(id string) bool {
		return id == capturedUID // Match the captured UID
	}), services.DEFAULT_CACHE_EXPIRY, exam).Return(&ent.CachedExam{ID: 1}, nil)

	_, err := examService.GenerateExamQuestionAndPopulateCache(ctx, examId)
	assert.NoError(t, err)

	mockExamRepository.AssertExpectations(t)
	mockExamSettingRepository.AssertExpectations(t)
	mockGenAIService.AssertExpectations(t)
	mockRedisService.AssertExpectations(t)
	mockCachedQuestionMetaDataRepository.AssertExpectations(t)
}

func TestGenerateExamQuestionAndPopulateCache_InactiveExam_Error(t *testing.T) {
	mockExamRepository := new(mocks.MockExamRepository)
	mockExamSettingRepository := new(mocks.MockExamSettingRepository)

	examService := services.NewExamService(
		nil, // Gen AI service not called for inactive exams
		nil, // Redis service not called for inactive exams
		mockExamRepository,
		nil, // Not used here
		mockExamSettingRepository,
		nil, // Not used here
	)

	ctx := context.Background()
	examId := 2
	exam := &ent.Exam{ID: examId, Name: "Inactive Exam", IsActive: false}

	mockExamRepository.On("GetById", ctx, examId).Return(exam, nil)

	_, err := examService.GenerateExamQuestionAndPopulateCache(ctx, examId)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Skipping inactive exam")

	mockExamRepository.AssertExpectations(t)
	mockExamSettingRepository.AssertNotCalled(t, "GetByExam", mock.Anything, mock.Anything)
}
