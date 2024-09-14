package repositories

import (
	"common/ent"
	"common/ent/user"
	"common/ent/usersubscription"
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

func (u *UserSubscriptioRepository) Update(ctx context.Context, updatedUserSubscription *ent.UserSubscription) error {
	return u.dbClient.UserSubscription.Update().
		Where(usersubscription.IDEQ(updatedUserSubscription.ID)).
		SetIsActive(updatedUserSubscription.IsActive).
		SetStartDate(updatedUserSubscription.StartDate).
		SetEndDate(updatedUserSubscription.EndDate).
		SetStatus(updatedUserSubscription.Status).
		Exec(ctx)
}

func (u *UserSubscriptioRepository) GetById(ctx context.Context, userSubscriptionId int, userId string) (*ent.UserSubscription, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return u.dbClient.UserSubscription.Query().
		Where(
			usersubscription.IDEQ(userSubscriptionId),
			usersubscription.HasUserWith(user.IDEQ(userUid))).
		WithSubscription().
		Only(ctx)
}

func (u *UserSubscriptioRepository) GetByUserId(ctx context.Context, userId string) ([]*ent.UserSubscription, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return u.dbClient.UserSubscription.Query().
		Where(usersubscription.HasUserWith(user.IDEQ(userUid)),
			usersubscription.IsActiveEQ(true),
			usersubscription.StatusEQ(usersubscription.StatusACTIVE),
		).
		WithSubscription().
		WithPayments().
		All(ctx)
}
