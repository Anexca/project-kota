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

	var descriptiveQuestions []models.DescriptiveQuestion

	_, err := w.examService.GetCachedQuestions(ctx, EXAM_TYPE, &descriptiveQuestions)
	if err != nil {
		return err
	}

	return nil
}
