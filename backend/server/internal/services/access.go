package services

import (
	"common/ent"
	"common/repositories"
	"context"
	"time"
)

type AccessService struct {
	subscriptionRepository     *repositories.SubscriptionRepository
	userSubscriptionRepository *repositories.UserSubscriptioRepository
}

func NewAccessService(dbClient *ent.Client) *AccessService {
	subscriptionRepository := repositories.NewSubscriptionRepository(dbClient)
	userSubscriptionRepository := repositories.NewUserSubscriptioRepository(dbClient)

	return &AccessService{
		subscriptionRepository:     subscriptionRepository,
		userSubscriptionRepository: userSubscriptionRepository,
	}
}

func (a *AccessService) UserHasAccessToExam(ctx context.Context, examId int, userId string) (bool, error) {
	userSubscriptions, err := a.userSubscriptionRepository.GetByUserId(ctx, userId)
	if err != nil {
		return false, err
	}

	now := time.Now()

	for _, userSubscription := range userSubscriptions {
		if userSubscription.StartDate.Before(now) && userSubscription.EndDate.After(now) {

			subscription, err := a.subscriptionRepository.GetById(ctx, userSubscription.Edges.Subscription.ID)
			if err != nil {
				return false, err
			}

			for _, exam := range subscription.Edges.Exams {
				if exam.ID == examId {
					return true, nil
				}
			}
		}
	}

	return false, nil
}
