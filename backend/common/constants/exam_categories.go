package constants

type ExamCategoryType string

const (
	Banking     ExamCategoryType = "BANKING"
	Engineering ExamCategoryType = "ENGINEERING"
	Medical     ExamCategoryType = "MEDICAL"
)

var EXAM_CATEGORIES = map[ExamCategoryType]string{
	Banking:     "banking",
	Engineering: "engineering",
	Medical:     "medical",
}
