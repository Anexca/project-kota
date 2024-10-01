package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"server/internal/middlewares"
	"server/pkg/config"
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
				r.Route("/descriptive", func(r chi.Router) {
					r.Get("/open", s.GetOpenQuestions)
					r.Post("/{id}/evaluate", s.EvaluateBankingDescriptiveExam)
				})

				r.Route("/mcq", func(r chi.Router) {
					r.Post("/{id}/evaluate", s.EvaluateBankingMCQExam)
				})

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", s.GetGeneratedExamsByExamGroupId)
				})
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
	})

	r.Route("/webhook", func(r chi.Router) {
		r.Post("/subscription/payment-success", s.HandleSubscriptionPaymentSuccess)
	})

	r.Get("/subscriptions", s.GetAllSubscriptions)
	r.Get("/categories/exams/{id}", s.GetExamById)
	r.Get("/exams/banking", s.GetBankingExamGroups)

	return r
}

func (s *Server) Sup(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Sup",
	}

	err := s.WriteJson(w, http.StatusOK, &response)
	if err != nil {
		s.HandleError(w, err, "Something went wrong while writing the response", http.StatusInternalServerError)
	}
}

func (s *Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.redisService.Health())
	response := Response{
		Data: string(jsonResp),
	}
	err := s.WriteJson(w, http.StatusOK, &response)
	if err != nil {
		s.HandleError(w, err, "Something went wrong while writing the response", http.StatusInternalServerError)
	}
}
