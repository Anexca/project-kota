package server

import "net/http"

func (s *Server) GetJEEPhysicsQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := s.questionService.GenerateMCQsForSubject(r.Context(), "JEE Mains", "physics", 5)
	if err != nil {
		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	response := Response{
		Data: questions,
	}

	s.WriteJson(w, http.StatusOK, &response)
}
