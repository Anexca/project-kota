package server

import (
	"common/ent"
	"errors"
	"net/http"
	"server/pkg/constants"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

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

	returnUrl := r.URL.Query().Get("returnUrl")

	userSubscription, err := s.subscriptionService.StartUserSubscription(r.Context(), subscriptionId, &returnUrl, userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("subscription not found"))
			return
		}

		if strings.Contains("user already has active subscription", err.Error()) {
			s.ErrorJson(w, err)
			return
		}

		if strings.Contains("payment for subscription was not successful", err.Error()) {
			s.ErrorJson(w, err)
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

	activatedSubscription, err := s.subscriptionService.ActivateUserSubscription(r.Context(), userSubscriptionId, userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("user subscription not found"))
			return
		}

		if strings.Contains(err.Error(), "payment for subscription was not successful") {
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
