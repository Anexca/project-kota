package repositories

import (
	"common/ent"
	"common/ent/user"
	"context"

	"github.com/google/uuid"
)

type UserRepository struct {
	dbClient *ent.Client
}

type UpdateUserModel struct {
	FirstName                 string
	LastName                  string
	PhoneNumber               string
	PaymentProviderCustomerID string
}

func NewUserRepository(dbClient *ent.Client) *UserRepository {
	return &UserRepository{
		dbClient: dbClient,
	}
}

func (u *UserRepository) Get(ctx context.Context, userId string) (*ent.User, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return u.dbClient.User.Query().Where(user.IDEQ(userUid)).Only(ctx)
}

func (u *UserRepository) GetByEmail(ctx context.Context, userEmail string) (*ent.User, error) {
	return u.dbClient.User.Query().Where(user.EmailEQ(userEmail)).Only(ctx)
}

func (u *UserRepository) Update(ctx context.Context, updatedUser *ent.User) (*ent.User, error) {
	return u.dbClient.User.
		UpdateOneID(updatedUser.ID).
		SetFirstName(updatedUser.FirstName).
		SetLastName(updatedUser.LastName).
		SetPhoneNumber(updatedUser.PhoneNumber).
		SetPaymentProviderCustomerID(updatedUser.PaymentProviderCustomerID).
		Save(ctx)
}
