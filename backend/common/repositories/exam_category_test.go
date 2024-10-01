package repositories_test

import (
	"context"
	"testing"

	"common/constants"
	"common/ent/examcategory"
	"common/repositories"

	"github.com/stretchr/testify/require"
)

func TestExamCategoryRepository_Get(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	_, err = client.ExamCategory.Create().
		SetName(examcategory.NameBANKING).
		SetDescription("Banking exams").
		Save(ctx)
	require.NoError(t, err)

	_, err = client.ExamCategory.Create().
		SetName(examcategory.NameBANKING).
		SetDescription("Engineering exams").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamCategoryRepository(client)

	categories, err := repo.Get(ctx)
	require.NoError(t, err)
	require.Len(t, categories, 2)
}

func TestExamCategoryRepository_GetByName(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	category, err := client.ExamCategory.Create().
		SetName(examcategory.NameBANKING).
		SetDescription("Banking exams").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamCategoryRepository(client)

	foundCategory, err := repo.GetByName(ctx, constants.ExamCategoryNameBanking)
	require.NoError(t, err)
	require.Equal(t, category.ID, foundCategory.ID)
	require.Equal(t, "BANKING", string(foundCategory.Name))

	_, err = repo.GetByName(ctx, "Non-Existent Category")
	require.Error(t, err)
}

func TestExamCategoryRepository_GetByName_WithExams(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	category, err := client.ExamCategory.Create().
		SetName(examcategory.Name(constants.ExamCategoryNameBanking)).
		SetDescription("Banking exams").
		Save(ctx)
	require.NoError(t, err)

	_, err = client.Exam.Create().
		SetName("Banking Exam 1").
		SetDescription("First banking exam").
		SetCategory(category).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamCategoryRepository(client)

	foundCategory, err := repo.GetByName(ctx, constants.ExamCategoryNameBanking)
	require.NoError(t, err)
	require.Len(t, foundCategory.Edges.Exams, 1)
	require.Equal(t, "Banking Exam 1", foundCategory.Edges.Exams[0].Name)
}

func TestExamCategoryRepository_GetByName_Inactive(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	category, err := client.ExamCategory.Create().
		SetName(examcategory.NameBANKING).
		SetDescription("An inactive category").
		Save(ctx)
	require.NoError(t, err)

	_, err = client.ExamCategory.UpdateOneID(category.ID).
		SetIsActive(false).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamCategoryRepository(client)

	foundCategory, err := repo.GetByName(ctx, constants.ExamCategoryNameBanking)
	require.NoError(t, err)
	require.Equal(t, category.ID, foundCategory.ID)
	require.False(t, foundCategory.IsActive)
}

func TestExamCategoryRepository_GetByName_InvalidName(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	repo := repositories.NewExamCategoryRepository(client)

	_, err = repo.GetByName(ctx, "")
	require.Error(t, err)
}

func TestExamCategoryRepository_GetByName_CategoryNotActive(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	category, err := client.ExamCategory.Create().
		SetName(examcategory.NameBANKING).
		SetDescription("Inactive category").
		SetIsActive(false).
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamCategoryRepository(client)

	foundCategory, err := repo.GetByName(ctx, constants.ExamCategoryNameBanking)
	require.NoError(t, err)
	require.Equal(t, category.ID, foundCategory.ID)
	require.False(t, foundCategory.IsActive)
}

func TestExamCategoryRepository_Get_NoCategories(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	repo := repositories.NewExamCategoryRepository(client)

	categories, err := repo.Get(ctx)
	require.NoError(t, err)
	require.Empty(t, categories)
}

func TestExamCategoryRepository_GetByName_EmptyDatabase(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	repo := repositories.NewExamCategoryRepository(client)

	_, err = repo.GetByName(ctx, "Non-Existent Category")
	require.Error(t, err)
}
