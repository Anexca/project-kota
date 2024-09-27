package repositories

import (
	"common/ent"
	"common/ent/examgroup"
	"context"
)

type ExamGroupRepository struct {
	dbClient *ent.Client
}

func NewExamGroupRepository(dbClient *ent.Client) *ExamGroupRepository {
	return &ExamGroupRepository{
		dbClient: dbClient,
	}
}

func (e *ExamGroupRepository) GetById(ctx context.Context, examGroupId int) (*ent.ExamGroup, error) {
	return e.dbClient.ExamGroup.Query().
		Where(examgroup.IDEQ(examGroupId)).
		Only(ctx)
}
