package server

import (
	"common/ent"
	"errors"
	"net/http"
	"server/pkg/constants"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

func (s *Server) GetOpenQuestions(w http.ResponseWriter, r *http.Request) {
	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
	}

	cachedQuestions, err := s.examGenerationService.GetOpenGeneratedExams(r.Context(), "descriptive", userId)

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

func (s *Server) GetGeneratedExamById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	generatedExamId, err := strconv.Atoi(idParam)
	if err != nil {
		s.ErrorJson(w, errors.New("invalid exam id"), http.StatusBadRequest)
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

	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
	}

	generatedExam, err := s.examGenerationService.GetGeneratedExamById(r.Context(), generatedExamId, userId, isOpen)
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
		return
	}

	// Pagination parameters
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

	// Date range parameters (optional)
	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")

	var from, to *time.Time

	if fromStr != "" {
		if fromParsed, err := time.Parse(time.RFC3339, fromStr); err == nil {
			from = &fromParsed
		} else {
			s.ErrorJson(w, errors.New("invalid 'from' date format, expected RFC3339"), http.StatusBadRequest)
			return
		}
	}

	if toStr != "" {
		if toParsed, err := time.Parse(time.RFC3339, toStr); err == nil {
			to = &toParsed
		} else {
			s.ErrorJson(w, errors.New("invalid 'to' date format, expected RFC3339"), http.StatusBadRequest)
			return
		}
	}

	// Exam type and category filters (optional)
	examTypeIdStr := r.URL.Query().Get("examTypeId")
	categoryIdStr := r.URL.Query().Get("categoryID")

	var examTypeId, categoryID *int

	if examTypeIdStr != "" {
		if examTypeParsed, err := strconv.Atoi(examTypeIdStr); err == nil && examTypeParsed > 0 {
			examTypeId = &examTypeParsed
		} else {
			s.ErrorJson(w, errors.New("invalid 'examTypeId' format, expected positive integer"), http.StatusBadRequest)
			return
		}
	}

	if categoryIdStr != "" {
		if categoryParsed, err := strconv.Atoi(categoryIdStr); err == nil && categoryParsed > 0 {
			categoryID = &categoryParsed
		} else {
			s.ErrorJson(w, errors.New("invalid 'categoryID' format, expected positive integer"), http.StatusBadRequest)
			return
		}
	}

	attempts, err := s.examAttemptService.GetAttempts(r.Context(), userId, page, limit, from, to, examTypeId, categoryID)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("exam not found"), http.StatusNotFound)
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
