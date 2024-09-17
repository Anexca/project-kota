package constants

type ExamCategoryName string

const (
	ExamCategoryTypeName ExamCategoryName = "BANKING"
)

var EXAM_CATEGORIES = map[ExamCategoryName]string{
	ExamCategoryTypeName: "banking",
}
