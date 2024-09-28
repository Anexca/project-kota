package server

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"common/ent"
	"github.com/go-chi/chi/v5"

	"server/pkg/constants"
)

func (s *Server) StartSubscription(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "subscriptionId")
	subscriptionId, err := strconv.Atoi(idParam)
	if err != nil {
		s.HandleError(w, err, "invalid subscription id", http.StatusBadRequest)
		return
	}

	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.HandleError(w, err, "unauthorized", http.StatusUnauthorized)
		return
	}

	returnUrl := r.URL.Query().Get("returnUrl")

	userSubscription, err := s.subscriptionService.StartUserSubscription(r.Context(), subscriptionId, &returnUrl, userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.HandleError(w, err, "subscription not found", http.StatusBadRequest)
			return
		}

		if strings.Contains(err.Error(), "user already has active subscription") {
			s.HandleError(w, err, "user already has active subscription", http.StatusBadRequest)
			return
		}

		if strings.Contains(err.Error(), "payment for subscription was not successful") {
			s.HandleError(w, err, "payment for subscription was not successful", http.StatusBadRequest)
			return
		}

		s.HandleError(w, err, err.Error(), http.StatusInternalServerError)
		return
	}

	responsePayload := Response{
		Data: userSubscription,
	}

	err = s.WriteJson(w, http.StatusCreated, &responsePayload)
	if err != nil {
		s.HandleError(w, err, "something went wrong", http.StatusInternalServerError)
	}
}
