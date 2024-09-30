package services

import (
	"context"

	"common/constants"
	"common/ent"
	commonInterfaces "common/interfaces"
	"common/repositories"

	"server/pkg/models"
)

type ExamCategoryService struct {
	examRepository         commonInterfaces.ExamRepositoryInterface
	examGroupRepository    commonInterfaces.ExamGroupRepositoryInterface
	examCategoryRepository commonInterfaces.ExamCategoryRepositoryInterface
}

func NewExamCategoryService(examRepo commonInterfaces.ExamRepositoryInterface, examGroupRepo commonInterfaces.ExamGroupRepositoryInterface, examCategoryRepo commonInterfaces.ExamCategoryRepositoryInterface) *ExamCategoryService {
	return &ExamCategoryService{
		examRepository:         examRepo,
		examGroupRepository:    examGroupRepo,
		examCategoryRepository: examCategoryRepo,
	}
}

func InitExamCategoryService(dbClient *ent.Client) *ExamCategoryService {
	examRepository := repositories.NewExamRepository(dbClient)
	examGroupRepository := repositories.NewExamGroupRepository(dbClient)
	examCategoryRepository := repositories.NewExamCategoryRepository(dbClient)

	return NewExamCategoryService(examRepository, examGroupRepository, examCategoryRepository)
}

func (e *ExamCategoryService) GetBankingExamGroups(ctx context.Context) ([]models.CategoryExamGroup, error) {
	category, err := e.examCategoryRepository.GetByName(ctx, constants.ExamCategoryNameBanking)
	if err != nil {
		return nil, err
	}

	var categoryExamGroups []models.CategoryExamGroup
	for _, examGroup := range category.Edges.Groups {
		examGroupModel := models.CategoryExamGroup{
			Id:           examGroup.ID,
			ExamName:     examGroup.Name,
			CategoryId:   category.ID,
			IsActive:     examGroup.IsActive,
			Description:  examGroup.Description,
			CategoryName: category.Name.String(),
			LogoUrl:      examGroup.LogoURL,
		}
		categoryExamGroups = append(categoryExamGroups, examGroupModel)
	}

	return categoryExamGroups, nil
}

func (e *ExamCategoryService) GetExamGroupById(ctx context.Context, examGroupId int) (*models.CategoryExamGroup, error) {
	examGroup, err := e.examGroupRepository.GetById(ctx, examGroupId)
	if err != nil {
		return nil, err
	}

	examGroupModel := models.CategoryExamGroup{
		Id:          examGroup.ID,
		ExamName:    examGroup.Name,
		IsActive:    examGroup.IsActive,
		Description: examGroup.Description,
		LogoUrl:     examGroup.LogoURL,
	}

	return &examGroupModel, nil
}
