package server

import "net/http"

func (s *Server) GetJEEPhysicsMCQs(w http.ResponseWriter, r *http.Request) {
	questions, err := s.questionService.GenerateQuestions(r.Context(), "Multiple Choice Question", "JEE_Mains", "physics", 5)
	if err != nil {
		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	response := Response{
		Data: questions,
	}

	s.WriteJson(w, http.StatusOK, &response)
}

func (s *Server) GetJEEPhysicsNVQs(w http.ResponseWriter, r *http.Request) {
	questions, err := s.questionService.GenerateQuestions(r.Context(), "Numerical Value Question", "JEE_Mains", "physics", 5)
	if err != nil {
		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	response := Response{
		Data: questions,
	}

	s.WriteJson(w, http.StatusOK, &response)
}
