package models

type DescriptiveQuestion struct {
	Type  string   `json:"type"`
	Topic string   `json:"topic"`
	Hints []string `json:"hints"`
}
