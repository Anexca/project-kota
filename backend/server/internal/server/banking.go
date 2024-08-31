package server

import (
	commonConstants "common/constants"

	"net/http"
)

const EXAM_CATEGORY_TYPE = commonConstants.Banking

func (s *Server) GetBankingDescriptiveQuestions(w http.ResponseWriter, r *http.Request) {
	s.WriteJson(w, http.StatusOK, &Response{})
}
