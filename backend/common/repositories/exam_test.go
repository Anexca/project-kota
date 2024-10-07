package repositories_test

import (
	"context"
	"testing"

	"common/constants"
	"common/ent"
	"common/ent/exam"
	"common/ent/examcategory"
	"common/repositories"

	"github.com/stretchr/testify/require"
)

func TestExamRepository_GetById(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	// Create an exam
	exam, err := client.Exam.Create().
		SetName("Sample Exam").
		SetDescription("Sample description").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamRepository(client)

	retrievedExam, err := repo.GetById(ctx, exam.ID)
	require.NoError(t, err)
	require.NotNil(t, retrievedExam)
	require.Equal(t, "Sample Exam", retrievedExam.Name)
}

func TestExamRepository_GetActiveByExamsGroupId(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	// Create a group
	group, err := client.ExamGroup.Create().
		SetName("Sample Group").
		SetDescription("Sample group description").
		SetIsActive(true).
		Save(ctx)
	require.NoError(t, err)

	// Create multiple exams, some active, some inactive
	activeExam, err := client.Exam.Create().
		SetName("Active Exam").
		SetIsActive(true).
		SetGroup(group).
		SetDescription("TEST").
		Save(ctx)
	require.NoError(t, err)

	_, err = client.Exam.Create().
		SetName("Inactive Exam").
		SetIsActive(false).
		SetDescription("TEST").
		SetGroup(group).
		Save(ctx)
	require.NoError(t, err)

	// Create a generated exam for the active exam
	generatedExamData := map[string]interface{}{"question": "What is Go?"}
	_, err = client.GeneratedExam.Create().
		SetRawExamData(generatedExamData).
		SetExam(activeExam).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamRepository(client)

	activeExams, err := repo.GetActiveByExamsGroupId(ctx, group.ID, true)
	require.NoError(t, err)
	require.Len(t, activeExams, 1)
	require.Equal(t, activeExam.ID, activeExams[0].ID)
}

func TestExamRepository_GetActiveById(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	exam, err := client.Exam.Create().
		SetName("Active Exam").
		SetIsActive(true).
		SetDescription("TEST").
		Save(ctx)
	require.NoError(t, err)

	// Create a generated exam for the active exam
	generatedExamData := map[string]interface{}{"question": "What is Go?"}
	_, err = client.GeneratedExam.Create().
		SetRawExamData(generatedExamData).
		SetExam(exam).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamRepository(client)

	retrievedExam, err := repo.GetActiveById(ctx, exam.ID, true)
	require.NoError(t, err)
	require.NotNil(t, retrievedExam)
	require.Equal(t, "Active Exam", retrievedExam.Name)

	// Test for inactive exam retrieval
	inactiveExam, err := client.Exam.Create().
		SetName("Inactive Exam").
		SetIsActive(false).
		SetDescription("TEST").
		Save(ctx)
	require.NoError(t, err)

	// Create a generated exam for the active exam
	generatedExamData = map[string]interface{}{"question": "What is Go?"}
	_, err = client.GeneratedExam.Create().
		SetRawExamData(generatedExamData).
		SetExam(inactiveExam).
		Save(ctx)
	require.NoError(t, err)

	_, err = repo.GetActiveById(ctx, inactiveExam.ID, true)
	require.Error(t, err)
}

func TestExamRepository_GetByExamCategory(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	// Create categories
	category, err := client.ExamCategory.Create().
		SetName(examcategory.Name(constants.ExamCategoryNameBanking)).
		SetDescription("Banking exams").
		Save(ctx)
	require.NoError(t, err)

	// Create exams in the category
	_, err = client.Exam.Create().
		SetName("Banking Exam").
		SetCategory(category).
		SetDescription("TEST").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamRepository(client)

	exams, err := repo.GetByExamCategory(ctx, category)
	require.NoError(t, err)
	require.Len(t, exams, 1)
}

func TestExamRepository_GetByName(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	// Create an exam
	_, err = client.Exam.Create().
		SetName("Unique Exam").
		SetDescription("Test Exam with unique name").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamRepository(client)

	// Retrieve by name
	retrievedExam, err := repo.GetByName(ctx, "Unique Exam")
	require.NoError(t, err)
	require.NotNil(t, retrievedExam)
	require.Equal(t, "Unique Exam", retrievedExam.Name)

	// Test retrieval of a non-existing name
	_, err = repo.GetByName(ctx, "Nonexistent Exam")
	require.Error(t, err)
}

func TestExamRepository_GetActiveByType(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	// Create exams with different types
	mcqExam, err := client.Exam.Create().
		SetName("MCQ Exam").
		SetType(exam.Type(constants.ExamTypeMCQ)).
		SetIsActive(true).
		SetDescription("TEST").
		Save(ctx)
	require.NoError(t, err)

	descriptiveExam, err := client.Exam.Create().
		SetName("Descriptive Exam").
		SetType(exam.Type(constants.ExamTypeDescriptive)).
		SetIsActive(true).
		SetDescription("TEST").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamRepository(client)

	// Retrieve active MCQ exams
	activeMcqExams, err := repo.GetActiveByType(ctx, constants.ExamTypeMCQ)
	require.NoError(t, err)
	require.Len(t, activeMcqExams, 1)
	require.Equal(t, mcqExam.Name, activeMcqExams[0].Name)

	// Retrieve active Descriptive exams
	activeDescriptiveExams, err := repo.GetActiveByType(ctx, constants.ExamTypeDescriptive)
	require.NoError(t, err)
	require.Len(t, activeDescriptiveExams, 1)
	require.Equal(t, descriptiveExam.Name, activeDescriptiveExams[0].Name)
}

func TestExamRepository_GetById_MultipleRecords(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	// Create multiple exams
	exam1, err := client.Exam.Create().
		SetName("First Exam").
		SetDescription("TEST").
		SetDescription("First exam description").
		Save(ctx)
	require.NoError(t, err)

	exam2, err := client.Exam.Create().
		SetName("Second Exam").
		SetDescription("TEST").
		SetDescription("Second exam description").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamRepository(client)

	// Retrieve the first exam
	retrievedExam1, err := repo.GetById(ctx, exam1.ID)
	require.NoError(t, err)
	require.Equal(t, "First Exam", retrievedExam1.Name)

	// Retrieve the second exam
	retrievedExam2, err := repo.GetById(ctx, exam2.ID)
	require.NoError(t, err)
	require.Equal(t, "Second Exam", retrievedExam2.Name)
}

func TestExamRepository_GetById_FailureCases(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	repo := repositories.NewExamRepository(client)

	// Try to retrieve an exam with an invalid ID
	_, err = repo.GetById(ctx, -1)
	require.Error(t, err)
}

func TestExamRepository_GetByCategory_FailureCases(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	repo := repositories.NewExamRepository(client)

	// Try to retrieve exams for a non-existent category
	category := &ent.ExamCategory{
		ID: -1,
	}
	e, err := repo.GetByExamCategory(ctx, category)
	require.Nil(t, err)
	require.Len(t, e, 0)
}

func TestExamRepository_GetById_EmptyFields(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	// Create an exam with empty description and type
	exam, err := client.Exam.Create().
		SetName("Empty Fields Exam").
		SetDescription("asas").
		SetType(exam.Type(constants.ExamTypeDescriptive)).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamRepository(client)

	// Retrieve the exam and check for empty fields
	retrievedExam, err := repo.GetById(ctx, exam.ID)
	require.NoError(t, err)
	require.Equal(t, "Empty Fields Exam", retrievedExam.Name)
	require.Equal(t, "asas", retrievedExam.Description)
	require.Equal(t, "DESCRIPTIVE", string(retrievedExam.Type))
}
