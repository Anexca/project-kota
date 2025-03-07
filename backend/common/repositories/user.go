package repositories

import (
	"context"

	"github.com/google/uuid"

	"common/ent"
	"common/ent/user"
)

// UpdateUserModel represents the model for updating a user's details.
type UpdateUserModel struct {
	FirstName                 string
	LastName                  string
	PhoneNumber               string
	PaymentProviderCustomerID string
}

// UserRepository is a concrete implementation of UserRepositoryInterface.
type UserRepository struct {
	dbClient *ent.Client
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(dbClient *ent.Client) *UserRepository {
	return &UserRepository{
		dbClient: dbClient,
	}
}

// Get retrieves a user by their ID.
func (u *UserRepository) Get(ctx context.Context, userId string) (*ent.User, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return u.dbClient.User.Query().Where(user.IDEQ(userUid)).Only(ctx)
}

// GetByEmail retrieves a user by their email address.
func (u *UserRepository) GetByEmail(ctx context.Context, userEmail string) (*ent.User, error) {
	return u.dbClient.User.Query().Where(user.EmailEQ(userEmail)).Only(ctx)
}

// Update updates a user's details in the database.
func (u *UserRepository) Update(ctx context.Context, updatedUser *ent.User) (*ent.User, error) {
	return u.dbClient.User.
		UpdateOneID(updatedUser.ID).
		SetFirstName(updatedUser.FirstName).
		SetLastName(updatedUser.LastName).
		SetPhoneNumber(updatedUser.PhoneNumber).
		SetPaymentProviderCustomerID(updatedUser.PaymentProviderCustomerID).
		Save(ctx)
}
