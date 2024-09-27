package server

import (
	"common/ent"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (s *Server) GetBankingDescriptiveCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := s.examCategoryService.GetBankingDescriptiveExamsGroups(r.Context())
	if err != nil {
		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	response := Response{
		Data: categories,
	}

	s.WriteJson(w, http.StatusOK, &response)
}

func (s *Server) GetBankingMCQCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := s.examCategoryService.GetBankingMCQExamGroups(r.Context())
	if err != nil {
		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	response := Response{
		Data: categories,
	}

	s.WriteJson(w, http.StatusOK, &response)
}

func (s *Server) GetExamById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	examId, err := strconv.Atoi(idParam)
	if err != nil {
		s.ErrorJson(w, errors.New("invalid exam id"), http.StatusBadRequest)
		return
	}

	categoryExam, err := s.examCategoryService.GetExamGroupById(r.Context(), examId)
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

	s.WriteJson(w, http.StatusOK, &Response{Data: categoryExam})
}
