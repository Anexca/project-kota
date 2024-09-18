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

func (s *Server) GetBankingDescriptiveQuestionsByExamId(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	examId, err := strconv.Atoi(idParam)
	if err != nil {
		s.ErrorJson(w, errors.New("invalid exam id"), http.StatusBadRequest)
		return
	}

	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	cachedQuestions, err := s.examGenerationService.GetGeneratedExamsByExamId(r.Context(), examId, userId)

	if err != nil {
		if strings.Contains(err.Error(), "forbidden") {
			s.ErrorJson(w, err, http.StatusForbidden)
			return
		}

		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("exam type not found"))
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	s.WriteJson(w, http.StatusOK, &Response{Data: cachedQuestions})
}

func (s *Server) EvaluateBankingDescriptiveExam(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	generatedExamId, err := strconv.Atoi(idParam)
	if err != nil {
		s.ErrorJson(w, errors.New("invalid exam id"), http.StatusBadRequest)
		return
	}

	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	isOpenStr := r.URL.Query().Get("isopen")

	var isOpen bool
	if isOpenStr != "" {
		isOpen, err = strconv.ParseBool(isOpenStr)
		if err != nil {
			s.ErrorJson(w, errors.New("invalid isopen query param, should be either true or false"), http.StatusBadRequest)
			return
		}

	} else {
		isOpen = false
	}

	var request services.DescriptiveExamAssesmentRequest

	if err := s.ReadJson(w, r, &request); err != nil {
		s.ErrorJson(w, errors.New("invalid json request body"))
		return
	}

	if err := ValidateInput(&request); err != nil {
		s.ErrorJson(w, err, http.StatusBadRequest)
		return
	}

	attempt, err := s.examAttemptService.CheckAndAddAttempt(r.Context(), generatedExamId, userId, isOpen)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("exam not found"))
			return
		}

		if strings.Contains(err.Error(), "max attempts for exam exceeded") {
			s.ErrorJson(w, err, http.StatusBadRequest)
			return
		}

		if strings.Contains(err.Error(), "forbidden") {
			s.ErrorJson(w, err, http.StatusForbidden)
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	assesment, err := s.examAssesmentService.StartNewDescriptiveAssesment(r.Context(), generatedExamId, attempt, &request, userId, isOpen)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("exam not found"))
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	responsePayload := Response{
		Data: assesment,
	}

	s.WriteJson(w, http.StatusAccepted, &responsePayload)
}
