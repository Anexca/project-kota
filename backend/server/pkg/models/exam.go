package models

import "time"

type ExamModelType int

const (
	DescriptiveExamType       ExamModelType = iota
	GeneratedExamOverviewType ExamModelType = iota
)

type DescriptiveExam struct {
	Type  string   `json:"type"`
	Topic string   `json:"topic"`
	Hints []string `json:"hints"`
}

type GeneratedExamOverview struct {
	Id                int                    `json:"id"`
	RawExamData       map[string]interface{} `json:"raw_exam_data"`
	UserAttempts      int                    `json:"user_attempts"`
	MaxAttempts       int                    `json:"max_attempts"`
	DurationSeconds   int                    `json:"duration_seconds"`
	NumberOfQuestions int                    `json:"number_of_questions"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
}

type DescriptiveExamAssessmentResult struct {
	Rating           string   `json:"rating,omitempty"`
	Strengths        []string `json:"strengths,omitempty"`
	Weakness         []string `json:"weakness,omitempty"`
	CorrectedVersion string   `json:"corrected_version,omitempty"`
}

type AssessmentDetails struct {
	Id               int                             `json:"id"`
	CompletedSeconds int                             `json:"completed_seconds"`
	RawAssesmentData DescriptiveExamAssessmentResult `json:"raw_assesment_data,omitempty"`
	Status           string                          `json:"status"`
}
