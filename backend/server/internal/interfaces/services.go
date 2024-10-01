package interfaces

import (
	"common/constants"
	"common/ent"
	"context"
	"server/pkg/models"

	cashfree_pg "github.com/cashfree/cashfree-pg/v4"
)

// PromptServiceInterface defines the contract for PromptService
type PromptServiceInterface interface {
	GetPromptResult(ctx context.Context, prompt string, model constants.GenAiModel) (string, error)
}

// ExamGenerationServiceInterface defines the contract for ExamGenerationService
type ExamGenerationServiceInterface interface {
	GetGeneratedExamById(ctx context.Context, generatedExamId int, userId string, isOpen bool) (*models.GeneratedExamOverview, error)
}

// AccessServiceInterface defines the interface for the AccessService.
type AccessServiceInterface interface {
	UserHasAccessToExam(ctx context.Context, examId int, userId string) (bool, error)
	GetAccessibleExamsForUser(ctx context.Context, exams []*ent.Exam, userId string) ([]*ent.Exam, error)
}

// ExamAssesmentServiceInterface defines the methods available for the ExamAssesmentService
type ExamAssesmentServiceInterface interface {
	StartNewDescriptiveAssesment(ctx context.Context, generatedExamId int, attempt *ent.ExamAttempt, request *models.DescriptiveExamAssesmentRequest, userId string, isOpen bool) (*models.AssessmentDetails, error)
	GetAssesmentById(ctx context.Context, assessmentId int, userId string) (*models.AssessmentDetails, error)
	GetExamAssessments(ctx context.Context, generatedExamId int, userId string) ([]models.AssessmentDetails, error)
	AssessDescriptiveExam(ctx context.Context, generatedExamId int, assessmentId int, content string, userId string, isOpen bool)
}

// ExamCategoryServiceInterface defines the operations available for an exam category service.
type ExamCategoryServiceInterface interface {
	GetBankingExamGroups(ctx context.Context) ([]models.CategoryExamGroup, error)
	GetExamGroupById(ctx context.Context, examGroupId int) (*models.CategoryExamGroup, error)
}

// PaymentServiceInterface defines the methods for the PaymentService.
type PaymentServiceInterface interface {
	CreateCustomer(model models.UpsertPaymentProviderCustomerModel) (*cashfree_pg.CustomerEntity, error)
	CreateOrder(model models.CreateOrderModel) (*cashfree_pg.OrderEntity, error)
	IsOrderSuccessful(orderId string) (bool, *cashfree_pg.PaymentEntity, error)
	VerifyWebhookSignature(signature, timestamp, body string) (*cashfree_pg.PGWebhookEvent, error)
}

// SubscriptionServiceInterface defines the methods for the subscription service.
type SubscriptionServiceInterface interface {
	GetAll(ctx context.Context) ([]models.SubscriptionOverview, error)
	StartUserSubscription(ctx context.Context, subscriptionId int, returnUrl *string, userId string) (*models.SubscriptionToActivate, error)
	ActivateUserSubscription(ctx context.Context, providerSubscriptionId string, userEmail string) (*models.ActivatedSubscription, error)
	UserHasActiveSubscription(ctx context.Context, subscription *ent.Subscription, user *ent.User) (bool, error)
}

// UserServiceInterface defines the methods for the UserService.
type UserServiceInterface interface {
	GetUserProfile(ctx context.Context, userId string) (models.UserProfileResponse, error)
	UpdateUser(ctx context.Context, userId string, request models.UpdateUserRequest) (*ent.User, error)
	GetUserTransactions(ctx context.Context, userId string) ([]models.SubscriptionPaymentDetails, error)
}
