package server

import "net/http"

func (s *Server) CreateOrder(w http.ResponseWriter, r *http.Request) {
	body, err := s.paymentService.CreateOrder()

	if err != nil {
		s.ErrorJson(w, err)
		return
	}

	responsePayload := Response{
		Data: body,
	}

	s.WriteJson(w, http.StatusCreated, &responsePayload)
}
