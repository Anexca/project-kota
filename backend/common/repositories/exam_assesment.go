package repositories

import (
	"common/constants"
	"common/ent"
	"common/ent/examassesment"
	"common/ent/examattempt"
	"common/ent/user"
	"context"

	"github.com/google/uuid"
)

type ExamAssesmentRepository struct {
	dbClient *ent.Client
}

type AssesmentModel struct {
	CompletedSeconds  int
	Status            constants.AssessmentStatus
	RawAssessmentData map[string]interface{}
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
		SetStatus(examassesment.Status(model.Status))

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

	_, err := query.Save(ctx)
	return err
}

func (e *ExamAssesmentRepository) GetById(ctx context.Context, assesmentId int, userId string) (*ent.ExamAssesment, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return e.dbClient.ExamAssesment.Query().
		Where(examassesment.HasAttemptWith(examattempt.HasUserWith(user.ID(userUid))), examassesment.ID(assesmentId)).
		Only(ctx)
}
