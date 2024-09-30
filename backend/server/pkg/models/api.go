package models

type DescriptiveExamAssesmentRequest struct {
	CompletedSeconds int    `json:"completed_seconds" validate:"required"`
	Content          string `json:"content" validate:"required"`
}
