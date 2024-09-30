package services_test

import (
	"common/constants"
	"common/ent"
	"common/mocks"
	"context"
	"server/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() (*services.ExamCategoryService, *mocks.MockExamRepository, *mocks.MockExamGroupRepository, *mocks.MockExamCategoryRepository) {
	mockExamRepo := new(mocks.MockExamRepository)
	mockExamGroupRepo := new(mocks.MockExamGroupRepository)
	mockExamCategoryRepo := new(mocks.MockExamCategoryRepository)
	service := services.NewExamCategoryService(mockExamRepo, mockExamGroupRepo, mockExamCategoryRepo)
	return service, mockExamRepo, mockExamGroupRepo, mockExamCategoryRepo
}

func TestExamCategoryService(t *testing.T) {
	service, _, mockExamGroupRepo, mockExamCategoryRepo := setup()
	ctx := context.Background()

	t.Run("GetBankingExamGroups", func(t *testing.T) {
		mockExamCategoryRepo.On("GetByName", ctx, constants.ExamCategoryNameBanking).Return(&ent.ExamCategory{
			Edges: ent.ExamCategoryEdges{
				Groups: []*ent.ExamGroup{
					{ID: 1, Name: "Group 1", IsActive: true, Description: "Description 1", LogoURL: "url1"},
					{ID: 2, Name: "Group 2", IsActive: true, Description: "Description 2", LogoURL: "url2"},
				},
			},
		}, nil)

		result, err := service.GetBankingExamGroups(ctx)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "Group 1", result[0].ExamName)
		mockExamCategoryRepo.AssertExpectations(t)
	})

	t.Run("GetExamGroupById", func(t *testing.T) {
		examGroupId := 1
		mockExamGroupRepo.On("GetById", ctx, examGroupId).Return(&ent.ExamGroup{
			ID:          1,
			Name:        "Group 1",
			IsActive:    true,
			Description: "Description 1",
			LogoURL:     "url1",
		}, nil)

		result, err := service.GetExamGroupById(ctx, examGroupId)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "Group 1", result.ExamName)
		mockExamGroupRepo.AssertExpectations(t)
	})
}
