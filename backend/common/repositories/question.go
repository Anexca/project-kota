package repositories

import (
	"common/ent"
	"common/ent/exam"
	"common/ent/question"
	"context"
	"fmt"
)

type QuestionRepository struct {
	dbClient *ent.Client
}

func NewQuestionRepository(dbClient *ent.Client) *QuestionRepository {
	return &QuestionRepository{
		dbClient: dbClient,
	}
}

func (q *QuestionRepository) AddMany(ctx context.Context, questions []any, ex *ent.Exam) ([]*ent.Question, error) {
	bulk := make([]*ent.QuestionCreate, len(questions))

	for i, questionData := range questions {
		jsonData, ok := questionData.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid question data type at index %d", i)
		}

		bulk[i] = q.dbClient.Question.Create().
			SetRawQuestionData(jsonData).
			SetExam(ex)
	}

	return q.dbClient.Question.CreateBulk(bulk...).Save(ctx)
}

func (q *QuestionRepository) GetByExam(ctx context.Context, ex *ent.Exam) ([]*ent.Question, error) {
	return q.dbClient.Question.Query().Where(question.HasExamWith(exam.ID(ex.ID)), question.IsActive(true)).All(ctx)
}
