package server

import (
	"log"
	"net/http"
	"time"

	"go.uber.org/ratelimit"

	"ai-service/pkg/models"
)

// Separate rate limiters for each function
var rlPromptResults = ratelimit.New(1, ratelimit.Per(time.Second))
var rlStructuredPromptResults = ratelimit.New(1, ratelimit.Per(time.Second))

func (s *Server) GetPromptResults(w http.ResponseWriter, r *http.Request) {
	rlPromptResults.Take()
	log.Println("GetPromptResults request allowed at", time.Now())

	var request models.GetPromptResultsRequest

	if err := s.ReadJson(w, r, &request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("invalid json request body"))
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
		}
		return
	}

	if err := ValidateInput(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
		}
		return
	}

	promptResults, err := s.promptService.GetPromptResults(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
		}
		return
	}

	_, err = w.Write([]byte(promptResults))
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}

func (s *Server) GetStructuredPromptResults(w http.ResponseWriter, r *http.Request) {
	rlStructuredPromptResults.Take()
	log.Println("GetStructuredPromptResults request allowed at", time.Now())

	var request models.GetPromptResultsRequest

	if err := s.ReadJson(w, r, &request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("invalid json request body"))
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
		}
		return
	}

	if err := ValidateInput(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
		}
		return
	}

	promptResults, err := s.promptService.GetStructuredPromptResults(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
		}
		return
	}

	_, err = w.Write([]byte(promptResults))
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}
