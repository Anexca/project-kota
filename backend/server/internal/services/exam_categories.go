package services

import (
	"common/constants"
	"common/ent"
	"common/repositories"
	"context"
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

func (e *ExamCategoryService) GetBankingDescriptiveExamsTypes(ctx context.Context) ([]*ent.Exam, error) {
	examName := constants.EXAMS[constants.Descriptive]
	return e.examRepository.GetAllByName(ctx, examName)
}
