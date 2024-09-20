package constants

type AssessmentStatusType string

const (
	ASSESSMENT_COMPLETED AssessmentStatusType = "COMPLETED"
	ASSESSMENT_REJECTED  AssessmentStatusType = "REJECTED"
	ASSESSMENT_PENDING   AssessmentStatusType = "PENDING"
)

// ExamType represents the various statuses a payment can have.
type ExamType string

const (
	ExamTypeDescriptive ExamType = "DESCRIPTIVE"
	ExamTypeMCQ         ExamType = "MCQ"
)
