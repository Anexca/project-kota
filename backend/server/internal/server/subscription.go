package server

import (
	"net/http"
)

func (s *Server) GetAllSubscriptions(w http.ResponseWriter, r *http.Request) {
	subscriptions, err := s.subscriptionService.GetAll(r.Context())
	if err != nil {
		s.HandleError(w, err, err.Error(), http.StatusInternalServerError)
		return
	}

	responsePayload := &Response{
		Data: subscriptions,
	}

	err = s.WriteJson(w, http.StatusOK, responsePayload)
	if err != nil {
		s.HandleError(w, err, "something went wrong", http.StatusInternalServerError)
	}
}
