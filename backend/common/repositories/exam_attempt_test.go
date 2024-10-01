package repositories_test

import (
	"context"
	"testing"

	"common/repositories"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestExamAttemptRepository_Create(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	userID := uuid.New()
	_, err = client.User.Create().SetID(userID).SetEmail("test@example.com").Save(ctx)
	require.NoError(t, err)

	generatedExam, err := client.GeneratedExam.Create().SetIsActive(true).SetIsOpen(false).Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamAttemptRepository(client)
	attempt, err := repo.Create(ctx, 0, generatedExam.ID, userID.String())
	require.NoError(t, err)
	require.NotNil(t, attempt)
	require.Equal(t, 1, attempt.AttemptNumber)
}

func TestExamAttemptRepository_GetById(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	userID := uuid.New()
	_, err = client.User.Create().SetID(userID).SetEmail("test@example.com").Save(ctx)
	require.NoError(t, err)

	generatedExam, err := client.GeneratedExam.Create().SetIsActive(true).SetIsOpen(false).Save(ctx)
	require.NoError(t, err)

	attempt, err := client.ExamAttempt.Create().SetUserID(userID).SetGeneratedexamID(generatedExam.ID).SetAttemptNumber(1).Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamAttemptRepository(client)
	retrievedAttempt, err := repo.GetById(ctx, attempt.ID, userID.String())
	require.NoError(t, err)
	require.NotNil(t, retrievedAttempt)
	require.Equal(t, attempt.ID, retrievedAttempt.ID)
}

func TestExamAttemptRepository_GetByUserId(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	userID := uuid.New()
	_, err = client.User.Create().SetID(userID).SetEmail("test@example.com").Save(ctx)
	require.NoError(t, err)

	generatedExam, err := client.GeneratedExam.Create().SetIsActive(true).SetIsOpen(false).Save(ctx)
	require.NoError(t, err)

	_, err = client.ExamAttempt.Create().SetUserID(userID).SetGeneratedexamID(generatedExam.ID).SetAttemptNumber(1).Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamAttemptRepository(client)
	attempts, err := repo.GetByUserId(ctx, userID.String())
	require.NoError(t, err)
	require.Len(t, attempts, 1)
}

func TestExamAttemptRepository_GetByExam(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	userID := uuid.New()
	_, err = client.User.Create().SetID(userID).SetEmail("test@example.com").Save(ctx)
	require.NoError(t, err)

	generatedExam, err := client.GeneratedExam.Create().SetIsActive(true).SetIsOpen(false).Save(ctx)
	require.NoError(t, err)

	_, err = client.ExamAttempt.Create().SetUserID(userID).SetGeneratedexamID(generatedExam.ID).SetAttemptNumber(1).Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamAttemptRepository(client)
	attempts, err := repo.GetByExam(ctx, generatedExam.ID, userID.String())
	require.NoError(t, err)
	require.Len(t, attempts, 1)
}

func TestExamAttemptRepository_GetById_InvalidUser(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	userID := uuid.New()
	_, err = client.User.Create().SetID(userID).SetEmail("test@example.com").Save(ctx)
	require.NoError(t, err)

	generatedExam, err := client.GeneratedExam.Create().SetIsActive(true).SetIsOpen(false).Save(ctx)
	require.NoError(t, err)

	attempt, err := client.ExamAttempt.Create().SetUserID(userID).SetGeneratedexamID(generatedExam.ID).SetAttemptNumber(1).Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamAttemptRepository(client)
	_, err = repo.GetById(ctx, attempt.ID, uuid.New().String())
	require.Error(t, err)
}

func TestExamAttemptRepository_Create_InvalidUser(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	generatedExam, err := client.GeneratedExam.Create().SetIsActive(true).SetIsOpen(false).Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamAttemptRepository(client)
	_, err = repo.Create(ctx, 0, generatedExam.ID, uuid.New().String())
	require.Error(t, err)
}
