package models

type PaginatedData struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	PerPage     int `json:"per_page"`
	TotalItems  int `json:"total_items"`
	Data        any `json:"data"`
}
