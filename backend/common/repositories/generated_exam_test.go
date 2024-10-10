package repositories_test

import (
	"context"
	"testing"

	"common/repositories"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestGeneratedExamRepository_AddMany(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	// Create an exam
	exam, err := client.Exam.Create().
		SetName("Test Exam").
		SetDescription("Test Description").
		Save(ctx)
	require.NoError(t, err)

	examData := []any{
		map[string]interface{}{"question": "What is Go?"},
		map[string]interface{}{"question": "Explain pointers in C."},
	}

	repo := repositories.NewGeneratedExamRepository(client)
	generatedExams, err := repo.AddMany(ctx, examData, exam)
	require.NoError(t, err)
	require.Len(t, generatedExams, 2)
	require.Equal(t, "What is Go?", generatedExams[0].RawExamData["question"])
	require.Equal(t, "Explain pointers in C.", generatedExams[1].RawExamData["question"])
}

func TestGeneratedExamRepository_Add(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	exam, err := client.Exam.Create().
		SetName("Test Exam").
		SetDescription("Test Description").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewGeneratedExamRepository(client)
	examData := map[string]interface{}{"question": "What is Go?"}
	generatedExam, err := repo.Add(ctx, examData, exam.ID)
	require.NoError(t, err)
	require.Equal(t, "What is Go?", generatedExam.RawExamData["question"])
}

func TestGeneratedExamRepository_UpdateMany(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	exam, err := client.Exam.Create().
		SetName("Test Exam").
		SetDescription("Test Description").
		Save(ctx)
	require.NoError(t, err)

	examData := []any{
		map[string]interface{}{"question": "What is Go?"},
		map[string]interface{}{"question": "Explain pointers in C."},
	}

	repo := repositories.NewGeneratedExamRepository(client)
	generatedExams, err := repo.AddMany(ctx, examData, exam)
	require.NoError(t, err)

	generatedExams[0].IsActive = false
	err = repo.UpdateMany(ctx, generatedExams)
	require.NoError(t, err)

	updatedExam, err := repo.GetById(ctx, generatedExams[0].ID)
	require.NoError(t, err)
	require.False(t, updatedExam.IsActive)
}

func TestGeneratedExamRepository_GetById(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	exam, err := client.Exam.Create().
		SetName("Test Exam").
		SetDescription("Test Description").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewGeneratedExamRepository(client)
	examData := map[string]interface{}{"question": "What is Go?"}
	generatedExam, err := repo.Add(ctx, examData, exam.ID)
	require.NoError(t, err)

	retrievedExam, err := repo.GetById(ctx, generatedExam.ID)
	require.NoError(t, err)
	require.Equal(t, "What is Go?", retrievedExam.RawExamData["question"])
}

func TestGeneratedExamRepository_GetOpenById(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	// Create an exam
	exam, err := client.Exam.Create().
		SetName("Test Exam").
		SetDescription("Test Description").
		Save(ctx) // We are not setting isActive because we don't care about it
	require.NoError(t, err)

	// Create a generated exam
	repo := repositories.NewGeneratedExamRepository(client)
	examData := map[string]interface{}{"question": "What is Go?"}
	generatedExam, err := repo.Add(ctx, examData, exam.ID)
	require.NoError(t, err)

	// Set the generated exam as open (no need to worry about isActive)
	_, err = client.GeneratedExam.UpdateOneID(generatedExam.ID).
		SetIsOpen(true).
		SetIsActive(false).
		Save(ctx)
	require.NoError(t, err)

	// Fetch the exam based on open status
	openExam, err := repo.GetOpenById(ctx, generatedExam.ID, true)
	require.NoError(t, err)
	require.NotNil(t, openExam)
	require.True(t, openExam.IsOpen) // Check that the exam is indeed open
}

func TestGeneratedExamRepository_GetActiveById(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	exam, err := client.Exam.Create().
		SetName("Test Exam").
		SetDescription("Test Description").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewGeneratedExamRepository(client)
	examData := map[string]interface{}{"question": "What is Go?"}
	generatedExam, err := repo.Add(ctx, examData, exam.ID)
	require.NoError(t, err)

	retrievedExam, err := repo.GetActiveById(ctx, generatedExam.ID, true)
	require.NoError(t, err)
	require.True(t, retrievedExam.IsActive)
}

func TestGeneratedExamRepository_GetPaginatedExamsByUserAndDate(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	// Create a user
	userID := uuid.New()
	_, err = client.User.Create().
		SetID(userID).
		SetEmail("test@example.com").
		Save(ctx)
	require.NoError(t, err)

	exam, err := client.Exam.Create().
		SetName("Test Exam").
		SetDescription("Test Description").
		Save(ctx)
	require.NoError(t, err)

	generatedExam, err := client.GeneratedExam.Create().
		SetRawExamData(map[string]interface{}{"question": "What is Go?"}).
		SetExam(exam).
		Save(ctx)
	require.NoError(t, err)

	_, err = client.ExamAttempt.Create().
		SetAttemptNumber(1).
		SetGeneratedexam(generatedExam).
		SetUserID(userID).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewGeneratedExamRepository(client)
	exams, err := repo.GetPaginatedExamsByUserAndDate(ctx, userID.String(), 1, 10, nil, nil, nil, nil)
	require.NoError(t, err)
	require.Len(t, exams, 1)
}

// func TestGeneratedExamRepository_GetByWeekOffset(t *testing.T) {
// 	client, err := setupTestDB(t)
// 	require.NoError(t, err)
// 	defer client.Close()

// 	ctx := context.Background()

// 	// Create an exam
// 	exam, err := client.Exam.Create().
// 		SetName("Test Exam").
// 		SetDescription("Test Description").
// 		Save(ctx)
// 	require.NoError(t, err)

// 	// Set a date for 7 days ago (last week)
// 	oneWeekAgo := time.Now().AddDate(0, 0, -7)

// 	// Create a generated exam with a past created_at date and inactive status
// 	repo := repositories.NewGeneratedExamRepository(client)
// 	examData := map[string]interface{}{"question": "What is Go?"}
// 	genExam, err := client.GeneratedExam.Create().
// 		SetRawExamData(examData).
// 		SetExam(exam).
// 		SetIsActive(false).
// 		SetCreatedAt(oneWeekAgo).
// 		SetUpdatedAt(oneWeekAgo).
// 		Save(ctx)
// 	require.NoError(t, err)

// 	// Print the actual CreatedAt for debugging
// 	fmt.Printf("Generated Exam CreatedAt: %v\n", genExam.CreatedAt)

// 	// Retrieve exams from the last week
// 	exams, err := repo.GetByWeekOffset(ctx, exam, 1, 10)
// 	require.NoError(t, err)

// 	// Ensure the exam is retrieved
// 	require.Len(t, exams, 1)
// 	require.WithinDuration(t, oneWeekAgo, exams[0].CreatedAt, time.Second)
// }
