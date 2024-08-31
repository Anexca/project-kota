package constants

type ExamType string

const (
	Descriptive ExamType = "DESCRIPTIVE"
)

var EXAMS = map[ExamType]string{
	Descriptive: "descriptive",
}
