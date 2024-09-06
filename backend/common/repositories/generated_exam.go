package repositories

import (
	"common/ent"
	"common/ent/exam"
	"common/ent/generatedexam"
	"context"
	"fmt"
)

type GeneratedExamRepository struct {
	dbClient *ent.Client
}

func NewGeneratedExamRepository(dbClient *ent.Client) *GeneratedExamRepository {
	return &GeneratedExamRepository{
		dbClient: dbClient,
	}
}

func (q *GeneratedExamRepository) AddMany(ctx context.Context, exams []any, ex *ent.Exam) ([]*ent.GeneratedExam, error) {
	bulk := make([]*ent.GeneratedExamCreate, len(exams))

	for i, exam := range exams {
		jsonData, ok := exam.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid exam data type at index %d", i)
		}

		bulk[i] = q.dbClient.GeneratedExam.Create().
			SetRawExamData(jsonData).
			SetExam(ex)
	}

	return q.dbClient.GeneratedExam.CreateBulk(bulk...).Save(ctx)
}

func (q *GeneratedExamRepository) GetById(ctx context.Context, generatedExamId int) (*ent.GeneratedExam, error) {
	return q.dbClient.GeneratedExam.Query().
		Where(generatedexam.ID(generatedExamId)).
		WithExam().
		Only(ctx)
}

func (q *GeneratedExamRepository) GetByExam(ctx context.Context, ex *ent.Exam) ([]*ent.GeneratedExam, error) {
	return q.dbClient.GeneratedExam.Query().
		Where(generatedexam.HasExamWith(exam.ID(ex.ID)), generatedexam.IsActive(true)).
		WithAttempts().
		WithExam().
		All(ctx)
}
