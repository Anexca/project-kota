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
