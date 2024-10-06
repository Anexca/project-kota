package models

type GetPromptResultsRequest struct {
	Prompt string `json:"prompt" validate:"required"`
	Model  string `json:"model" validate:"required,oneof=gemini-1.5-pro gemini-1.5-flash gemini-1.0-pro"`
}

type GenerateQuestionResponse struct {
	ExamName         string `json:"exam_name"`
	CachedMetaDataId int    `json:"cached_meta_data_id"`
}
