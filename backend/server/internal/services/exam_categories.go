package services

import (
	"common/constants"
	"common/ent"
	"common/repositories"
	"context"
	"server/pkg/models"
)

type ExamCategoryService struct {
	examRepository         *repositories.ExamRepository
	examCategoryRepository *repositories.ExamCategoryRepository
}

func NewExamCategoryService(dbClient *ent.Client) *ExamCategoryService {
	examRepository := repositories.NewExamRespository(dbClient)
	examCategoryRepository := repositories.NewExamCategoryRepository(dbClient)

	return &ExamCategoryService{
		examRepository:         examRepository,
		examCategoryRepository: examCategoryRepository,
	}
}

func (e *ExamCategoryService) GetBankingDescriptiveExamsTypes(ctx context.Context) ([]models.CategoryExamType, error) {
	category, err := e.examCategoryRepository.GetByName(ctx, constants.ExamCategoryNameBanking, constants.ExamTypeDescriptive)
	if err != nil {
		return nil, err
	}

	var categoryExamTypes []models.CategoryExamType

	for _, examGroup := range category.Edges.Groups {
		categoryExamType := models.CategoryExamType{
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

func (e *ExamCategoryService) GetBankingMCQExamTypes(ctx context.Context) ([]models.CategoryExamType, error) {
	category, err := e.examCategoryRepository.GetByName(ctx, constants.ExamCategoryNameBanking, constants.ExamTypeMCQ)
	if err != nil {
		return nil, err
	}

	var categoryExamTypes []models.CategoryExamType

	for _, examGroup := range category.Edges.Groups {
		categoryExamType := models.CategoryExamType{
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

func (e *ExamCategoryService) GetCategoryExamById(ctx context.Context, examId int) (*models.CategoryExamType, error) {
	exam, err := e.examRepository.GetById(ctx, examId)
	if err != nil {
		return nil, err
	}

	categoryExamType := models.CategoryExamType{
		Id:          exam.ID,
		ExamName:    exam.Name,
		IsActive:    exam.IsActive,
		Description: exam.Description,
		TypeOfExam:  exam.Type.String(),
		LogoUrl:     exam.LogoURL,
	}

	return &categoryExamType, nil
}
