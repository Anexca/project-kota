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

	return nil
}
