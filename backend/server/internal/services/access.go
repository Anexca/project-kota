package services

import (
	"context"
	"time"

	"common/ent"
	commonInterfaces "common/interfaces"
	"common/repositories"
)

type AccessService struct {
	subscriptionRepository     commonInterfaces.SubscriptionRepositoryInterface
	userSubscriptionRepository commonInterfaces.UserSubscriptionRepositoryInterface
}

func NewAccessService(subscriptionRepo commonInterfaces.SubscriptionRepositoryInterface, userSubscriptionRepo commonInterfaces.UserSubscriptionRepositoryInterface) *AccessService {
	return &AccessService{
		subscriptionRepository:     subscriptionRepo,
		userSubscriptionRepository: userSubscriptionRepo,
	}
}

func InitAccessService(dbClient *ent.Client) *AccessService {
	subscriptionRepository := repositories.NewSubscriptionRepository(dbClient)
	userSubscriptionRepository := repositories.NewUserSubscriptionRepository(dbClient)

	return NewAccessService(subscriptionRepository, userSubscriptionRepository)
}

func (a *AccessService) UserHasAccessToExam(ctx context.Context, examId int, userId string) (bool, error) {
	userSubscriptions, err := a.userSubscriptionRepository.GetByUserId(ctx, userId)
	if err != nil {
		return false, err
	}

	if len(userSubscriptions) == 0 {
		return false, nil // No subscriptions found
	}

	now := time.Now()

	for _, userSubscription := range userSubscriptions {
		if userSubscription.StartDate.Before(now) && userSubscription.EndDate.After(now) {
			subscription, err := a.subscriptionRepository.GetById(ctx, userSubscription.Edges.Subscription.ID)
			if err != nil {
				return false, err
			}

			// Create a map of exam IDs for quick lookup
			examMap := make(map[int]struct{})
			for _, exam := range subscription.Edges.Exams {
				examMap[exam.Edges.Exam.ID] = struct{}{}
			}

			// Check if the requested examId exists in the subscription exams
			if _, exists := examMap[examId]; exists {
				return true, nil
			}
		}
	}

	return false, nil
}

func (a *AccessService) GetAccessibleExamsForUser(ctx context.Context, exams []*ent.Exam, userId string) ([]*ent.Exam, error) {
	userSubscriptions, err := a.userSubscriptionRepository.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	subscriptions, err := a.subscriptionRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	var accessibleExams []*ent.Exam

	for _, userSubscription := range userSubscriptions {
		if userSubscription.StartDate.Before(now) && userSubscription.EndDate.After(now) {

			subscription, err := a.GetSubscriptionById(subscriptions, userSubscription.Edges.Subscription.ID)
			if err != nil {
				return nil, err
			}

			examMap := make(map[int]struct{})
			for _, exam := range subscription.Edges.Exams {
				examMap[exam.Edges.Exam.ID] = struct{}{}
			}

			for _, exam := range exams {
				if _, found := examMap[exam.ID]; found {
					accessibleExams = append(accessibleExams, exam)
				}
			}
		}
	}

	return accessibleExams, nil
}

func (a *AccessService) GetSubscriptionById(subscriptions []*ent.Subscription, subscriptionId int) (*ent.Subscription, error) {
	for _, subscription := range subscriptions {
		if subscription.ID == subscriptionId {
			return subscription, nil
		}
	}

	return nil, &ent.NotFoundError{}
}
