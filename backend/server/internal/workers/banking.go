package workers

import (
	commonConstants "common/constants"
	"context"
	"log"
	"server/pkg/models"
)

func (w *Worker) AddDescriptiveQuestionsInDatabase() error {
	ctx := context.Background()

	exams, err := w.examService.GetActiveExams(ctx, commonConstants.ExamTypeDescriptive)
	if err != nil {
		return err
	}

	for _, exam := range exams {
		err := w.examService.GenerateExams(ctx, exam.ID, models.DescriptiveExamType)
		if err != nil {
			log.Printf("Failed to Add Descriptive Question in Database: %v", err)
		}

		err = w.examService.MarkExpiredExamsInactive(ctx, exam.ID)
		if err != nil {
			log.Printf("Failed to Add Descriptive Question in Database: %v", err)
		}
	}

	return nil
}

func (w *Worker) MarkDescriptiveQuestionsAsOpenInDatabase() error {
	ctx := context.Background()
	examId := 1 // only set for general questions, need to make this dynamic
	return w.examService.MarkQuestionsAsOpen(ctx, examId)
}
