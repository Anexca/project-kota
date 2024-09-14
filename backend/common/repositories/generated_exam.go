package repositories

import (
	"common/ent"
	"common/ent/exam"
	"common/ent/examattempt"
	"common/ent/examcategory"
	"common/ent/generatedexam"
	"common/ent/user"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
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

func (q *GeneratedExamRepository) GetPaginatedExamsByUserAndDate(ctx context.Context, userId string, page, limit int, from, to *time.Time, examTypeId, categoryId *int) ([]*ent.GeneratedExam, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	query := q.dbClient.GeneratedExam.Query().
		Where(generatedexam.HasAttemptsWith(examattempt.HasUserWith(user.IDEQ(userUid)))).
		WithExam(
			func(query *ent.ExamQuery) {
				query.WithCategory()
			},
		).
		WithAttempts(
			func(query *ent.ExamAttemptQuery) {
				query.WithAssesment().
					Order(ent.Desc(examattempt.FieldUpdatedAt))
			},
		).
		Order(ent.Desc(generatedexam.FieldUpdatedAt)).
		Limit(limit).
		Offset(offset)

	if from != nil && to != nil {
		query = query.Where(generatedexam.UpdatedAtGTE(*from), generatedexam.UpdatedAtLTE(*to))
	} else if from != nil {
		query = query.Where(generatedexam.UpdatedAtGTE(*from))
	} else if to != nil {
		query = query.Where(generatedexam.UpdatedAtLTE(*to))
	}

	if examTypeId != nil {
		query = query.Where(generatedexam.HasExamWith(exam.IDEQ(*examTypeId)))
	}

	if categoryId != nil {
		query = query.Where(generatedexam.HasExamWith(exam.HasCategoryWith(examcategory.IDEQ(*categoryId))))
	}

	return query.All(ctx)
}
