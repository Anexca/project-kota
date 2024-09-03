package models

import "time"

type ExamModelType int

const (
	DescriptiveExamType       ExamModelType = iota
	GeneratedExamOverviewType ExamModelType = iota
)

type DescriptiveExam struct {
	Type  string   `json:"type"`
	Topic string   `json:"topic"`
	Hints []string `json:"hints"`
}

type GeneratedExamOverview struct {
	Id                int                    `json:"id"`
	RawExamData       map[string]interface{} `json:"raw_exam_data"`
	UserAttempts      int                    `json:"user_attempts"`
	MaxAttempts       int                    `json:"max_attempts"`
	DurationMinutes   int                    `json:"duration_minutes"`
	NumberOfQuestions int                    `json:"number_of_questions"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
}
