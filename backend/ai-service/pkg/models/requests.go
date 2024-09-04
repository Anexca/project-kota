package models

type GetPromptResultsRequest struct {
	Prompt string `json:"prompt" validate:"required"`
	Model  string `json:"model" validate:"required,oneof=pro flash"`
}
