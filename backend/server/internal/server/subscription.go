package server

import "net/http"

func (s *Server) GetAllSubscriptions(w http.ResponseWriter, r *http.Request) {
	subscriptions, err := s.subscriptionService.GetAll(r.Context())
	if err != nil {
		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	responsePayload := &Response{
		Data: subscriptions,
	}

	s.WriteJson(w, http.StatusOK, responsePayload)
}
