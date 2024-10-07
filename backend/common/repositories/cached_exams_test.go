package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"common/repositories"
)

func TestCachedExamRepository_Create(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	examEntity, err := client.Exam.Create().
		SetName("Sample Exam").
		SetDescription("This is a sample exam.").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewCachedExamRepository(client)

	cacheUID := uuid.New().String()
	expiry := time.Hour * 24

	cachedExam, err := repo.Create(ctx, cacheUID, expiry, examEntity)
	require.NoError(t, err)
	require.Equal(t, cacheUID, cachedExam.CacheUID)
	require.Equal(t, false, cachedExam.IsUsed)
	require.WithinDuration(t, time.Now().Add(expiry), cachedExam.ExpiresAt, time.Second)
}

func TestCachedExamRepository_GetByExam(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	examEntity, err := client.Exam.Create().
		SetName("Sample Exam").
		SetDescription("This is a sample exam.").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewCachedExamRepository(client)

	cacheUID := uuid.New().String()
	expiry := time.Hour * 24

	_, err = repo.Create(ctx, cacheUID, expiry, examEntity)
	require.NoError(t, err)

	cachedExams, err := repo.GetByExam(ctx, examEntity)
	require.NoError(t, err)
	require.Len(t, cachedExams, 1)
	require.Equal(t, cacheUID, cachedExams[0].CacheUID)
}

func TestCachedExamRepository_GetByExam_NoCachedExams(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	examEntity, err := client.Exam.Create().
		SetName("Sample Exam").
		SetDescription("This is a sample exam.").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewCachedExamRepository(client)

	cachedExams, err := repo.GetByExam(ctx, examEntity)
	require.NoError(t, err)
	require.Empty(t, cachedExams)
}

func TestCachedExamRepository_MarkAsUsed(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	examEntity, err := client.Exam.Create().
		SetName("Sample Exam").
		SetDescription("This is a sample exam.").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewCachedExamRepository(client)

	cacheUID := uuid.New().String()
	expiry := time.Hour * 24

	cachedExam, err := repo.Create(ctx, cacheUID, expiry, examEntity)
	require.NoError(t, err)

	err = repo.MarkAsUsed(ctx, cachedExam.ID)
	require.NoError(t, err)

	updatedCachedExam, err := client.CachedExam.Get(ctx, cachedExam.ID)
	require.NoError(t, err)
	require.True(t, updatedCachedExam.IsUsed)
}

func TestCachedExamRepository_MarkAsUsed_InvalidId(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	repo := repositories.NewCachedExamRepository(client)

	err = repo.MarkAsUsed(ctx, 999) // Assuming this ID doesn't exist
	require.Error(t, err)
}
