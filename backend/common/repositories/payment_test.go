package repositories_test

import (
	"context"
	"testing"
	"time"

	"common/ent/payment"
	"common/repositories"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestPaymentRepository_GetByUserId(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	userID := uuid.New()

	// Create a user
	_, err = client.User.Create().
		SetID(userID).
		SetEmail("test@example.com").
		Save(ctx)
	require.NoError(t, err)

	// Create a subscription
	subscription, err := client.Subscription.Create().
		SetProviderPlanID("basic-plan").
		SetName("Basic Plan").
		SetDurationInMonths(6).
		SetIsActive(true).
		Save(ctx)
	require.NoError(t, err)

	// Create a UserSubscription to link with Payment
	userSubscription, err := client.UserSubscription.Create().
		SetUserID(userID).
		SetSubscriptionID(subscription.ID).
		SetProviderSubscriptionID("provider-subscription-id-123").
		SetIsActive(true).
		Save(ctx)
	require.NoError(t, err)

	// Create a payment linked to the UserSubscription
	_, err = client.Payment.Create().
		SetUserID(userID).
		SetSubscriptionID(userSubscription.ID). // Link the Payment to UserSubscription
		SetAmount(100.50).
		SetPaymentMethod("Credit Card").
		SetStatus(payment.StatusSUCCESS).
		SetProviderPaymentID("payment-123").
		SetPaymentDate(time.Now()).
		Save(ctx)
	require.NoError(t, err)

	// Now retrieve payments by user ID
	repo := repositories.NewPaymentRepository(client)

	payments, err := repo.GetByUserId(ctx, userID.String())
	require.NoError(t, err)
	require.Len(t, payments, 1)
	require.Equal(t, "payment-123", payments[0].ProviderPaymentID)
}

func TestPaymentRepository_Create(t *testing.T) {
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

	subscription, err := client.Subscription.Create().
		SetProviderPlanID("pro-plan").
		SetName("Pro Plan").
		SetDurationInMonths(12).
		SetIsActive(true).
		Save(ctx)
	require.NoError(t, err)

	// Create UserSubscription
	userSubscription, err := client.UserSubscription.Create().
		SetUserID(userID).
		SetSubscriptionID(subscription.ID).
		SetProviderSubscriptionID("provider-sub-id-123").
		SetIsActive(true).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewPaymentRepository(client)

	model := repositories.CreatePaymentModel{
		UserSubscriptionId: userSubscription.ID, // Use UserSubscription ID
		Amount:             250.75,
		PaymentDate:        time.Now(),
		Status:             string(payment.StatusPENDING),
		PaymentMethod:      "Debit Card",
		ProviderPaymentId:  "provider-payment-456",
	}

	payment, err := repo.Create(ctx, model, userID.String())
	require.NoError(t, err)
	require.Equal(t, "provider-payment-456", payment.ProviderPaymentID)
	require.Equal(t, model.Amount, payment.Amount)
}

func TestPaymentRepository_GetByProviderPaymentId(t *testing.T) {
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

	subscription, err := client.Subscription.Create().
		SetProviderPlanID("premium-plan").
		SetName("Premium Plan").
		SetDurationInMonths(12).
		SetIsActive(true).
		Save(ctx)
	require.NoError(t, err)

	// Create UserSubscription
	userSubscription, err := client.UserSubscription.Create().
		SetUserID(userID).
		SetSubscriptionID(subscription.ID).
		SetProviderSubscriptionID("provider-sub-id-456").
		SetIsActive(true).
		Save(ctx)
	require.NoError(t, err)

	_, err = client.Payment.Create().
		SetUserID(userID).
		SetSubscriptionID(userSubscription.ID). // Use UserSubscription ID
		SetAmount(500.00).
		SetPaymentMethod("Bank Transfer").
		SetStatus(payment.StatusSUCCESS).
		SetProviderPaymentID("provider-payment-789").
		SetPaymentDate(time.Now()).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewPaymentRepository(client)

	payment, err := repo.GetByProviderPaymentId(ctx, "provider-payment-789")
	require.NoError(t, err)
	require.Equal(t, "provider-payment-789", payment.ProviderPaymentID)
	require.Equal(t, 500.00, payment.Amount)
}

func TestPaymentRepository_GetByProviderPaymentId_InvalidId(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	repo := repositories.NewPaymentRepository(client)

	_, err = repo.GetByProviderPaymentId(ctx, "invalid-payment-id")
	require.Error(t, err)
}
