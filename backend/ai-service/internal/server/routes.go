package server

import (
	"ai-service/internal/middlewares"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)

	r.Get("/sup", s.SupHandler)
	r.Get("/health", s.HealthCheck)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.RequireAccessKeyMiddleware())
	})

	return r
}

func (s *Server) SupHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Sup",
	}

	s.WriteJson(w, http.StatusOK, &response)
}

func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.redisService.Health())
	response := Response{
		Data: string(jsonResp),
	}
	s.WriteJson(w, http.StatusOK, &response)

}
