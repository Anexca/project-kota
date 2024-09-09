package server

import (
	"ai-service/pkg/models"
	"net/http"
)

func (s *Server) GetPromptResults(w http.ResponseWriter, r *http.Request) {
	var request models.GetPromptResultsRequest

	if err := s.ReadJson(w, r, &request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid json request body"))
		return
	}

	if err := ValidateInput(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	promptResults, err := s.promptService.GetPromptResults(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(promptResults))
}
