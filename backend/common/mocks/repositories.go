package mocks

import (
	"common/constants"
	"common/ent"
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

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
