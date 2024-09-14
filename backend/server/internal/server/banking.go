package server

import (
	commonConstants "common/constants"
	"common/ent"
	"errors"
	"server/internal/services"
	"server/pkg/constants"
	"strconv"
	"strings"

	"net/http"

	"github.com/go-chi/chi/v5"
)

const EXAM_CATEGORY_TYPE = commonConstants.Banking

func (s *Server) GetBankingDescriptiveQuestions(w http.ResponseWriter, r *http.Request) {
	const EXAM_TYPE = commonConstants.Descriptive

	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
	}

	cachedQuestions, err := s.examGenerationService.GetGeneratedExams(r.Context(), EXAM_TYPE, userId)
	if err != nil {
		if strings.Contains(err.Error(), "forbidden") {
			s.ErrorJson(w, err, http.StatusForbidden)
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

	attempt, err := s.examAttemptService.CheckAndAddAttempt(r.Context(), generatedExamId, userId)
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

	assesment, err := s.examAssesmentService.StartNewDescriptiveAssesment(r.Context(), generatedExamId, attempt, &request, userId)
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
