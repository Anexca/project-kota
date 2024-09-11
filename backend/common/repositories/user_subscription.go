package repositories

import (
	"common/ent"
	"context"

	"github.com/google/uuid"
)

type UserSubscriptioRepository struct {
	dbClient *ent.Client
}

type UserSubscriptionModel struct {
	SubscriptionId         int
	UserId                 string
	ProviderSubscriptionId string
	IsActive               bool
}

func NewUserSubscriptioRepository(dbClient *ent.Client) *UserSubscriptioRepository {
	return &UserSubscriptioRepository{
		dbClient: dbClient,
	}
}

func (u *UserSubscriptioRepository) Create(ctx context.Context, model UserSubscriptionModel) (*ent.UserSubscription, error) {
	userUid, err := uuid.Parse(model.UserId)
	if err != nil {
		return nil, err
	}

	return u.dbClient.UserSubscription.Create().
		SetUserID(userUid).
		SetSubscriptionID(model.SubscriptionId).
		SetProviderSubscriptionID(model.ProviderSubscriptionId).
		SetIsActive(false).
		Save(ctx)
}
