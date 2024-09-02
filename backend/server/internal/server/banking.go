package server

import (
	commonConstants "common/constants"

	"net/http"
)

const EXAM_CATEGORY_TYPE = commonConstants.Banking

func (s *Server) GetBankingDescriptiveQuestions(w http.ResponseWriter, r *http.Request) {
	const EXAM_TYPE = commonConstants.Descriptive

	cachedQuestions, err := s.examService.GetGeneratedExams(r.Context(), EXAM_TYPE)
	if err != nil {
		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	s.WriteJson(w, http.StatusOK, &Response{Data: cachedQuestions})
}
