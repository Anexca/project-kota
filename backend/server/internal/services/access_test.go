package services_test

import (
	"context"
	"errors"
	"server/internal/services"
	"testing"
	"time"

	"common/ent"
	"common/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
func TestUserHasAccessToExam_Success(t *testing.T) {
	mockSubscriptionRepo := new(MockSubscriptionRepository)
	mockUserSubscriptionRepo := new(MockUserSubscriptionRepository)

	accessService := services.NewAccessService(mockSubscriptionRepo, mockUserSubscriptionRepo)

	ctx := context.Background()
	userId := "test-user-id"
	examId := 1

	// Mock user subscriptions
	userSubscription := &ent.UserSubscription{
		Edges: ent.UserSubscriptionEdges{
			Subscription: &ent.Subscription{ID: 1}, // Subscription ID matches
		},
		StartDate: time.Now().Add(-1 * time.Hour),
		EndDate:   time.Now().Add(1 * time.Hour),
	}
	mockUserSubscriptionRepo.On("GetByUserId", ctx, userId).Return([]*ent.UserSubscription{userSubscription}, nil)

	// Mock subscription exams
	subscription := &ent.Subscription{
		Edges: ent.SubscriptionEdges{
			Exams: []*ent.SubscriptionExam{
				{
					Edges: ent.SubscriptionExamEdges{
						Exam: &ent.Exam{ID: examId}, // Set the exam directly
					},
				},
			},
		},
	}
	mockSubscriptionRepo.On("GetById", ctx, userSubscription.Edges.Subscription.ID).Return(subscription, nil)

	// Execute the method
	hasAccess, err := accessService.UserHasAccessToExam(ctx, examId, userId)
	assert.NoError(t, err)
	assert.True(t, hasAccess)

	// Assert expectations
	mockUserSubscriptionRepo.AssertExpectations(t)
	mockSubscriptionRepo.AssertExpectations(t)
}

func TestUserHasAccessToExam_NoAccess(t *testing.T) {
	mockSubscriptionRepo := new(MockSubscriptionRepository)
	mockUserSubscriptionRepo := new(MockUserSubscriptionRepository)

	accessService := services.NewAccessService(mockSubscriptionRepo, mockUserSubscriptionRepo)

	ctx := context.Background()
	userId := "test-user-id"
	examId := 1

	// Mock user subscriptions
	userSubscription := &ent.UserSubscription{
		Edges: ent.UserSubscriptionEdges{
			Subscription: &ent.Subscription{ID: 1}, // Subscription ID matches
		},
		StartDate: time.Now().Add(-1 * time.Hour),
		EndDate:   time.Now().Add(1 * time.Hour),
	}
	mockUserSubscriptionRepo.On("GetByUserId", ctx, userId).Return([]*ent.UserSubscription{userSubscription}, nil)

	// Mock subscription exams with no matching exam
	subscription := &ent.Subscription{
		Edges: ent.SubscriptionEdges{
			Exams: []*ent.SubscriptionExam{

				{
					Edges: ent.SubscriptionExamEdges{
						Exam: &ent.Exam{ID: 2}, // Reference to another Exam
					},
				},
			},
		},
	}
	mockSubscriptionRepo.On("GetById", ctx, userSubscription.Edges.Subscription.ID).Return(subscription, nil)

	// Execute the method
	hasAccess, err := accessService.UserHasAccessToExam(ctx, examId, userId)
	assert.NoError(t, err)
	assert.False(t, hasAccess)

	// Assert expectations
	mockUserSubscriptionRepo.AssertExpectations(t)
	mockSubscriptionRepo.AssertExpectations(t)
}

func TestGetAccessibleExamsForUser_Success(t *testing.T) {
	mockSubscriptionRepo := new(MockSubscriptionRepository)
	mockUserSubscriptionRepo := new(MockUserSubscriptionRepository)

	accessService := services.NewAccessService(mockSubscriptionRepo, mockUserSubscriptionRepo)

	ctx := context.Background()
	userId := "test-user-id"

	// Mock user subscriptions
	userSubscription := &ent.UserSubscription{
		Edges: ent.UserSubscriptionEdges{
			Subscription: &ent.Subscription{ID: 1},
		},
		StartDate: time.Now().Add(-1 * time.Hour),
		EndDate:   time.Now().Add(1 * time.Hour),
	}
	mockUserSubscriptionRepo.On("GetByUserId", ctx, userId).Return([]*ent.UserSubscription{userSubscription}, nil)

	// Mock subscription exams
	subscription := &ent.Subscription{
		Edges: ent.SubscriptionEdges{
			Exams: []*ent.SubscriptionExam{
				{
					Edges: ent.SubscriptionExamEdges{
						Exam: &ent.Exam{ID: 1}, // Reference to the Exam
					},
				},
				{
					Edges: ent.SubscriptionExamEdges{
						Exam: &ent.Exam{ID: 2}, // Reference to another Exam
					},
				},
			},
		},
	}
	mockSubscriptionRepo.On("GetById", ctx, userSubscription.Edges.Subscription.ID).Return(subscription, nil)

	// Mock available exams
	exams := []*ent.Exam{
		{ID: 1, Name: "Exam 1"},
		{ID: 2, Name: "Exam 2"},
		{ID: 3, Name: "Exam 3"},
	}

	// Execute the method
	accessibleExams, err := accessService.GetAccessibleExamsForUser(ctx, exams, userId)
	assert.NoError(t, err)
	assert.Len(t, accessibleExams, 2) // Should only include exams 1 and 2

	// Assert expectations
	mockUserSubscriptionRepo.AssertExpectations(t)
	mockSubscriptionRepo.AssertExpectations(t)
}

func TestGetAccessibleExamsForUser_NoAccessibleExams(t *testing.T) {
	mockSubscriptionRepo := new(MockSubscriptionRepository)
	mockUserSubscriptionRepo := new(MockUserSubscriptionRepository)

	accessService := services.NewAccessService(mockSubscriptionRepo, mockUserSubscriptionRepo)

	ctx := context.Background()
	userId := "test-user-id"

	// Mock user subscriptions
	userSubscription := &ent.UserSubscription{
		Edges: ent.UserSubscriptionEdges{
			Subscription: &ent.Subscription{ID: 1},
		},
		StartDate: time.Now().Add(-1 * time.Hour),
		EndDate:   time.Now().Add(1 * time.Hour),
	}
	mockUserSubscriptionRepo.On("GetByUserId", ctx, userId).Return([]*ent.UserSubscription{userSubscription}, nil)

	// Mock subscription exams with no matching exams
	subscription := &ent.Subscription{
		Edges: ent.SubscriptionEdges{
			Exams: []*ent.SubscriptionExam{
				{
					Edges: ent.SubscriptionExamEdges{
						Exam: &ent.Exam{ID: 4}, // Non-matching exam ID
					},
				},
			},
		},
	}
	mockSubscriptionRepo.On("GetById", ctx, userSubscription.Edges.Subscription.ID).Return(subscription, nil)

	// Mock available exams that do not match subscription exams
	exams := []*ent.Exam{
		{ID: 1, Name: "Exam 1"},
		{ID: 2, Name: "Exam 2"},
		{ID: 3, Name: "Exam 3"}, // None of these match exam ID 4
	}

	// Execute the method
	accessibleExams, err := accessService.GetAccessibleExamsForUser(ctx, exams, userId)
	assert.NoError(t, err)
	assert.Len(t, accessibleExams, 0) // Should return no accessible exams

	// Assert expectations
	mockUserSubscriptionRepo.AssertExpectations(t)
	mockSubscriptionRepo.AssertExpectations(t)
}

func TestUserHasAccessToExam_ErrorFetchingSubscriptions(t *testing.T) {
	mockSubscriptionRepo := new(MockSubscriptionRepository)
	mockUserSubscriptionRepo := new(MockUserSubscriptionRepository)

	accessService := services.NewAccessService(mockSubscriptionRepo, mockUserSubscriptionRepo)

	ctx := context.Background()
	userId := "test-user-id"
	examId := 1

	// Simulate an error when fetching user subscriptions
	mockUserSubscriptionRepo.On("GetByUserId", ctx, userId).Return([]*ent.UserSubscription{}, errors.New("fetch error"))

	hasAccess, err := accessService.UserHasAccessToExam(ctx, examId, userId)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fetch error")
	assert.False(t, hasAccess)

	// Assert expectations
	mockUserSubscriptionRepo.AssertExpectations(t)
}

func TestGetAccessibleExamsForUser_ErrorFetchingSubscriptions(t *testing.T) {
	mockSubscriptionRepo := new(MockSubscriptionRepository)
	mockUserSubscriptionRepo := new(MockUserSubscriptionRepository)

	accessService := services.NewAccessService(mockSubscriptionRepo, mockUserSubscriptionRepo)

	ctx := context.Background()
	userId := "test-user-id"

	// Simulate an error when fetching user subscriptions
	mockUserSubscriptionRepo.On("GetByUserId", ctx, userId).Return([]*ent.UserSubscription{}, errors.New("fetch error"))

	exams := []*ent.Exam{}
	accessibleExams, err := accessService.GetAccessibleExamsForUser(ctx, exams, userId)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fetch error")
	assert.Nil(t, accessibleExams)

	// Assert expectations
	mockUserSubscriptionRepo.AssertExpectations(t)
}

func TestUserHasAccessToExam_ErrorFetchingSubscription(t *testing.T) {
	mockSubscriptionRepo := new(MockSubscriptionRepository)
	mockUserSubscriptionRepo := new(MockUserSubscriptionRepository)

	accessService := services.NewAccessService(mockSubscriptionRepo, mockUserSubscriptionRepo)

	ctx := context.Background()
	userId := "test-user-id"
	examId := 1

	// Mock user subscriptions
	mockUserSubscriptionRepo.On("GetByUserId", ctx, userId).Return([]*ent.UserSubscription{}, nil)

	// Execute the method
	hasAccess, err := accessService.UserHasAccessToExam(ctx, examId, userId)
	assert.NoError(t, err)
	assert.False(t, hasAccess)

	// Assert expectations
	mockUserSubscriptionRepo.AssertExpectations(t)
}
