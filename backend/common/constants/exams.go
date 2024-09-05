package constants

type ExamType string
type AssessmentStatus string

const (
	Descriptive ExamType = "DESCRIPTIVE"
)

var EXAMS = map[ExamType]string{
	Descriptive: "descriptive",
}

const (
	ASSESSMENT_COMPLETED AssessmentStatus = "COMPLETED"
	ASSESSMENT_REJECTED  AssessmentStatus = "REJECTED"
	ASSESSMENT_PENDING   AssessmentStatus = "PENDING"
)
