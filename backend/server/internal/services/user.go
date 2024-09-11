package services

import (
	"common/ent"
	"common/repositories"
	"context"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

type UpdateUserRequest struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone,omitempty"`
}

func NewUserService(dbClient *ent.Client) *UserService {
	userRepository := repositories.NewUserRepository(dbClient)

	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) GetUser(ctx context.Context, userId string) (*ent.User, error) {
	return u.userRepository.Get(ctx, userId)
}

func (u *UserService) UpdateUser(ctx context.Context, userId string, request UpdateUserRequest) (*ent.User, error) {
	userToUpdate, err := u.GetUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	if request.FirstName != "" {
		userToUpdate.FirstName = request.FirstName
	}

	if request.LastName != "" {
		userToUpdate.LastName = request.LastName
	}

	if request.PhoneNumber != "" {
		userToUpdate.PhoneNumber = request.PhoneNumber
	}

	updatedUser, err := u.userRepository.Update(ctx, userToUpdate)
	return updatedUser, err
}
