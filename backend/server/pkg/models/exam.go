package models

import "time"

type ExamModelType int

const (
	DescriptiveExamType       ExamModelType = iota
	MCQExamType               ExamModelType = iota
	GeneratedExamOverviewType ExamModelType = iota
)

type DescriptiveExam struct {
	Type                    string   `json:"type" validate:"required,oneof=formal_letter essay"`
	Topic                   string   `json:"topic" validate:"required"`
	Hints                   []string `json:"hints" validate:"required"`
	MaxNumberOfWordsAllowed string   `json:"max_number_of_words" validate:"required"`
	TotalMarks              string   `json:"total_marks" validate:"required"`
}

type GeneratedMCQExam struct {
	Questions     []MCQExamQuestion     `json:"questions"`
	ContentGroups []MCQExamContentGroup `json:"content_groups"`
}

type MCQExamQuestion struct {
	ContentReferenceId string   `json:"content_reference_id"`
	Question           string   `json:"question"`
	QuestionNumber     int      `json:"question_number"`
	Answer             int      `json:"answer"`
	Options            []string `json:"options"`
	Explanation        string   `json:"explanation"`
}

type MCQExamContentGroup struct {
	ContentId    string      `json:"content_id"`
	Instructions string      `json:"instructions"`
	Content      interface{} `json:"content"` // String or Object depending on the content type
}

type GeneratedExamOverview struct {
	Id                  int                    `json:"exam_id"`
	ExamType            string                 `json:"exam_type"`
	ExamName            string                 `json:"exam_name"`
	ExamStage           string                 `json:"exam_stage"`
	IsSectional         bool                   `json:"is_sectional"`
	RawExamData         map[string]interface{} `json:"raw_exam_data,omitempty"`
	UserAttempts        int                    `json:"user_attempts"`
	MaxAttempts         int                    `json:"max_attempts"`
	DurationSeconds     int                    `json:"duration_seconds"`
	NumberOfQuestions   int                    `json:"number_of_questions"`
	NegativeMarking     float64                `json:"negative_marking,omitempty"`
	UserHasAccessToExam bool                   `json:"has_access"`
	CreatedAt           time.Time              `json:"created_at"`
	UpdatedAt           time.Time              `json:"updated_at"`
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

type CategoryExamGroup struct {
	Id           int    `json:"exam_group_id"`
	ExamName     string `json:"exam_group_name"`
	Description  string `json:"description"`
	TypeOfExam   string `json:"type_of_exam,omitempty"`
	IsActive     bool   `json:"is_active"`
	CategoryName string `json:"category_name,omitempty"`
	CategoryId   int    `json:"category_id,omitempty"`
	LogoUrl      string `json:"logo_url"`
}
