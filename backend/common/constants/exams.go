package constants

type ExamType string
type AssessmentStatus string

const (
	Descriptive     ExamType = "DESCRIPTIVE"
	IBPSDescriptive ExamType = "IBPS_DESCRIPTIVE"
)

var EXAMS = map[ExamType]string{
	Descriptive:     "descriptive",
	IBPSDescriptive: "descriptive_ibps_po",
}

const (
	ASSESSMENT_COMPLETED AssessmentStatus = "COMPLETED"
	ASSESSMENT_REJECTED  AssessmentStatus = "REJECTED"
	ASSESSMENT_PENDING   AssessmentStatus = "PENDING"
)
