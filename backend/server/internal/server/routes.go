package server

import (
	"encoding/json"
	"net/http"
	"server/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/sup", s.Sup)
	r.Get("/health", s.HealthCheck)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.RequireAuthMiddleware(s.authService))

		r.Route("/exams", func(r chi.Router) {
			r.Route("/banking", func(r chi.Router) {
				r.Get("/descriptive", s.GetBankingDescriptiveQuestions)
			})
		})
	})

	return r
}

func (s *Server) Sup(w http.ResponseWriter, r *http.Request) {
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
