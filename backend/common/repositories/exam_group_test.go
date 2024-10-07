package repositories_test

import (
	"context"
	"testing"

	"common/repositories"

	"github.com/stretchr/testify/require"
)

func TestExamGroupRepository_GetById(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	group, err := client.ExamGroup.Create().
		SetName("Test Group").
		SetDescription("Description of Test Group").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamGroupRepository(client)

	retrievedGroup, err := repo.GetById(ctx, group.ID)
	require.NoError(t, err)
	require.Equal(t, group.ID, retrievedGroup.ID)
	require.Equal(t, "Test Group", retrievedGroup.Name)
}

func TestExamGroupRepository_GetActiveByIdWithExams(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	group, err := client.ExamGroup.Create().
		SetName("Active Group").
		SetDescription("Active group description").
		SetIsActive(true).
		Save(ctx)
	require.NoError(t, err)

	exam1, err := client.Exam.Create().
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

	repo := repositories.NewExamGroupRepository(client)

	retrievedGroup, err := repo.GetActiveByIdWithExams(ctx, group.ID, true)
	require.NoError(t, err)
	require.Equal(t, group.ID, retrievedGroup.ID)
	require.Len(t, retrievedGroup.Edges.Exams, 1)
	require.Equal(t, exam1.ID, retrievedGroup.Edges.Exams[0].ID)
}

func TestExamGroupRepository_GetActiveByIdWithExams_NotFound(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	repo := repositories.NewExamGroupRepository(client)

	_, err = repo.GetActiveByIdWithExams(ctx, 9999, true)
	require.Error(t, err)
}

func TestExamGroupRepository_GetById_NotFound(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	repo := repositories.NewExamGroupRepository(client)

	_, err = repo.GetById(ctx, 9999)
	require.Error(t, err)
}

func TestExamGroupRepository_CreateAndRetrieve(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	repo := repositories.NewExamGroupRepository(client)

	group, err := client.ExamGroup.Create().
		SetName("New Group").
		SetDescription("New group description").
		Save(ctx)
	require.NoError(t, err)

	retrievedGroup, err := repo.GetById(ctx, group.ID)
	require.NoError(t, err)
	require.Equal(t, group.ID, retrievedGroup.ID)
}

func TestExamGroupRepository_Update(t *testing.T) {
	client, err := setupTestDB(t)
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	group, err := client.ExamGroup.Create().
		SetName("Update Group").
		SetDescription("Description before update").
		Save(ctx)
	require.NoError(t, err)

	repo := repositories.NewExamGroupRepository(client)

	group, err = client.ExamGroup.UpdateOne(group).
		SetDescription("Updated description").
		Save(ctx)
	require.NoError(t, err)

	retrievedGroup, err := repo.GetById(ctx, group.ID)
	require.NoError(t, err)
	require.Equal(t, "Updated description", retrievedGroup.Description)
}
