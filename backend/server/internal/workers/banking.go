package workers

import (
	commonConstants "common/constants"
	"context"
	"server/pkg/models"
)

const EXAM_CATEGORY_TYPE = commonConstants.Banking

func (w *Worker) AddDescriptiveQuestionsInDatabase() error {
	ctx := context.Background()
	const EXAM_TYPE = commonConstants.Descriptive

	err := w.examService.GenerateExams(ctx, EXAM_TYPE, models.DescriptiveExamType)
	if err != nil {
		return err
	}

	err = w.examService.MarkExpiredExamsInactive(ctx, EXAM_TYPE)
	if err != nil {
		return err
	}

	return nil
}

func (w *Worker) MarkDescriptiveQuestionsAsOpenInDatabase() error {
	ctx := context.Background()
	const EXAM_TYPE = commonConstants.Descriptive

	return w.examService.MarkQuestionsAsOpen(ctx, EXAM_TYPE)
}
