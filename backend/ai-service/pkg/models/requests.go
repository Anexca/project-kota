package models

type GetPromptResultsRequest struct {
	Prompt      string                 `json:"prompt" validate:"required"`
	Model       string                 `json:"model" validate:"required,oneof=gemini-1.5-pro gemini-1.5-flash"`
}
