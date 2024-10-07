package mocks

import (
	"common/constants"
	"common/ent"
	"common/repositories"
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

// Mock for UserSubscriptionRepositoryInterface
type MockUserSubscriptionRepository struct {
	mock.Mock
}

func (m *MockUserSubscriptionRepository) Create(ctx context.Context, model repositories.UserSubscriptionModel) (*ent.UserSubscription, error) {
	args := m.Called(ctx, model)
	return args.Get(0).(*ent.UserSubscription), args.Error(1)
}

func (m *MockUserSubscriptionRepository) Update(ctx context.Context, updatedUserSubscription *ent.UserSubscription) error {
	args := m.Called(ctx, updatedUserSubscription)
	return args.Error(0)
}

func (m *MockUserSubscriptionRepository) GetById(ctx context.Context, userSubscriptionId int, userId string) (*ent.UserSubscription, error) {
	args := m.Called(ctx, userSubscriptionId, userId)
	return args.Get(0).(*ent.UserSubscription), args.Error(1)
}

func (m *MockUserSubscriptionRepository) GetByUserId(ctx context.Context, userId string) ([]*ent.UserSubscription, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).([]*ent.UserSubscription), args.Error(1)
}

func (m *MockUserSubscriptionRepository) GetByProviderSubscriptionId(ctx context.Context, providerSubscriptionId string, userId string) (*ent.UserSubscription, error) {
	args := m.Called(ctx, providerSubscriptionId, userId)
	return args.Get(0).(*ent.UserSubscription), args.Error(1)
}

type MockSubscriptionRepository struct {
	mock.Mock
}

func (m *MockSubscriptionRepository) GetById(ctx context.Context, id int) (*ent.Subscription, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*ent.Subscription), args.Error(1)
}

func (m *MockSubscriptionRepository) GetAll(ctx context.Context) ([]*ent.Subscription, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*ent.Subscription), args.Error(1)
}

// Mock for GeneratedExamRepositoryInterface
type MockGeneratedExamRepository struct {
	mock.Mock
}

func (m *MockGeneratedExamRepository) AddMany(ctx context.Context, exams []any, ex *ent.Exam) ([]*ent.GeneratedExam, error) {
	args := m.Called(ctx, exams, ex)
	return args.Get(0).([]*ent.GeneratedExam), args.Error(1)
}

func (m *MockGeneratedExamRepository) Add(ctx context.Context, exam map[string]interface{}, examId int) (*ent.GeneratedExam, error) {
	args := m.Called(ctx, exam, examId)
	return args.Get(0).(*ent.GeneratedExam), args.Error(1)
}

func (m *MockGeneratedExamRepository) UpdateMany(ctx context.Context, generatedExams []*ent.GeneratedExam) error {
	args := m.Called(ctx, generatedExams)
	return args.Error(0)
}

func (m *MockGeneratedExamRepository) GetById(ctx context.Context, generatedExamId int) (*ent.GeneratedExam, error) {
	args := m.Called(ctx, generatedExamId)
	return args.Get(0).(*ent.GeneratedExam), args.Error(1)
}

func (m *MockGeneratedExamRepository) GetOpenById(ctx context.Context, generatedExamId int, isOpen bool) (*ent.GeneratedExam, error) {
	args := m.Called(ctx, generatedExamId, isOpen)
	return args.Get(0).(*ent.GeneratedExam), args.Error(1)
}

func (m *MockGeneratedExamRepository) GetActiveById(ctx context.Context, generatedExamId int, isActive bool) (*ent.GeneratedExam, error) {
	args := m.Called(ctx, generatedExamId, isActive)
	return args.Get(0).(*ent.GeneratedExam), args.Error(1)
}

func (m *MockGeneratedExamRepository) GetByExam(ctx context.Context, ex *ent.Exam) ([]*ent.GeneratedExam, error) {
	args := m.Called(ctx, ex)
	return args.Get(0).([]*ent.GeneratedExam), args.Error(1)
}

func (m *MockGeneratedExamRepository) GetByOpenFlag(ctx context.Context, examId int) ([]*ent.GeneratedExam, error) {
	args := m.Called(ctx, examId)
	return args.Get(0).([]*ent.GeneratedExam), args.Error(1)
}

func (m *MockGeneratedExamRepository) GetByMonthOffset(ctx context.Context, ex *ent.Exam, monthOffset, limit int) ([]*ent.GeneratedExam, error) {
	args := m.Called(ctx, ex, monthOffset, limit)
	return args.Get(0).([]*ent.GeneratedExam), args.Error(1)
}

func (m *MockGeneratedExamRepository) GetByWeekOffset(ctx context.Context, ex *ent.Exam, weekOffset, limit int) ([]*ent.GeneratedExam, error) {
	args := m.Called(ctx, ex, weekOffset, limit)
	return args.Get(0).([]*ent.GeneratedExam), args.Error(1)
}

func (m *MockGeneratedExamRepository) GetPaginatedExamsByUserAndDate(ctx context.Context, userId string, page, limit int, from, to *time.Time, examTypeId, categoryID *int) ([]*ent.GeneratedExam, error) {
	args := m.Called(ctx, userId, page, limit, from, to, examTypeId, categoryID)
	return args.Get(0).([]*ent.GeneratedExam), args.Error(1)
}

func (m *MockGeneratedExamRepository) GetCountOfFilteredExamsDataByUserAndDate(ctx context.Context, userId string, from, to *time.Time, examTypeId, categoryID *int) (int, error) {
	args := m.Called(ctx, userId, from, to, examTypeId, categoryID)
	return args.Int(0), args.Error(1)
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

// MockExamGroupRepository is a mock type for the ExamGroupRepositoryInterface
type MockExamGroupRepository struct {
	mock.Mock
}

// GetById mocks the GetById method
func (m *MockExamGroupRepository) GetById(ctx context.Context, examGroupId int) (*ent.ExamGroup, error) {
	args := m.Called(ctx, examGroupId)
	return args.Get(0).(*ent.ExamGroup), args.Error(1)
}

// GetActiveByIdWithExams mocks the GetActiveByIdWithExams method
func (m *MockExamGroupRepository) GetActiveByIdWithExams(ctx context.Context, examGroupId int, isActive bool) (*ent.ExamGroup, error) {
	args := m.Called(ctx, examGroupId, isActive)
	return args.Get(0).(*ent.ExamGroup), args.Error(1)
}

// MockUserRepository is an autogenerated mock type for the UserRepositoryInterface
type MockUserRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, userId
func (m *MockUserRepository) Get(ctx context.Context, userId string) (*ent.User, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).(*ent.User), args.Error(1)
}

// GetByEmail provides a mock function with given fields: ctx, userEmail
func (m *MockUserRepository) GetByEmail(ctx context.Context, userEmail string) (*ent.User, error) {
	args := m.Called(ctx, userEmail)
	return args.Get(0).(*ent.User), args.Error(1)
}

// Update provides a mock function with given fields: ctx, updatedUser
func (m *MockUserRepository) Update(ctx context.Context, updatedUser *ent.User) (*ent.User, error) {
	args := m.Called(ctx, updatedUser)
	return args.Get(0).(*ent.User), args.Error(1)
}

// MockPaymentRepository is an autogenerated mock type for the PaymentRepositoryInterface
type MockPaymentRepository struct {
	mock.Mock
}

// GetByUserId provides a mock function with given fields: ctx, userId
func (m *MockPaymentRepository) GetByUserId(ctx context.Context, userId string) ([]*ent.Payment, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).([]*ent.Payment), args.Error(1)
}

// Create provides a mock function with given fields: ctx, model, userId
func (m *MockPaymentRepository) Create(ctx context.Context, model repositories.CreatePaymentModel, userId string) (*ent.Payment, error) {
	args := m.Called(ctx, model, userId)
	return args.Get(0).(*ent.Payment), args.Error(1)
}

// GetByProviderPaymentId provides a mock function with given fields: ctx, paymentProviderId
func (m *MockPaymentRepository) GetByProviderPaymentId(ctx context.Context, paymentProviderId string) (*ent.Payment, error) {
	args := m.Called(ctx, paymentProviderId)
	return args.Get(0).(*ent.Payment), args.Error(1)
}
