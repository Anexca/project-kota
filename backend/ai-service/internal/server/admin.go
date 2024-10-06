package server

import (
	"net/http"
)

func (s *Server) GenerateExamQuestionAndPopulateCache(w http.ResponseWriter, r *http.Request) {
	// examId, err := ParseIDParam(r, "id")
	// if err != nil {
	// 	s.HandleError(w, err, "invalid exam id", http.StatusBadRequest)
	// 	return
	// }

	// generatedExamData, err := s.examService.GenerateExamQuestionAndPopulateCache(r.Context(), examId)
	// if err != nil {
	// 	var notFoundError *ent.NotFoundError
	// 	if errors.As(err, &notFoundError) {
	// 		s.HandleError(w, err, "exam not found", http.StatusNotFound)
	// 		return
	// 	}

	// 	s.HandleError(w, err, "internal server error", http.StatusInternalServerError)
	// 	return
	// }

	err := s.WriteJson(w, http.StatusOK, &Response{Data: nil})
	if err != nil {
		s.HandleError(w, err, "something went wrong", http.StatusInternalServerError)
	}
}
