package server

import (
	"encoding/json"
	"log"
	"net/http"
	"server/internal/middlewares"
	"server/pkg/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	env, err := config.LoadEnvironment()
	if err != nil {
		log.Fatalln(err)
	}

	// Start the rate limiter cleanup routine
	middlewares.StartCleanupRoutine()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)

	var allowedOrigins []string
	if env.IsProduction {
		allowedOrigins = []string{env.CorsAllowedOrigin}
	} else {
		allowedOrigins = []string{"*"}
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middlewares.RateLimiterMiddleware)

	r.Get("/sup", s.Sup)
	r.Get("/health", s.HealthCheck)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.RequireAuthMiddleware(s.authService))

		r.Route("/user", func(r chi.Router) {
			r.Get("/", s.GetUserProfile)
			r.Put("/", s.UpdateUser)
			r.Get("/transactions", s.GetUserTransactions)
		})

		r.Route("/exams", func(r chi.Router) {
			r.Route("/banking", func(r chi.Router) {
				r.Get("/descriptive", s.GetBankingDescriptiveCategories)
				r.Get("/descriptive/{id}", s.GetBankingDescriptiveQuestionsByExamId)
				r.Post("/descriptive/{id}/evaluate", s.EvaluateBankingDescriptiveExam)
			})

			r.Route("/assesments", func(r chi.Router) {
				r.Get("/{id}", s.GetAssesmentById)
			})

			r.Route("/history", func(r chi.Router) {
				r.Get("/", s.GetExamAttempts)
			})

			r.Get("/{id}", s.GetGeneratedExamById)
			r.Get("/{id}/assessments", s.GetExamAssessments)
		})

		r.Route("/subscriptions", func(r chi.Router) {
			r.Post("/{subscriptionId}/buy", s.StartSubscription)
		})

		r.Route("/user-subscriptions", func(r chi.Router) {
			r.Post("/{userSubscriptionId}/cancel", s.CancelUserSubscription)
			r.Post("/{userSubscriptionId}/activate", s.ActivateUserSubscription)
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
