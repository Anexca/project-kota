package repositories

import (
	"context"

	"github.com/google/uuid"

	"common/constants"
	"common/ent"
	"common/ent/examassesment"
	"common/ent/examattempt"
	"common/ent/generatedexam"
	"common/ent/user"
)

type ExamAssesmentRepository struct {
	dbClient *ent.Client
}

type AssesmentModel struct {
	CompletedSeconds  int
	Status            constants.AssessmentStatusType
	RawAssessmentData map[string]interface{}
	RawUserSubmission map[string]interface{}
	Remarks           string
}

func NewExamAssesmentRepository(dbClient *ent.Client) *ExamAssesmentRepository {
	return &ExamAssesmentRepository{
		dbClient: dbClient,
	}
}

func (e *ExamAssesmentRepository) Create(ctx context.Context, attemptId int, model AssesmentModel) (*ent.ExamAssesment, error) {
	query := e.dbClient.ExamAssesment.Create().
		SetAttemptID(attemptId).
		SetCompletedSeconds(model.CompletedSeconds).
		SetStatus(examassesment.Status(model.Status)).
		SetRawUserSubmission(model.RawUserSubmission)

	if model.RawAssessmentData != nil {
		query.SetRawAssesmentData(model.RawAssessmentData)
	}

	return query.Save(ctx)
}

func (e *ExamAssesmentRepository) Update(ctx context.Context, assessmentId int, model AssesmentModel) error {
	query := e.dbClient.ExamAssesment.Update().
		Where(examassesment.ID(assessmentId)).
		SetStatus(examassesment.Status(model.Status))

	if model.RawAssessmentData != nil {
		query.SetRawAssesmentData(model.RawAssessmentData)
	}

	if model.Remarks != "" {
		query.SetRemarks(model.Remarks)
	}

	_, err := query.Save(ctx)
	return err
}

func (e *ExamAssesmentRepository) GetById(ctx context.Context, assesmentId int, userId string) (*ent.ExamAssesment, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return e.dbClient.ExamAssesment.Query().
		Where(
			examassesment.HasAttemptWith(examattempt.HasUserWith(user.ID(userUid))),
			examassesment.ID(assesmentId),
		).
		WithAttempt().
		Only(ctx)
}

func (e *ExamAssesmentRepository) GetByExam(ctx context.Context, generatedExamId int, userId string) ([]*ent.ExamAssesment, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return e.dbClient.ExamAssesment.Query().
		Where(examassesment.HasAttemptWith(
			examattempt.HasUserWith(user.ID(userUid)),
			examattempt.HasGeneratedexamWith(generatedexam.ID(generatedExamId)),
		)).All(ctx)
}
