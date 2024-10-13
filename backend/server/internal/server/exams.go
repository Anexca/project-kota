package server

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"common/ent"
)

func (s *Server) GetOpenQuestions(w http.ResponseWriter, r *http.Request) {
	userId, err := s.GetUserIdFromRequest(r)
	if err != nil {
		s.HandleError(w, err, err.Error(), http.StatusUnauthorized)
		return
	}

	cachedQuestions, err := s.examGenerationService.GetOpenGeneratedExams(r.Context(), "descriptive", userId)
	if err != nil {
		if strings.Contains(err.Error(), "forbidden") {
			s.HandleError(w, err, "forbidden", http.StatusForbidden)
			return
		}

		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.HandleError(w, err, "exam type not found", http.StatusNotFound)
			return
		}

		s.HandleError(w, err, "internal server error", http.StatusInternalServerError)
		return
	}

	err = s.WriteJson(w, http.StatusOK, &Response{Data: cachedQuestions})
	if err != nil {
		s.HandleError(w, err, "Something went wrong while writing the response", http.StatusInternalServerError)
	}
}

func (s *Server) GetGeneratedExamById(w http.ResponseWriter, r *http.Request) {
	generatedExamId, err := ParseIDParam(r, "id")
	if err != nil {
		s.HandleError(w, err, "invalid exam id", http.StatusBadRequest)
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

	userId, err := s.GetUserIdFromRequest(r)
	if err != nil {
		s.HandleError(w, err, err.Error(), http.StatusUnauthorized)
		return
	}

	generatedExam, err := s.examGenerationService.GetGeneratedExamById(r.Context(), generatedExamId, userId, isOpen)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.HandleError(w, err, "exam not found", http.StatusNotFound)
			return
		}

		if strings.Contains(err.Error(), "forbidden") {
			s.HandleError(w, err, "forbidden", http.StatusForbidden)
			return
		}

		s.HandleError(w, err, "internal server error", http.StatusInternalServerError)
		return
	}

	responsePayload := Response{
		Data: generatedExam,
	}

	err = s.WriteJson(w, http.StatusOK, &responsePayload)
	if err != nil {
		s.HandleError(w, err, "Something went wrong while writing the response", http.StatusInternalServerError)
	}
}

func (s *Server) GetGeneratedExamsByExamGroupId(w http.ResponseWriter, r *http.Request) {
	examGroupId, err := ParseIDParam(r, "id")
	if err != nil {
		s.HandleError(w, err, "invalid exam group id", http.StatusBadRequest)
		return
	}

	userId, err := s.GetUserIdFromRequest(r)
	if err != nil {
		s.HandleError(w, err, err.Error(), http.StatusUnauthorized)
		return
	}

	exams, err := s.examGenerationService.GetExamsByExamGroupIdAndExamType(r.Context(), examGroupId, userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.HandleError(w, err, "exam group not found", http.StatusNotFound)
			return
		}

		if strings.Contains(err.Error(), "forbidden") {
			s.HandleError(w, err, "forbidden", http.StatusForbidden)
			return
		}

		s.HandleError(w, err, "internal server error", http.StatusInternalServerError)
		return
	}

	responsePayload := Response{
		Data: exams,
	}

	err = s.WriteJson(w, http.StatusOK, &responsePayload)
	if err != nil {
		s.HandleError(w, err, "Something went wrong while writing the response", http.StatusInternalServerError)
	}
}

func (s *Server) GetAssesmentById(w http.ResponseWriter, r *http.Request) {
	assesmentId, err := ParseIDParam(r, "id")
	if err != nil {
		s.HandleError(w, err, "invalid assesment id", http.StatusBadRequest)
		return
	}

	userId, err := s.GetUserIdFromRequest(r)
	if err != nil {
		s.HandleError(w, err, err.Error(), http.StatusUnauthorized)
		return
	}

	assesment, err := s.examAssesmentService.GetAssesmentById(r.Context(), assesmentId, userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.HandleError(w, err, "assesment not found", http.StatusNotFound)
			return
		}

		if strings.Contains(err.Error(), "forbidden") {
			s.HandleError(w, err, "forbidden", http.StatusForbidden)
			return
		}

		s.HandleError(w, err, "internal server error", http.StatusInternalServerError)
		return
	}

	responsePayload := Response{
		Data: assesment,
	}

	err = s.WriteJson(w, http.StatusOK, &responsePayload)
	if err != nil {
		s.HandleError(w, err, "Something went wrong while writing the response", http.StatusInternalServerError)
	}
}

func (s *Server) GetExamAssessments(w http.ResponseWriter, r *http.Request) {
	generatedExamId, err := ParseIDParam(r, "id")
	if err != nil {
		s.HandleError(w, err, "invalid exam id", http.StatusBadRequest)
		return
	}

	userId, err := s.GetUserIdFromRequest(r)
	if err != nil {
		s.HandleError(w, err, err.Error(), http.StatusUnauthorized)
		return
	}

	assessments, err := s.examAssesmentService.GetExamAssessments(r.Context(), generatedExamId, userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.HandleError(w, err, "exam not found", http.StatusNotFound)
			return
		}

		if strings.Contains(err.Error(), "forbidden") {
			s.HandleError(w, err, "forbidden", http.StatusForbidden)
			return
		}

		s.HandleError(w, err, "internal server error", http.StatusInternalServerError)
		return
	}

	response := &Response{
		Data: assessments,
	}

	err = s.WriteJson(w, http.StatusOK, response)
	if err != nil {
		s.HandleError(w, err, "Something went wrong while writing the response", http.StatusInternalServerError)
	}
}

func (s *Server) GetExamAttempts(w http.ResponseWriter, r *http.Request) {
	userId, err := s.GetUserIdFromRequest(r)
	if err != nil {
		s.HandleError(w, err, err.Error(), http.StatusUnauthorized)
		return
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

	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")

	var from, to *time.Time

	if fromStr != "" {
		if fromParsed, err := time.Parse(time.RFC3339, fromStr); err == nil {
			from = &fromParsed
		} else {
			s.HandleError(w, err, "invalid 'from' date format, expected RFC3339", http.StatusBadRequest)
			return
		}
	}

	if toStr != "" {
		if toParsed, err := time.Parse(time.RFC3339, toStr); err == nil {
			to = &toParsed
		} else {
			s.HandleError(w, err, "invalid 'to' date format, expected RFC3339", http.StatusBadRequest)
			return
		}
	}

	examTypeIdStr := r.URL.Query().Get("examTypeId")
	categoryIdStr := r.URL.Query().Get("categoryID")

	var examTypeId, categoryID *int

	if examTypeIdStr != "" {
		if examTypeParsed, err := strconv.Atoi(examTypeIdStr); err == nil && examTypeParsed > 0 {
			examTypeId = &examTypeParsed
		} else {
			s.HandleError(w, err, "invalid 'examTypeId' format, expected positive integer", http.StatusBadRequest)
			return
		}
	}

	if categoryIdStr != "" {
		if categoryParsed, err := strconv.Atoi(categoryIdStr); err == nil && categoryParsed > 0 {
			categoryID = &categoryParsed
		} else {
			s.HandleError(w, err, "invalid 'categoryID' format, expected positive integer", http.StatusBadRequest)
			return
		}
	}

	paginatedAttempts, err := s.examAttemptService.GetAttempts(r.Context(), userId, page, limit, from, to, examTypeId, categoryID)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.HandleError(w, err, "exam not found", http.StatusNotFound)
			return
		}

		if strings.Contains(err.Error(), "forbidden") {
			s.HandleError(w, err, "forbidden", http.StatusForbidden)
			return
		}

		s.HandleError(w, err, "internal server error", http.StatusInternalServerError)
		return
	}

	response := &Response{
		Data: paginatedAttempts.Data,
		Pagination: ResponsePagination{
			CurrentPage: paginatedAttempts.CurrentPage,
			TotalPages:  paginatedAttempts.TotalPages,
			PerPage:     paginatedAttempts.PerPage,
			TotalItems:  paginatedAttempts.TotalItems,
		},
	}

	err = s.WriteJson(w, http.StatusOK, response)
	if err != nil {
		s.HandleError(w, err, "Something went wrong while writing the response", http.StatusInternalServerError)
	}
}
