package repositories

import (
	"common/ent"
	"common/ent/exam"
	"common/ent/examassesment"
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
			SetIsOpen(generatedExam.IsOpen).
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
		Where(generatedexam.IDEQ(generatedExamId)).
		WithExam().
		Only(ctx)
}

func (q *GeneratedExamRepository) GetOpenById(ctx context.Context, generatedExamId int, isOpen bool) (*ent.GeneratedExam, error) {
	return q.dbClient.GeneratedExam.Query().
		Where(generatedexam.IDEQ(generatedExamId), generatedexam.IsOpenEQ(isOpen), generatedexam.IsActive(!isOpen)).
		WithExam().
		Only(ctx)
}

func (q GeneratedExamRepository) GetActiveById(ctx context.Context, generatedExamId int, IsActive bool) (*ent.GeneratedExam, error) {
	return q.dbClient.GeneratedExam.Query().
		Where(generatedexam.IDEQ(generatedExamId), generatedexam.IsActiveEQ(IsActive)).
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

func (q *GeneratedExamRepository) GetByOpenFlag(ctx context.Context, examId int) ([]*ent.GeneratedExam, error) {
	return q.dbClient.GeneratedExam.Query().
		Where(generatedexam.IsOpenEQ(true), generatedexam.HasExamWith(exam.ID(examId))).
		WithExam().
		All(ctx)
}

func (q *GeneratedExamRepository) GetByMonthOffset(ctx context.Context, ex *ent.Exam, monthOffset, limit int) ([]*ent.GeneratedExam, error) {
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
		Order(ent.Desc(generatedexam.FieldCreatedAt)).
		Limit(limit).
		All(ctx)
}

func (q *GeneratedExamRepository) GetByWeekOffset(ctx context.Context, ex *ent.Exam, weekOffset, limit int) ([]*ent.GeneratedExam, error) {
	now := time.Now()

	targetWeekStart := now.AddDate(0, 0, -7*weekOffset)

	for targetWeekStart.Weekday() != time.Monday {
		targetWeekStart = targetWeekStart.AddDate(0, 0, -1)
	}

	targetWeekEnd := targetWeekStart.AddDate(0, 0, 6).Add(time.Hour*23 + time.Minute*59 + time.Second*59)

	fmt.Println("Start of week:", targetWeekStart)
	fmt.Println("End of week:", targetWeekEnd)

	return q.dbClient.GeneratedExam.Query().
		Where(
			generatedexam.HasExamWith(exam.IDEQ(ex.ID)),
			generatedexam.IsActive(false), // assuming you need inactive records
			generatedexam.CreatedAtGTE(targetWeekStart),
			generatedexam.CreatedAtLTE(targetWeekEnd),
		).
		WithAttempts().
		WithExam().
		Order(ent.Desc(generatedexam.FieldCreatedAt)).
		Limit(limit).
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
		Order(ent.Desc(examassesment.FieldUpdatedAt)).
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

func (q *GeneratedExamRepository) GetCountOfFilteredExamsDataByUserAndDate(ctx context.Context, userId string, from, to *time.Time, examTypeId, categoryID *int) (int, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return 0, err
	}

	query := q.dbClient.GeneratedExam.Query().
		Where(generatedexam.HasAttemptsWith(examattempt.HasUserWith(user.IDEQ(userUid))))

	if examTypeId != nil {
		query = query.Where(generatedexam.HasExamWith(exam.IDEQ(*examTypeId)))
	}

	if categoryID != nil {
		query = query.Where(generatedexam.HasExamWith(exam.HasCategoryWith(examcategory.IDEQ(*categoryID))))
	}

	if from != nil && to != nil {
		query = query.Where(generatedexam.HasAttemptsWith(examattempt.UpdatedAtGTE(*from), examattempt.UpdatedAtLTE(*to)))
	} else if from != nil {
		query = query.Where(generatedexam.HasAttemptsWith(examattempt.UpdatedAtGTE(*from)))
	} else if to != nil {
		query = query.Where(generatedexam.HasAttemptsWith(examattempt.UpdatedAtLTE(*to)))
	}

	// Get total count
	totalCount, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}

	return totalCount, nil
}
