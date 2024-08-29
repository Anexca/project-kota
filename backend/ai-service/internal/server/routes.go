package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/sup", s.SupHandler)
	r.Get("/health", s.HealthCheckHandler)

	r.Route("/questions", func(r chi.Router) {
		r.Route("/jee", func(r chi.Router) {
			r.Get("/mcq/physics", s.GetJEEPhysicsMCQs)
			r.Get("/nvq/physics", s.GetJEEPhysicsNVQs)
		})

		r.Route("/banking", func(r chi.Router) {
			r.Get("/descriptive", s.GetDescriptiveQuestions)
		})
	})

	return r
}

func (s *Server) SupHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Sup",
	}

	s.WriteJson(w, http.StatusOK, &response)
}

func (s *Server) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.redisService.Health())
	response := Response{
		Data: string(jsonResp),
	}
	s.WriteJson(w, http.StatusOK, &response)

}
