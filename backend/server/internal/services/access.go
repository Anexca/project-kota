package services

import (
	"context"
	"time"

	"common/ent"
	"common/repositories"
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
				if exam.Edges.Exam.ID == examId {
					return true, nil
				}
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

	now := time.Now()
	var accessibleExams []*ent.Exam

	for _, userSubscription := range userSubscriptions {
		if userSubscription.StartDate.Before(now) && userSubscription.EndDate.After(now) {

			subscription, err := a.subscriptionRepository.GetById(ctx, userSubscription.Edges.Subscription.ID)
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
