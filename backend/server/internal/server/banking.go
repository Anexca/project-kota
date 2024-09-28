package server

import (
	"common/ent"
	"errors"
	"server/internal/services"
	"server/pkg/constants"
	"strconv"
	"strings"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) EvaluateBankingDescriptiveExam(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")
	generatedExamId, err := strconv.Atoi(idParam)
	if err != nil {
		s.HandleError(w, err, "invalid exam id", http.StatusBadRequest)
		return
	}

	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.HandleError(w, err, "unauthorized", http.StatusUnauthorized)
		return
	}

	isOpenStr := r.URL.Query().Get("isopen")
	isOpen := false
	if isOpenStr != "" {
		isOpen, err = strconv.ParseBool(isOpenStr)
		if err != nil {
			s.HandleError(w, err, "invalid isopen query param, should be either true or false", http.StatusBadRequest)
			return
		}
	}

	var request services.DescriptiveExamAssesmentRequest
	if err := s.ReadJson(w, r, &request); err != nil {
		s.HandleError(w, err, "invalid json request body", http.StatusBadRequest)
		return
	}

	if err := ValidateInput(&request); err != nil {
		s.HandleError(w, err, err.Error(), http.StatusBadRequest)
		return
	}

	attempt, err := s.examAttemptService.CheckAndAddAttempt(r.Context(), generatedExamId, userId, isOpen)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.HandleError(w, err, "exam not found", http.StatusNotFound)
			return
		}

		switch {
		case strings.Contains(err.Error(), "max attempts for exam exceeded"):
			s.HandleError(w, err, err.Error(), http.StatusBadRequest)
		case strings.Contains(err.Error(), "forbidden"):
			s.HandleError(w, err, err.Error(), http.StatusForbidden)
		default:
			s.HandleError(w, err, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	assesment, err := s.examAssesmentService.StartNewDescriptiveAssesment(r.Context(), generatedExamId, attempt, &request, userId, isOpen)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.HandleError(w, err, "exam not found", http.StatusNotFound)
			return
		}

		s.HandleError(w, err, err.Error(), http.StatusInternalServerError)
		return
	}

	responsePayload := Response{Data: assesment}
	if err := s.WriteJson(w, http.StatusAccepted, &responsePayload); err != nil {
		s.HandleError(w, err, "Something went wrong", http.StatusInternalServerError)
	}
}
