package server

import (
	"common/ent"
	"errors"
	"net/http"
	"server/pkg/constants"
	"strconv"

	"github.com/go-chi/chi/v5"
)

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

func (s *Server) StartSubscription(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "subscriptionId")
	subscriptionId, err := strconv.Atoi(idParam)
	if err != nil {
		s.ErrorJson(w, errors.New("invalid subscription id"), http.StatusBadRequest)
		return
	}

	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
	}

	userSubscription, err := s.subscriptionService.StartUserSubscription(r.Context(), subscriptionId, userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("subscription not found"))
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	responsePayload := Response{
		Data: userSubscription,
	}

	s.WriteJson(w, http.StatusCreated, &responsePayload)
}
