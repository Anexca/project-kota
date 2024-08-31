package server

import (
	commonConstants "common/constants"

	"net/http"
)

const EXAM_CATEGORY_TYPE = commonConstants.Banking

func (s *Server) GetBankingDescriptiveQuestions(w http.ResponseWriter, r *http.Request) {
	const EXAM_TYPE = commonConstants.Descriptive

	// Define the type for the questions
	type DescriptiveQuestion struct {
		Type  string   `json:"type"`
		Topic string   `json:"topic"`
		Hints []string `json:"hints"`
	}

	// Create a variable of the type you want to use for unmarshaling
	var descriptiveQuestions []DescriptiveQuestion

	// Fetch the cached questions
	cachedQuestions, err := s.examService.GetCachedQuestions(r.Context(), EXAM_TYPE, &descriptiveQuestions)
	if err != nil {
		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	s.WriteJson(w, http.StatusOK, &Response{Data: cachedQuestions})
}
