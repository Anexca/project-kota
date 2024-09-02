package services

import (
	"context"

	"github.com/nedpals/supabase-go"
)

type AuthService struct {
	supabaseClient *supabase.Client
}

func NewAuthService(supabaseClient *supabase.Client) *AuthService {
	return &AuthService{
		supabaseClient: supabaseClient,
	}
}

func (a *AuthService) VerifyUserToken(ctx context.Context, userToken string) (*supabase.User, error) {
	return a.supabaseClient.Auth.User(ctx, userToken)
}
