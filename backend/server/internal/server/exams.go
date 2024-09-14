package server

import (
	"common/ent"
	"errors"
	"net/http"
	"server/pkg/constants"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (s *Server) GetGeneratedExamById(w http.ResponseWriter, r *http.Request) {
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

	generatedExam, err := s.examGenerationService.GetGeneratedExamById(r.Context(), generatedExamId, userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("exam not found"))
			return
		}

		if strings.Contains(err.Error(), "forbidden") {
			s.ErrorJson(w, err, http.StatusForbidden)
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	responsePayload := Response{
		Data: generatedExam,
	}

	s.WriteJson(w, http.StatusOK, &responsePayload)
}

func (s *Server) GetAssesmentById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	assesmentId, err := strconv.Atoi(idParam)
	if err != nil {
		s.ErrorJson(w, errors.New("invalid assesment id"), http.StatusBadRequest)
		return
	}

	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
	}

	assesment, err := s.examAssesmentService.GetAssesmentById(r.Context(), assesmentId, userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("assesment not found"))
			return
		}

		if strings.Contains(err.Error(), "forbidden") {
			s.ErrorJson(w, err, http.StatusForbidden)
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	responsePayload := Response{
		Data: assesment,
	}

	s.WriteJson(w, http.StatusOK, &responsePayload)
}

func (s *Server) GetExamAssessments(w http.ResponseWriter, r *http.Request) {
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

	assessments, err := s.examAssesmentService.GetExamAssessments(r.Context(), generatedExamId, userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("exam not found"))
			return
		}

		if strings.Contains(err.Error(), "forbidden") {
			s.ErrorJson(w, err, http.StatusForbidden)
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	response := &Response{
		Data: assessments,
	}

	s.WriteJson(w, http.StatusOK, response)
}

func (s *Server) GetExamAttempts(w http.ResponseWriter, r *http.Request) {
	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
	}

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	attempts, err := s.examAttemptService.GetAttempts(r.Context(), userId, page, limit)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("exam not found"))
			return
		}

		if strings.Contains(err.Error(), "forbidden") {
			s.ErrorJson(w, err, http.StatusForbidden)
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	response := &Response{
		Data: attempts,
	}

	s.WriteJson(w, http.StatusOK, response)
}
