package constants

type ExamCategoryName string

const (
	ExamCategoryNameBanking ExamCategoryName = "BANKING"
)

var EXAM_CATEGORIES = map[ExamCategoryName]string{
	ExamCategoryNameBanking: "banking",
}
