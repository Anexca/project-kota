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
	category, err := e.examCategoryRepository.GetByName(ctx, constants.ExamCategoryNameBanking)
	if err != nil {
		return nil, err
	}

	var categoryExamTypes []models.CategoryExamType

	for _, exam := range category.Edges.Exams {
		categoryExamType := models.CategoryExamType{
			Id:           exam.ID,
			ExamName:     exam.Name,
			CategoryId:   category.ID,
			IsActive:     exam.IsActive,
			Description:  exam.Description,
			TypeOfExam:   exam.Type.String(),
			CategoryName: category.Name.String(),
		}

		categoryExamTypes = append(categoryExamTypes, categoryExamType)
	}

	return categoryExamTypes, nil
}
