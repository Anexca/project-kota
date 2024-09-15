package services

import (
	"common/ent"
	"common/repositories"
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

// func GetExamByCategory
