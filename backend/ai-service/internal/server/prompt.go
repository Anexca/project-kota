package server

import (
	"ai-service/pkg/models"
	"errors"
	"net/http"
)

func (s *Server) GetPromptResults(w http.ResponseWriter, r *http.Request) {
	var request models.GetPromptResultsRequest

	if err := s.ReadJson(w, r, &request); err != nil {
		s.ErrorJson(w, errors.New("invalid json request body"))
		return
	}

	if err := ValidateInput(&request); err != nil {
		s.ErrorJson(w, err, http.StatusBadRequest)
		return
	}

	promptResults, err := s.promptService.GetPromptResults(r.Context(), &request)
	if err != nil {
		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	responsePayload := Response{
		Data: promptResults,
	}

	s.WriteJson(w, http.StatusOK, &responsePayload)
}
