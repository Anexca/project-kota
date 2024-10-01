package interfaces

import (
	"common/constants"
	"common/ent"
	"common/repositories"
	"context"
	"time"
)

// CachedExamRepositoryInterface defines the contract for the cached exam repository.
type CachedExamRepositoryInterface interface {
	Create(ctx context.Context, cacheUID string, expiry time.Duration, exam *ent.Exam) (*ent.CachedExam, error)
	GetByExam(ctx context.Context, ex *ent.Exam) ([]*ent.CachedExam, error)
	MarkAsUsed(ctx context.Context, id int) error
}

// ExamAssessmentRepositoryInterface defines the contract for interacting with exam assessments.
type ExamAssessmentRepositoryInterface interface {
	Create(ctx context.Context, attemptId int, model repositories.AssessmentModel) (*ent.ExamAssesment, error)
	Update(ctx context.Context, assessmentId int, model repositories.AssessmentModel) error
	GetById(ctx context.Context, assessmentId int, userId string) (*ent.ExamAssesment, error)
	GetByExam(ctx context.Context, generatedExamId int, userId string) ([]*ent.ExamAssesment, error)
}

// ExamAttemptRepositoryInterface defines the contract for the exam attempt repository.
type ExamAttemptRepositoryInterface interface {
	GetById(ctx context.Context, attemptId int, userId string) (*ent.ExamAttempt, error)
	GetByUserId(ctx context.Context, userId string) ([]*ent.ExamAttempt, error)
	GetByExam(ctx context.Context, generatedExamId int, userId string) ([]*ent.ExamAttempt, error)
	Create(ctx context.Context, currentAttempt int, generatedExamId int, userId string) (*ent.ExamAttempt, error)
}

// ExamCategoryRepositoryInterface defines the contract for the exam category repository.
type ExamCategoryRepositoryInterface interface {
	Get(ctx context.Context) ([]*ent.ExamCategory, error)
	GetByName(ctx context.Context, categoryName constants.ExamCategoryName) (*ent.ExamCategory, error)
}

// ExamGroupRepositoryInterface defines the contract for the exam group repository.
type ExamGroupRepositoryInterface interface {
	GetById(ctx context.Context, examGroupId int) (*ent.ExamGroup, error)
	GetActiveByIdWithExams(ctx context.Context, examGroupId int, isActive bool) (*ent.ExamGroup, error)
}

// ExamSettingRepositoryInterface defines the contract for the exam setting repository.
type ExamSettingRepositoryInterface interface {
	GetByExam(ctx context.Context, examId int) (*ent.ExamSetting, error)
}

// ExamRepositoryInterface defines the contract for the exam repository.
type ExamRepositoryInterface interface {
	GetById(ctx context.Context, examId int) (*ent.Exam, error)
	GetActiveByExamsGroupId(ctx context.Context, examGroupId int, isActive bool) ([]*ent.Exam, error)
	GetActiveById(ctx context.Context, examId int, isActive bool) (*ent.Exam, error)
	GetByExamCategory(ctx context.Context, examCategory *ent.ExamCategory) ([]*ent.Exam, error)
	GetActiveByType(ctx context.Context, examType constants.ExamType) ([]*ent.Exam, error)
	GetByName(ctx context.Context, examName string) (*ent.Exam, error)
}

// GeneratedExamRepositoryInterface defines the contract for the GeneratedExam repository.
type GeneratedExamRepositoryInterface interface {
	AddMany(ctx context.Context, exams []any, ex *ent.Exam) ([]*ent.GeneratedExam, error)
	Add(ctx context.Context, exam map[string]interface{}, examId int) (*ent.GeneratedExam, error)
	UpdateMany(ctx context.Context, generatedExams []*ent.GeneratedExam) error
	GetById(ctx context.Context, generatedExamId int) (*ent.GeneratedExam, error)
	GetOpenById(ctx context.Context, generatedExamId int, isOpen bool) (*ent.GeneratedExam, error)
	GetActiveById(ctx context.Context, generatedExamId int, isActive bool) (*ent.GeneratedExam, error)
	GetByExam(ctx context.Context, ex *ent.Exam) ([]*ent.GeneratedExam, error)
	GetByOpenFlag(ctx context.Context, examId int) ([]*ent.GeneratedExam, error)
	GetByMonthOffset(ctx context.Context, ex *ent.Exam, monthOffset, limit int) ([]*ent.GeneratedExam, error)
	GetByWeekOffset(ctx context.Context, ex *ent.Exam, weekOffset, limit int) ([]*ent.GeneratedExam, error)
	GetPaginatedExamsByUserAndDate(ctx context.Context, userId string, page, limit int, from, to *time.Time, examTypeId, categoryID *int) ([]*ent.GeneratedExam, error)
	GetCountOfFilteredExamsDataByUserAndDate(ctx context.Context, userId string, from, to *time.Time, examTypeId, categoryID *int) (int, error)
}

// PaymentRepositoryInterface defines the contract for the Payment repository.
type PaymentRepositoryInterface interface {
	GetByUserId(ctx context.Context, userId string) ([]*ent.Payment, error)
	Create(ctx context.Context, model repositories.CreatePaymentModel, userId string) (*ent.Payment, error)
	GetByProviderPaymentId(ctx context.Context, paymentProviderId string) (*ent.Payment, error)
}

// SubscriptionRepositoryInterface defines the contract for the Subscription repository.
type SubscriptionRepositoryInterface interface {
	GetAll(ctx context.Context) ([]*ent.Subscription, error)
	GetById(ctx context.Context, subscriptionId int) (*ent.Subscription, error)
}

// UserSubscriptionRepositoryInterface defines the contract for the UserSubscription repository.
type UserSubscriptionRepositoryInterface interface {
	Create(ctx context.Context, model repositories.UserSubscriptionModel) (*ent.UserSubscription, error)
	Update(ctx context.Context, updatedUserSubscription *ent.UserSubscription) error
	GetById(ctx context.Context, userSubscriptionId int, userId string) (*ent.UserSubscription, error)
	GetByUserId(ctx context.Context, userId string) ([]*ent.UserSubscription, error)
	GetByProviderSubscriptionId(ctx context.Context, providerSubscriptionId string, userId string) (*ent.UserSubscription, error)
}

// UserRepositoryInterface defines the contract for the User repository.
type UserRepositoryInterface interface {
	Get(ctx context.Context, userId string) (*ent.User, error)
	GetByEmail(ctx context.Context, userEmail string) (*ent.User, error)
	Update(ctx context.Context, updatedUser *ent.User) (*ent.User, error)
}
