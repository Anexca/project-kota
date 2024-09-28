package services

import (
	"context"

	"common/constants"
	"common/ent"
	"common/repositories"

	"server/pkg/models"
)

type ExamCategoryService struct {
	examRepository         *repositories.ExamRepository
	examGroupRepository    *repositories.ExamGroupRepository
	examCategoryRepository *repositories.ExamCategoryRepository
}

func NewExamCategoryService(dbClient *ent.Client) *ExamCategoryService {
	examRepository := repositories.NewExamRespository(dbClient)
	examGroupRepository := repositories.NewExamGroupRepository(dbClient)
	examCategoryRepository := repositories.NewExamCategoryRepository(dbClient)

	return &ExamCategoryService{
		examRepository:         examRepository,
		examGroupRepository:    examGroupRepository,
		examCategoryRepository: examCategoryRepository,
	}
}

func (e *ExamCategoryService) GetBankingExamGroups(ctx context.Context) ([]models.CategoryExamGroup, error) {
	category, err := e.examCategoryRepository.GetByName(ctx, constants.ExamCategoryNameBanking)
	if err != nil {
		return nil, err
	}

	var categoryExamTypes []models.CategoryExamGroup

	for _, examGroup := range category.Edges.Groups {
		categoryExamType := models.CategoryExamGroup{
			Id:           examGroup.ID,
			ExamName:     examGroup.Name,
			CategoryId:   category.ID,
			IsActive:     examGroup.IsActive,
			Description:  examGroup.Description,
			CategoryName: category.Name.String(),
			LogoUrl:      examGroup.LogoURL,
		}

		categoryExamTypes = append(categoryExamTypes, categoryExamType)
	}

	return categoryExamTypes, nil
}

func (e *ExamCategoryService) GetExamGroupById(ctx context.Context, examGroupId int) (*models.CategoryExamGroup, error) {
	exam, err := e.examGroupRepository.GetById(ctx, examGroupId)
	if err != nil {
		return nil, err
	}

	categoryExamType := models.CategoryExamGroup{
		Id:          exam.ID,
		ExamName:    exam.Name,
		IsActive:    exam.IsActive,
		Description: exam.Description,
		LogoUrl:     exam.LogoURL,
	}

	return &categoryExamType, nil
}
