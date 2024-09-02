package workers

import (
	commonConstants "common/constants"
	"context"
)

const EXAM_CATEGORY_TYPE = commonConstants.Banking

func (w *Worker) AddDescriptiveQuestionsInDatabase() error {
	ctx := context.Background()
	const EXAM_TYPE = commonConstants.Descriptive

	err := w.examService.AddCachedQuestionInDatabase(ctx, EXAM_TYPE)
	if err != nil {
		return err
	}

	return nil
}
