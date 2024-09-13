package server

import (
	"common/ent"
	"errors"
	"net/http"
	"server/internal/services"
	"server/pkg/constants"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (s *Server) ActivateUserSubscription(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "userSubscriptionId")
	userSubscriptionId, err := strconv.Atoi(idParam)
	if err != nil {
		s.ErrorJson(w, errors.New("invalid user subscription id"), http.StatusBadRequest)
		return
	}

	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
	}

	var request services.ActivateUserSubscriptionRequest
	if err := s.ReadJson(w, r, &request); err != nil {
		s.ErrorJson(w, errors.New("invalid json request body"))
		return
	}

	if err := ValidateInput(&request); err != nil {
		s.ErrorJson(w, err)
		return
	}

	activatedSubscription, err := s.subscriptionService.ActivateUserSubscription(r.Context(), request, userSubscriptionId, userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("user subscription not found"))
			return
		}

		if strings.Contains(err.Error(), "payment verification failed") {
			s.ErrorJson(w, err)
			return
		}

		if strings.Contains(err.Error(), "already exists") {
			s.ErrorJson(w, err)
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	responsePayload := Response{
		Data: activatedSubscription,
	}

	s.WriteJson(w, http.StatusOK, &responsePayload)

}

func (s *Server) CancelUserSubscription(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "userSubscriptionId")
	userSubscriptionId, err := strconv.Atoi(idParam)
	if err != nil {
		s.ErrorJson(w, errors.New("invalid user subscription id"), http.StatusBadRequest)
		return
	}

	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
	}

	cancelledSubscription, err := s.subscriptionService.CancelUserSubscription(r.Context(), userSubscriptionId, userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("user subscription not found"))
			return
		}

		if strings.Contains(err.Error(), "user subscrption is already cancelled") {
			s.ErrorJson(w, err)
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	responsePayload := Response{
		Data: cancelledSubscription,
	}

	s.WriteJson(w, http.StatusOK, &responsePayload)
}
