package server

import (
	"ai-service/pkg/models"
	"net/http"
	"time"

	"go.uber.org/ratelimit"
)

var rl = ratelimit.New(1, ratelimit.Per(time.Minute))

func (s *Server) GetPromptResults(w http.ResponseWriter, r *http.Request) {
	rl.Take()

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
