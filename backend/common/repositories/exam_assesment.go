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

// ExamAssessmentRepository is a concrete implementation of ExamAssessmentRepositoryInterface.
type ExamAssessmentRepository struct {
	dbClient *ent.Client
}

// AssessmentModel is used to pass data for creating or updating assessments.
type AssessmentModel struct {
	CompletedSeconds  int
	Status            constants.AssessmentStatusType
	RawAssessmentData map[string]interface{}
	RawUserSubmission map[string]interface{}
	Remarks           string
}

// NewExamAssessmentRepository creates a new instance of ExamAssessmentRepository.
func NewExamAssessmentRepository(dbClient *ent.Client) *ExamAssessmentRepository {
	return &ExamAssessmentRepository{
		dbClient: dbClient,
	}
}

// Create creates a new exam assessment for a given attempt.
func (e *ExamAssessmentRepository) Create(ctx context.Context, attemptId int, model AssessmentModel) (*ent.ExamAssesment, error) {
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

// Update updates an existing exam assessment by ID.
func (e *ExamAssessmentRepository) Update(ctx context.Context, assessmentId int, model AssessmentModel) error {
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

// GetById retrieves a specific exam assessment by its ID and the user's UUID.
func (e *ExamAssessmentRepository) GetById(ctx context.Context, assessmentId int, userId string) (*ent.ExamAssesment, error) {
	userUid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return e.dbClient.ExamAssesment.Query().
		Where(
			examassesment.HasAttemptWith(examattempt.HasUserWith(user.ID(userUid))),
			examassesment.ID(assessmentId),
		).
		WithAttempt().
		Only(ctx)
}

// GetByExam retrieves all exam assessments for a specific generated exam and user.
func (e *ExamAssessmentRepository) GetByExam(ctx context.Context, generatedExamId int, userId string) ([]*ent.ExamAssesment, error) {
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
