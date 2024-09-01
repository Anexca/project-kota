package server

import (
	commonConstants "common/constants"
	"server/pkg/models"

	"net/http"
)

const EXAM_CATEGORY_TYPE = commonConstants.Banking

func (s *Server) GetBankingDescriptiveQuestions(w http.ResponseWriter, r *http.Request) {
	const EXAM_TYPE = commonConstants.Descriptive

	var descriptiveQuestions []models.DescriptiveQuestion

	cachedQuestions, err := s.examService.GetCachedQuestions(r.Context(), EXAM_TYPE, &descriptiveQuestions)
	if err != nil {
		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	s.WriteJson(w, http.StatusOK, &Response{Data: cachedQuestions})
}
