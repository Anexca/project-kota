package models

import "time"

type ExamModelType int

const (
	DescriptiveExamType       ExamModelType = iota
	GeneratedExamOverviewType ExamModelType = iota
)

type DescriptiveExam struct {
	Type                    string   `json:"type" validate:"required,oneof=formal_letter essay"`
	Topic                   string   `json:"topic" validate:"required"`
	Hints                   []string `json:"hints" validate:"required"`
	MaxNumberOfWordsAllowed string   `json:"max_number_of_words" validate:"required"`
	TotalMarks              string   `json:"total_marks" validate:"required"`
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
	Rating           string   `json:"rating"`
	Strengths        []string `json:"strengths"`
	Weakness         []string `json:"weakness"`
	CorrectedVersion string   `json:"corrected_version"`
	ProfanityCheck   string   `json:"profanity_check,omitempty" `
}

type AssessmentDetails struct {
	Id                int                    `json:"id"`
	CompletedSeconds  int                    `json:"completed_seconds"`
	RawAssesmentData  map[string]interface{} `json:"raw_assesment_data,omitempty"`
	RawUserSubmission map[string]interface{} `json:"raw_user_submission,omitempty"`
	Status            string                 `json:"status"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
}

type UserExamAttempt struct {
	Id           int
	IsActive     bool
	ExamName     string
	ExamCategory string
	Topic        string
	Type         string
	Attempts     []Attempt
}

type Attempt struct {
	AttemptId     int
	AttemptNumber int
	AssessmentId  int
	AttemptDate   time.Time
}
