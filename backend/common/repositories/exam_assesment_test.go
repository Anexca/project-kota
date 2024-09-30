package repositories_test

import (
	"context"
	"testing"

	"common/constants"
	"common/repositories"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestExamAssessmentRepository_Create(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	userID := uuid.New()
	_, err = client.User.Create().SetID(userID).SetEmail("test@example.com").Save(ctx)
	require.NoError(t, err)

	generatedExam, err := client.GeneratedExam.Create().SetIsActive(true).Save(ctx)
	require.NoError(t, err)

	attempt, err := client.ExamAttempt.Create().
		SetUserID(userID).
		SetGeneratedexamID(generatedExam.ID).
		SetAttemptNumber(1).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamAssessmentRepository(client)

	model := repositories.AssessmentModel{
		CompletedSeconds: 120,
		Status:           constants.ASSESSMENT_PENDING,
		RawAssessmentData: map[string]interface{}{
			"score": 85,
		},
		RawUserSubmission: map[string]interface{}{
			"answers": []string{"A", "C", "B"},
		},
		Remarks: "Good attempt",
	}

	assessment, err := repo.Create(ctx, attempt.ID, model)
	require.NoError(t, err)
	require.Equal(t, model.RawAssessmentData, assessment.RawAssesmentData)
	require.Equal(t, string(model.Status), string(assessment.Status))
}

func TestExamAssessmentRepository_Update(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	userID := uuid.New()
	_, err = client.User.Create().SetID(userID).SetEmail("test@example.com").Save(ctx)
	require.NoError(t, err)

	generatedExam, err := client.GeneratedExam.Create().SetIsActive(true).Save(ctx)
	require.NoError(t, err)

	attempt, err := client.ExamAttempt.Create().
		SetUserID(userID).
		SetGeneratedexamID(generatedExam.ID).
		SetAttemptNumber(1).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamAssessmentRepository(client)

	model := repositories.AssessmentModel{
		CompletedSeconds: 120,
		Status:           constants.ASSESSMENT_PENDING,
	}

	assessment, err := repo.Create(ctx, attempt.ID, model)
	require.NoError(t, err)

	updateModel := repositories.AssessmentModel{
		CompletedSeconds: 150,
		Status:           constants.ASSESSMENT_COMPLETED,
		Remarks:          "Well done",
	}

	err = repo.Update(ctx, assessment.ID, updateModel)
	require.NoError(t, err)

	updatedAssessment, err := repo.GetById(ctx, assessment.ID, userID.String())
	require.NoError(t, err)
	require.Equal(t, string(updateModel.Status), string(updatedAssessment.Status))
	require.Equal(t, string(updateModel.Remarks), string(updatedAssessment.Remarks))
}

func TestExamAssessmentRepository_GetById(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	userID := uuid.New()
	_, err = client.User.Create().SetID(userID).SetEmail("test@example.com").Save(ctx)
	require.NoError(t, err)

	generatedExam, err := client.GeneratedExam.Create().SetIsActive(true).Save(ctx)
	require.NoError(t, err)

	attempt, err := client.ExamAttempt.Create().
		SetUserID(userID).
		SetGeneratedexamID(generatedExam.ID).
		SetAttemptNumber(1).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamAssessmentRepository(client)

	model := repositories.AssessmentModel{
		CompletedSeconds: 120,
		Status:           constants.ASSESSMENT_PENDING,
	}

	assessment, err := repo.Create(ctx, attempt.ID, model)
	require.NoError(t, err)

	retrievedAssessment, err := repo.GetById(ctx, assessment.ID, userID.String())
	require.NoError(t, err)
	require.Equal(t, assessment.ID, retrievedAssessment.ID)
}

func TestExamAssessmentRepository_GetByExam(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	userID := uuid.New()
	_, err = client.User.Create().SetID(userID).SetEmail("test@example.com").Save(ctx)
	require.NoError(t, err)

	generatedExam, err := client.GeneratedExam.Create().SetIsActive(true).Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamAssessmentRepository(client)

	attempt, err := client.ExamAttempt.Create().
		SetUserID(userID).
		SetGeneratedexamID(generatedExam.ID).
		SetAttemptNumber(1).
		Save(ctx)
	require.NoError(t, err)

	model := repositories.AssessmentModel{
		CompletedSeconds: 120,
		Status:           constants.ASSESSMENT_PENDING,
	}

	_, err = repo.Create(ctx, attempt.ID, model)
	require.NoError(t, err)

	assessments, err := repo.GetByExam(ctx, generatedExam.ID, userID.String())
	require.NoError(t, err)
	require.Len(t, assessments, 1)
}

func TestExamAssessmentRepository_GetById_InvalidId(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	repo := repositories.NewExamAssessmentRepository(client)

	_, err = repo.GetById(ctx, 99999, uuid.New().String())
	require.Error(t, err)
}

func TestExamAssessmentRepository_Create_InvalidUser(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	generatedExam, err := client.GeneratedExam.Create().SetIsActive(true).Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamAssessmentRepository(client)

	model := repositories.AssessmentModel{
		CompletedSeconds: 120,
		Status:           constants.ASSESSMENT_PENDING,
	}

	_, err = repo.Create(ctx, generatedExam.ID, model)
	require.Error(t, err)
}
