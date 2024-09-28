package server

import (
	"errors"
	"net/http"

	"common/ent"

	"server/internal/services"
	"server/pkg/constants"
)

func (s *Server) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.HandleError(w, err, "unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := s.userService.GetUserProfile(r.Context(), userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.HandleError(w, err, "user not found", http.StatusBadRequest)
			return
		}

		s.HandleError(w, err, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Data: user,
	}

	err = s.WriteJson(w, http.StatusOK, &response)
	if err != nil {
		s.HandleError(w, err, "something went wrong", http.StatusInternalServerError)
	}
}

func (s *Server) GetUserTransactions(w http.ResponseWriter, r *http.Request) {
	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.HandleError(w, err, "unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := s.userService.GetUserTransactions(r.Context(), userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.HandleError(w, err, "user not found", http.StatusBadRequest)
			return
		}

		s.HandleError(w, err, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Data: user,
	}

	err = s.WriteJson(w, http.StatusOK, &response)
	if err != nil {
		s.HandleError(w, err, "something went wrong", http.StatusInternalServerError)
	}
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.HandleError(w, err, "unauthorized", http.StatusUnauthorized)
		return
	}

	var request services.UpdateUserRequest

	if err := s.ReadJson(w, r, &request); err != nil {
		s.HandleError(w, err, "invalid json request body", http.StatusBadRequest)
		return
	}

	user, err := s.userService.UpdateUser(r.Context(), userId, request)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.HandleError(w, err, "user not found", http.StatusBadRequest)
			return
		}

		s.HandleError(w, err, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Data: user,
	}

	err = s.WriteJson(w, http.StatusOK, &response)
	if err != nil {
		s.HandleError(w, err, "something went wrong", http.StatusInternalServerError)
	}
}
