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

func (q *GeneratedExamRepository) UpdateMany(ctx context.Context, generatedExams []*ent.GeneratedExam) error {
	tx, err := q.dbClient.Tx(ctx)
	if err != nil {
		return err
	}

	for _, generatedExam := range generatedExams {
		_, err := tx.GeneratedExam.UpdateOneID(generatedExam.ID).
			SetIsActive(generatedExam.IsActive).
			SetRawExamData(generatedExam.RawExamData).
			Save(ctx)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return rbErr
			}
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
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
		Order(ent.Desc(generatedexam.FieldUpdatedAt)).
		All(ctx)
}

func (q *GeneratedExamRepository) GetOpenGeneratedExamsByExam(ctx context.Context, ex *ent.Exam, monthOffset int) ([]*ent.GeneratedExam, error) {
	now := time.Now()

	targetMonth := now.AddDate(0, -monthOffset, 0)
	firstOfMonth := time.Date(targetMonth.Year(), targetMonth.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	return q.dbClient.GeneratedExam.Query().
		Where(
			generatedexam.HasExamWith(exam.ID(ex.ID)),
			generatedexam.IsActive(false),
			generatedexam.CreatedAtGTE(firstOfMonth),
			generatedexam.CreatedAtLTE(lastOfMonth),
		).
		WithAttempts().
		WithExam().
		Order(ent.Asc(generatedexam.FieldCreatedAt)).
		Limit(5).
		All(ctx)
}

func (q *GeneratedExamRepository) GetPaginatedExamsByUserAndDate(ctx context.Context, userId string, page, limit int, from, to *time.Time, examTypeId, categoryID *int) ([]*ent.GeneratedExam, error) {
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
			func(attemptQuery *ent.ExamAttemptQuery) {
				attemptQuery.WithAssesment().
					Order(ent.Desc(examattempt.FieldUpdatedAt))

				if from != nil && to != nil {
					attemptQuery.Where(examattempt.UpdatedAtGTE(*from), examattempt.UpdatedAtLTE(*to))
				} else if from != nil {
					attemptQuery.Where(examattempt.UpdatedAtGTE(*from))
				} else if to != nil {
					attemptQuery.Where(examattempt.UpdatedAtLTE(*to))
				}
			},
		).
		Order(ent.Desc(generatedexam.FieldUpdatedAt)).
		Limit(limit).
		Offset(offset)

	if examTypeId != nil {
		query = query.Where(generatedexam.HasExamWith(exam.IDEQ(*examTypeId)))
	}

	if categoryID != nil {
		query = query.Where(generatedexam.HasExamWith(exam.HasCategoryWith(examcategory.IDEQ(*categoryID))))
	}

	return query.All(ctx)
}
