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
	ExamType          string                 `json:"exam_type"`
	ExamName          string                 `json:"exam_name"`
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
	AttemptedExamId int       `json:"attempted_exam_id"`
	IsActive        bool      `json:"is_active"`
	ExamType        string    `json:"exam_type"`
	ExamName        string    `json:"exam_name"`
	ExamTypeId      int       `json:"exam_type_id"`
	ExamCategory    string    `json:"exam_category"`
	ExamCategoryId  int       `json:"exam_category_id"`
	Topic           string    `json:"topic"`
	Type            string    `json:"type"`
	Attempts        []Attempt `json:"attempts"`
}

type Attempt struct {
	AttemptId        int       `json:"attempt_id"`
	AttemptNumber    int       `json:"attempt_number"`
	AssessmentStatus string    `json:"assessment_status"`
	AssessmentId     int       `json:"assessment_id"`
	AttemptDate      time.Time `json:"attempt_date"`
}

type CategoryExamType struct {
	Id           int    `json:"exam_type_id"`
	ExamName     string `json:"exam_name"`
	Description  string `json:"description"`
	TypeOfExam   string `json:"type_of_exam"`
	IsActive     bool   `json:"is_active"`
	CategoryName string `json:"category_name,omitempty"`
	CategoryId   int    `json:"category_id,omitempty"`
	LogoUrl      string `json:"logo_url"`
}
