package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"ai-service/internal/middlewares"
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
		r.Post("/prompt", s.GetPromptResults)
		r.Post("/prompt/structured", s.GetStructuredPromptResults)
	})

	r.Route("/admin", func(r chi.Router) {
		r.Use(middlewares.RequireAdminKeyMiddleware())
		r.Route("/generate", func(r chi.Router) {
			r.Post("/exam/{id}", s.GenerateExamQuestionAndPopulateCache)
			r.Post("/all/descriptive", s.GenerateAllDescriptiveQuestions)
		})
	})

	return r
}

func (s *Server) SupHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Sup",
	}

	err := s.WriteJson(w, http.StatusOK, &response)
	if err != nil {
		s.HandleError(w, err, "something went wrong", http.StatusInternalServerError)
	}
}

func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.redisService.Health())
	response := Response{
		Data: string(jsonResp),
	}
	err := s.WriteJson(w, http.StatusOK, &response)
	if err != nil {
		s.HandleError(w, err, "something went wrong", http.StatusInternalServerError)
	}
}
