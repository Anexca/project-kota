package repositories_test

import (
	"context"
	"testing"

	"common/repositories"

	"github.com/stretchr/testify/require"
)

func TestExamSettingRepository_GetByExam(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	exam, err := client.Exam.Create().
		SetName("Test Exam").
		SetDescription("Test Description").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamSettingRepository(client)

	setting, err := client.ExamSetting.Create().
		SetExamID(exam.ID).
		SetNumberOfQuestions(10).
		SetDurationSeconds(3600).
		Save(ctx)
	require.NoError(t, err)

	retrievedSetting, err := repo.GetByExam(ctx, exam.ID)
	require.NoError(t, err)
	require.Equal(t, setting.ID, retrievedSetting.ID)
	require.Equal(t, setting.NumberOfQuestions, retrievedSetting.NumberOfQuestions)
	require.Equal(t, setting.DurationSeconds, retrievedSetting.DurationSeconds)

	// Test retrieving an exam setting with optional fields
	settingWithOptional := setting.Update().
		SetNegativeMarking(0.5).
		SetAiPrompt("Sample AI Prompt").
		SaveX(ctx)

	retrievedSettingOptional, err := repo.GetByExam(ctx, exam.ID)
	require.NoError(t, err)
	require.Equal(t, settingWithOptional.NegativeMarking, retrievedSettingOptional.NegativeMarking)
	require.Equal(t, settingWithOptional.AiPrompt, retrievedSettingOptional.AiPrompt)
}

func TestExamSettingRepository_GetByExam_NonExistentExam(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	repo := repositories.NewExamSettingRepository(client)

	// Using a non-existent exam ID
	_, err = repo.GetByExam(ctx, 999)
	require.Error(t, err)
}

func TestExamSettingRepository_GetByExam_ExamWithoutSetting(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	exam, err := client.Exam.Create().
		SetName("Test Exam Without Setting").
		SetDescription("Test Description").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamSettingRepository(client)

	// Attempt to get an exam setting for an exam that has none
	retrievedSetting, err := repo.GetByExam(ctx, exam.ID)
	require.Error(t, err)
	require.Nil(t, retrievedSetting)
}
