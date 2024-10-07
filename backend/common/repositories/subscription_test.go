package repositories_test

import (
	"context"
	"testing"

	"common/repositories"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestSubscriptionRepository_GetAll(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	_, err = client.Subscription.Create().
		SetProviderPlanID("basic-plan-123").
		SetName("Basic Plan").
		SetDurationInMonths(6).
		SetIsActive(true).
		Save(ctx)
	require.NoError(t, err)

	_, err = client.Subscription.Create().
		SetProviderPlanID("premium-plan-456").
		SetName("Premium Plan").
		SetDurationInMonths(12).
		SetIsActive(true).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewSubscriptionRepository(client)

	subscriptions, err := repo.GetAll(ctx)
	require.NoError(t, err)
	require.Len(t, subscriptions, 2)
}

func TestSubscriptionRepository_GetById(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	_, err = client.User.Create().
		SetID(uuid.New()).
		SetEmail("testuser@example.com").
		Save(ctx)
	require.NoError(t, err)

	exam, err := client.Exam.Create().
		SetName("Mock Exam").
		SetDescription("A mock exam description").
		Save(ctx)
	require.NoError(t, err)

	subscription, err := client.Subscription.Create().
		SetProviderPlanID("premium-plan-789").
		SetName("Premium Plan").
		SetDurationInMonths(12).
		SetIsActive(true).
		Save(ctx)
	require.NoError(t, err)

	_, err = client.SubscriptionExam.Create().
		SetSubscription(subscription).
		SetExam(exam).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewSubscriptionRepository(client)

	retrievedSub, err := repo.GetById(ctx, subscription.ID)
	require.NoError(t, err)
	require.Equal(t, subscription.Name, retrievedSub.Name)
	require.Len(t, retrievedSub.Edges.Exams, 1)
	require.Equal(t, exam.Name, retrievedSub.Edges.Exams[0].Edges.Exam.Name)
}

func TestSubscriptionRepository_GetById_InvalidId(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	repo := repositories.NewSubscriptionRepository(client)

	_, err = repo.GetById(ctx, 999)
	require.Error(t, err)
}
