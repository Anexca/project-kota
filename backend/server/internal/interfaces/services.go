package interfaces

import (
	"common/constants"
	"common/ent"
	"context"
	"server/pkg/models"
)

// PromptServiceInterface defines the contract for PromptService
type PromptServiceInterface interface {
	GetPromptResult(ctx context.Context, prompt string, model constants.GenAiModel) (string, error)
}

// ExamGenerationServiceInterface defines the contract for ExamGenerationService
type ExamGenerationServiceInterface interface {
	GetGeneratedExamById(ctx context.Context, generatedExamId int, userId string, isOpen bool) (*models.GeneratedExamOverview, error)
}

// AccessServiceInterface defines the interface for the AccessService.
type AccessServiceInterface interface {
	UserHasAccessToExam(ctx context.Context, examId int, userId string) (bool, error)
	GetAccessibleExamsForUser(ctx context.Context, exams []*ent.Exam, userId string) ([]*ent.Exam, error)
}

// ExamAssesmentServiceInterface defines the methods available for the ExamAssesmentService
type ExamAssesmentServiceInterface interface {
	StartNewDescriptiveAssesment(ctx context.Context, generatedExamId int, attempt *ent.ExamAttempt, request *models.DescriptiveExamAssesmentRequest, userId string, isOpen bool) (*models.AssessmentDetails, error)
	GetAssesmentById(ctx context.Context, assessmentId int, userId string) (*models.AssessmentDetails, error)
	GetExamAssessments(ctx context.Context, generatedExamId int, userId string) ([]models.AssessmentDetails, error)
	AssessDescriptiveExam(ctx context.Context, generatedExamId int, assessmentId int, content string, userId string, isOpen bool)
}
