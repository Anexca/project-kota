package server

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"common/ent"
	"github.com/go-chi/chi/v5"
)

func (s *Server) GetBankingExamGroups(w http.ResponseWriter, r *http.Request) {
	categories, err := s.examCategoryService.GetBankingExamGroups(r.Context())
	if err != nil {
		s.HandleError(w, err, "failed to retrieve exam categories", http.StatusInternalServerError)
		return
	}

	response := Response{
		Data: categories,
	}

	err = s.WriteJson(w, http.StatusOK, &response)
	if err != nil {
		s.HandleError(w, err, "failed to write response", http.StatusInternalServerError)
	}
}

func (s *Server) GetExamById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	examId, err := strconv.Atoi(idParam)
	if err != nil {
		s.HandleError(w, err, "invalid exam id", http.StatusBadRequest)
		return
	}

	categoryExam, err := s.examCategoryService.GetExamGroupById(r.Context(), examId)
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

	err = s.WriteJson(w, http.StatusOK, &Response{Data: categoryExam})
	if err != nil {
		s.HandleError(w, err, "Something went wrong while sending the response", http.StatusInternalServerError)
	}
}
