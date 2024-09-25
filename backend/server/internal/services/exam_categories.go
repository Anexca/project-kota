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

	for _, exam := range category.Edges.Exams {
		categoryExamType := models.CategoryExamType{
			Id:           exam.ID,
			ExamName:     exam.Name,
			CategoryId:   category.ID,
			IsActive:     exam.IsActive,
			Description:  exam.Description,
			TypeOfExam:   exam.Type.String(),
			CategoryName: category.Name.String(),
			LogoUrl:      exam.LogoURL,
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

	for _, exam := range category.Edges.Exams {
		categoryExamType := models.CategoryExamType{
			Id:           exam.ID,
			ExamName:     exam.Name,
			CategoryId:   category.ID,
			IsActive:     exam.IsActive,
			Description:  exam.Description,
			TypeOfExam:   exam.Type.String(),
			CategoryName: category.Name.String(),
			LogoUrl:      exam.LogoURL,
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
