package repositories

import (
	"context"

	"github.com/google/uuid"

	"common/ent"
	"common/ent/user"
	"common/ent/usersubscription"
)

// UserSubscriptionRepositoryInterface defines the contract for the UserSubscription repository.
type UserSubscriptionRepositoryInterface interface {
	Create(ctx context.Context, model UserSubscriptionModel) (*ent.UserSubscription, error)
	Update(ctx context.Context, updatedUserSubscription *ent.UserSubscription) error
	GetById(ctx context.Context, userSubscriptionId int, userId string) (*ent.UserSubscription, error)
	GetByUserId(ctx context.Context, userId string) ([]*ent.UserSubscription, error)
	GetByProviderSubscriptionId(ctx context.Context, providerSubscriptionId string, userId string) (*ent.UserSubscription, error)
}

// UserSubscriptionModel is the model to create a new user subscription.
type UserSubscriptionModel struct {
	SubscriptionId         int
	UserId                 string
	ProviderSubscriptionId string
	IsActive               bool
}

// UserSubscriptionRepository is a concrete implementation of UserSubscriptionRepositoryInterface.
type UserSubscriptionRepository struct {
	dbClient *ent.Client
}

// NewUserSubscriptionRepository creates a new instance of UserSubscriptionRepository.
func NewUserSubscriptionRepository(dbClient *ent.Client) *UserSubscriptionRepository {
	return &UserSubscriptionRepository{
		dbClient: dbClient,
	}
}

// Create adds a new user subscription record to the database.
func (u *UserSubscriptionRepository) Create(ctx context.Context, model UserSubscriptionModel) (*ent.UserSubscription, error) {
	userUid, err := uuid.Parse(model.UserId)
	if err != nil {
		return nil, err
	}

	return u.dbClient.UserSubscription.Create().
		SetUserID(userUid).
		SetSubscriptionID(model.SubscriptionId).
		SetProviderSubscriptionID(model.ProviderSubscriptionId).
		SetIsActive(model.IsActive).
		Save(ctx)
}

// Update modifies an existing user subscription in the database.
func (u *UserSubscriptionRepository) Update(ctx context.Context, updatedUserSubscription *ent.UserSubscription) error {
	return u.dbClient.UserSubscription.Update().
		Where(usersubscription.IDEQ(updatedUserSubscription.ID)).
		SetIsActive(updatedUserSubscription.IsActive).
		SetStartDate(updatedUserSubscription.StartDate).
		SetEndDate(updatedUserSubscription.EndDate).
		SetStatus(updatedUserSubscription.Status).
		Exec(ctx)
}

// GetById retrieves a user subscription by its ID and the user's ID.
func (u *UserSubscriptionRepository) GetById(ctx context.Context, userSubscriptionId int, userId string) (*ent.UserSubscription, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return u.dbClient.UserSubscription.Query().
		Where(
			usersubscription.IDEQ(userSubscriptionId),
			usersubscription.HasUserWith(user.IDEQ(userUid)),
		).
		WithSubscription().
		Only(ctx)
}

// GetByUserId retrieves all active user subscriptions for a given user.
func (u *UserSubscriptionRepository) GetByUserId(ctx context.Context, userId string) ([]*ent.UserSubscription, error) {
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

// GetByProviderSubscriptionId retrieves a user subscription by the provider's subscription ID and the user's ID.
func (u *UserSubscriptionRepository) GetByProviderSubscriptionId(ctx context.Context, providerSubscriptionId string, userId string) (*ent.UserSubscription, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return u.dbClient.UserSubscription.Query().
		Where(usersubscription.ProviderSubscriptionIDEQ(providerSubscriptionId),
			usersubscription.HasUserWith(user.IDEQ(userUid)),
		).
		WithSubscription().
		Only(ctx)
}
