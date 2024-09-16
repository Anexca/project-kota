package constants

type AssessmentStatusType string

const (
	Descriptive     ExamType = "DESCRIPTIVE"
	IBPSDescriptive ExamType = "IBPS_DESCRIPTIVE"
)

var EXAMS = map[ExamType]string{
	Descriptive:     "descriptive",
	IBPSDescriptive: "descriptive_ibps_po",
}

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
