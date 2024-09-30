package repositories_test

import (
	"context"
	"testing"
	"time"

	"common/ent/usersubscription"
	"common/repositories"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestUserSubscriptionRepository_CreateAndUpdate(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	userID := uuid.New()
	subscriptionID := 1

	_, err = client.User.Create().
		SetID(userID).
		SetEmail("test@example.com").
		Save(ctx)
	require.NoError(t, err)

	// Add required fields for Subscription, including duration_in_months
	_, err = client.Subscription.Create().
		SetProviderPlanID("pro-plan-123").
		SetName("Pro Plan").
		SetDurationInMonths(12).
		SetIsActive(true).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewUserSubscriptionRepository(client)

	model := repositories.UserSubscriptionModel{
		SubscriptionId:         subscriptionID,
		UserId:                 userID.String(),
		ProviderSubscriptionId: "provider-id-123",
		IsActive:               true,
	}
	userSub, err := repo.Create(ctx, model)
	require.NoError(t, err)
	require.NotNil(t, userSub)
	require.Equal(t, "provider-id-123", userSub.ProviderSubscriptionID)
	require.True(t, userSub.IsActive)

	userSub.IsActive = false
	userSub.StartDate = time.Now()
	userSub.EndDate = time.Now().Add(30 * time.Hour)
	userSub.Status = usersubscription.StatusACTIVE

	err = repo.Update(ctx, userSub)
	require.NoError(t, err)

	updatedSub, err := client.UserSubscription.Get(ctx, userSub.ID)
	require.NoError(t, err)
	require.Equal(t, false, updatedSub.IsActive)
	require.Equal(t, usersubscription.StatusACTIVE, updatedSub.Status)
	require.NotNil(t, updatedSub.StartDate)
	require.NotNil(t, updatedSub.EndDate)
}

func TestUserSubscriptionRepository_GetByIdAndFailures(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	userID := uuid.New()
	subscriptionID := 1

	_, err = client.User.Create().
		SetID(userID).
		SetEmail("test@example.com").
		Save(ctx)
	require.NoError(t, err)

	_, err = client.Subscription.Create().
		SetProviderPlanID("pro-plan-123").
		SetName("Pro Plan").
		SetDurationInMonths(12).
		SetIsActive(true).
		Save(ctx)
	require.NoError(t, err)

	userSub, err := client.UserSubscription.Create().
		SetUserID(userID).
		SetSubscriptionID(subscriptionID).
		SetProviderSubscriptionID("provider-id-123").
		SetIsActive(true).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewUserSubscriptionRepository(client)

	retrievedSub, err := repo.GetById(ctx, userSub.ID, userID.String())
	require.NoError(t, err)
	require.NotNil(t, retrievedSub)
	require.Equal(t, "provider-id-123", retrievedSub.ProviderSubscriptionID)
	require.Equal(t, userSub.ID, retrievedSub.ID)

	invalidUserId := uuid.New().String()
	_, err = repo.GetById(ctx, userSub.ID, invalidUserId)
	require.Error(t, err)
	require.Contains(t, err.Error(), "user_subscription not found")
}

func TestUserSubscriptionRepository_GetByProviderSubscriptionId_InvalidData(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	userID := uuid.New()
	subscriptionID := 1

	_, err = client.User.Create().
		SetID(userID).
		SetEmail("test@example.com").
		Save(ctx)
	require.NoError(t, err)

	_, err = client.Subscription.Create().
		SetProviderPlanID("pro-plan-123").
		SetName("Pro Plan").
		SetDurationInMonths(12).
		SetIsActive(true).
		Save(ctx)
	require.NoError(t, err)

	_, err = client.UserSubscription.Create().
		SetUserID(userID).
		SetSubscriptionID(subscriptionID).
		SetProviderSubscriptionID("provider-id-123").
		SetIsActive(true).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewUserSubscriptionRepository(client)

	_, err = repo.GetByProviderSubscriptionId(ctx, "invalid-provider-id", userID.String())
	require.Error(t, err)
	require.Contains(t, err.Error(), "user_subscription not found")
}

func TestUserSubscriptionRepository_GetByUserId_NoSubscriptions(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	userID := uuid.New()

	_, err = client.User.Create().
		SetID(userID).
		SetEmail("test@example.com").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewUserSubscriptionRepository(client)

	subscriptions, err := repo.GetByUserId(ctx, userID.String())
	require.NoError(t, err)
	require.Len(t, subscriptions, 0)
}
