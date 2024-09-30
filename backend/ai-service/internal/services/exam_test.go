package services_test

import (
	"ai-service/internal/services"
	"context"
	"errors"
	"testing"
	"time"

	"common/constants"
	"common/ent"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for GenAIServiceInterface
type MockGenAIService struct {
	mock.Mock
}

func (m *MockGenAIService) GetContentStream(ctx context.Context, prompt string, model constants.GenAiModel) (string, error) {
	args := m.Called(ctx, prompt, model)
	return args.String(0), args.Error(1)
}

// Mock for RedisServiceInterface
type MockRedisService struct {
	mock.Mock
}

// Mock implementation for Store method
func (m *MockRedisService) Store(ctx context.Context, key string, value string, expiration time.Duration) error {
	args := m.Called(ctx, key, value, expiration)
	return args.Error(0)
}

// Mock implementation for Get method
func (m *MockRedisService) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

// Mock implementation for Delete method
func (m *MockRedisService) Delete(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}

// Mock implementation for Health method
func (m *MockRedisService) Health() map[string]string {
	args := m.Called()
	return args.Get(0).(map[string]string)
}

// Mock implementation for CheckRedisHealth method
func (m *MockRedisService) CheckRedisHealth(ctx context.Context, stats map[string]string) map[string]string {
	args := m.Called(ctx, stats)
	return args.Get(0).(map[string]string)
}

// Mock for ExamRepositoryInterface
type MockExamRepository struct {
	mock.Mock
}

// Mock implementation for GetById method
func (m *MockExamRepository) GetById(ctx context.Context, examId int) (*ent.Exam, error) {
	args := m.Called(ctx, examId)
	return args.Get(0).(*ent.Exam), args.Error(1)
}

// Mock implementation for GetActiveByExamsGroupId method
func (m *MockExamRepository) GetActiveByExamsGroupId(ctx context.Context, examGroupId int, isActive bool) ([]*ent.Exam, error) {
	args := m.Called(ctx, examGroupId, isActive)
	return args.Get(0).([]*ent.Exam), args.Error(1)
}

// Mock implementation for GetActiveById method
func (m *MockExamRepository) GetActiveById(ctx context.Context, examId int, isActive bool) (*ent.Exam, error) {
	args := m.Called(ctx, examId, isActive)
	return args.Get(0).(*ent.Exam), args.Error(1)
}

// Mock implementation for GetByExamCategory method
func (m *MockExamRepository) GetByExamCategory(ctx context.Context, examCategory *ent.ExamCategory) ([]*ent.Exam, error) {
	args := m.Called(ctx, examCategory)
	return args.Get(0).([]*ent.Exam), args.Error(1)
}

// Mock implementation for GetActiveByType method
func (m *MockExamRepository) GetActiveByType(ctx context.Context, examType constants.ExamType) ([]*ent.Exam, error) {
	args := m.Called(ctx, examType)
	return args.Get(0).([]*ent.Exam), args.Error(1)
}

// Mock implementation for GetByName method
func (m *MockExamRepository) GetByName(ctx context.Context, examName string) (*ent.Exam, error) {
	args := m.Called(ctx, examName)
	return args.Get(0).(*ent.Exam), args.Error(1)
}

type MockExamCategoryRepository struct {
	mock.Mock
}

// Mock implementation for Get method
func (m *MockExamCategoryRepository) Get(ctx context.Context) ([]*ent.ExamCategory, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*ent.ExamCategory), args.Error(1)
}

// Mock implementation for GetByName method
func (m *MockExamCategoryRepository) GetByName(ctx context.Context, categoryName constants.ExamCategoryName) (*ent.ExamCategory, error) {
	args := m.Called(ctx, categoryName)
	return args.Get(0).(*ent.ExamCategory), args.Error(1)
}

// Mock for ExamSettingRepositoryInterface
type MockExamSettingRepository struct {
	mock.Mock
}

func (m *MockExamSettingRepository) GetByExam(ctx context.Context, examID int) (*ent.ExamSetting, error) {
	args := m.Called(ctx, examID)
	return args.Get(0).(*ent.ExamSetting), args.Error(1)
}

// Mock for CachedExamRepositoryInterface
type MockCachedExamRepository struct {
	mock.Mock
}

// Mock implementation for Create method
func (m *MockCachedExamRepository) Create(ctx context.Context, uid string, expiration time.Duration, exam *ent.Exam) (*ent.CachedExam, error) {
	args := m.Called(ctx, uid, expiration, exam)
	return args.Get(0).(*ent.CachedExam), args.Error(1)
}

// Mock implementation for GetByExam method
func (m *MockCachedExamRepository) GetByExam(ctx context.Context, ex *ent.Exam) ([]*ent.CachedExam, error) {
	args := m.Called(ctx, ex)
	return args.Get(0).([]*ent.CachedExam), args.Error(1)
}

// Mock implementation for MarkAsUsed method
func (m *MockCachedExamRepository) MarkAsUsed(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestPopulateExamQuestionCache_Success(t *testing.T) {
	mockGenAIService := new(MockGenAIService)
	mockRedisService := new(MockRedisService)
	mockExamRepository := new(MockExamRepository)
	mockExamCategoryRepository := new(MockExamCategoryRepository)
	mockExamSettingRepository := new(MockExamSettingRepository)
	mockCachedExamRepository := new(MockCachedExamRepository)

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
	mockGenAIService.On("GetContentStream", ctx, mock.Anything, constants.PRO_15).Return("Validation successful", nil)

	// Mocking Redis service
	mockRedisService.On("Store", ctx, mock.Anything, "Validation successful", services.DEFAULT_CACHE_EXPIRY).Return(nil)

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
	mockGenAIService := new(MockGenAIService)
	mockRedisService := new(MockRedisService)
	mockExamRepository := new(MockExamRepository)
	mockExamCategoryRepository := new(MockExamCategoryRepository)
	mockExamSettingRepository := new(MockExamSettingRepository)
	mockCachedExamRepository := new(MockCachedExamRepository)

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
	mockGenAIService := new(MockGenAIService)
	mockRedisService := new(MockRedisService)
	mockExamRepository := new(MockExamRepository)
	mockExamCategoryRepository := new(MockExamCategoryRepository)
	mockExamSettingRepository := new(MockExamSettingRepository)
	mockCachedExamRepository := new(MockCachedExamRepository)

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
	mockGenAIService := new(MockGenAIService)
	mockRedisService := new(MockRedisService)
	mockExamRepository := new(MockExamRepository)
	mockExamCategoryRepository := new(MockExamCategoryRepository)
	mockExamSettingRepository := new(MockExamSettingRepository)
	mockCachedExamRepository := new(MockCachedExamRepository)

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
	mockGenAIService := new(MockGenAIService)
	mockRedisService := new(MockRedisService)
	mockExamRepository := new(MockExamRepository)
	mockExamCategoryRepository := new(MockExamCategoryRepository)
	mockExamSettingRepository := new(MockExamSettingRepository)
	mockCachedExamRepository := new(MockCachedExamRepository)

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
	mockGenAIService := new(MockGenAIService)
	mockRedisService := new(MockRedisService)
	mockExamRepository := new(MockExamRepository)
	mockExamCategoryRepository := new(MockExamCategoryRepository)
	mockExamSettingRepository := new(MockExamSettingRepository)
	mockCachedExamRepository := new(MockCachedExamRepository)

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

	// Mock the second call to GetContentStream for validation
	mockGenAIService.On("GetContentStream", ctx, mock.Anything, constants.PRO_15).Return("Validation successful", nil)

	// Simulate an error when storing in Redis
	mockRedisService.On("Store", ctx, mock.Anything, "Validation successful", services.DEFAULT_CACHE_EXPIRY).Return(errors.New("redis store error"))

	err := examService.PopulateExamQuestionCache(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "redis store error")
}

func TestPopulateExamQuestionCache_ErrorSavingCachedMetadata(t *testing.T) {
	mockGenAIService := new(MockGenAIService)
	mockRedisService := new(MockRedisService)
	mockExamRepository := new(MockExamRepository)
	mockExamCategoryRepository := new(MockExamCategoryRepository)
	mockExamSettingRepository := new(MockExamSettingRepository)
	mockCachedExamRepository := new(MockCachedExamRepository)

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
	mockGenAIService.On("GetContentStream", ctx, mock.Anything, constants.PRO_15).Return("Validation successful", nil)

	mockRedisService.On("Store", ctx, mock.Anything, "Validation successful", services.DEFAULT_CACHE_EXPIRY).Return(nil)

	// Create a valid CachedExam instance to return from the mock
	cachedExam := &ent.CachedExam{ID: 1}
	mockCachedExamRepository.On("Create", ctx, mock.Anything, services.DEFAULT_CACHE_EXPIRY, exam).Return(cachedExam, errors.New("cache metadata error"))

	err := examService.PopulateExamQuestionCache(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cache metadata error")
}
