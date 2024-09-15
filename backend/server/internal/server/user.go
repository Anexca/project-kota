package server

import (
	"common/ent"
	"errors"
	"net/http"
	"server/internal/services"
	"server/pkg/constants"
)

func (s *Server) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	user, err := s.userService.GetUserProfile(r.Context(), userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("user not found"))
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	response := Response{
		Data: user,
	}

	s.WriteJson(w, http.StatusOK, &response)
}

func (s *Server) GetUserTransactions(w http.ResponseWriter, r *http.Request) {
	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	user, err := s.userService.GetUserTransactions(r.Context(), userId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("user not found"))
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	response := Response{
		Data: user,
	}

	s.WriteJson(w, http.StatusOK, &response)
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

	userId, err := GetHttpRequestContextValue(r, constants.UserIDKey)
	if err != nil {
		s.ErrorJson(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	var request services.UpdateUserRequest

	if err := s.ReadJson(w, r, &request); err != nil {
		s.ErrorJson(w, errors.New("invalid json request body"))
		return
	}

	user, err := s.userService.UpdateUser(r.Context(), userId, request)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			s.ErrorJson(w, errors.New("user not found"))
			return
		}

		s.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	response := Response{
		Data: user,
	}

	s.WriteJson(w, http.StatusOK, &response)
}
