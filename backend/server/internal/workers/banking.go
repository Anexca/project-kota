package workers

import (
	commonConstants "common/constants"
	"context"
	"log"
	"server/pkg/models"
)

const EXAM_CATEGORY_TYPE = commonConstants.ExamCategoryNameBanking

func (w *Worker) AddDescriptiveQuestionsInDatabase() error {
	ctx := context.Background()

	exams, err := w.examService.GetActiveExams(ctx, commonConstants.ExamTypeDescriptive)
	if err != nil {
		return err
	}

	for _, exam := range exams {
		err := w.examService.VGenerateExams(ctx, exam.ID, models.DescriptiveExamType)
		if err != nil {
			log.Printf("Failed to Add Descriptive Question in Database: %v", err)
		}

		err = w.examService.VMarkExpiredExamsInactive(ctx, exam.ID)
		if err != nil {
			log.Printf("Failed to Add Descriptive Question in Database: %v", err)
		}
	}

	return nil
}

func (w *Worker) MarkDescriptiveQuestionsAsOpenInDatabase() error {
	ctx := context.Background()
	const EXAM_TYPE = commonConstants.Descriptive

	return w.examService.MarkQuestionsAsOpen(ctx, EXAM_TYPE)
}
