package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"ai-service/pkg/config"
)

func RequireAdminKeyMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			environment, err := config.LoadEnvironment()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			accessKey, err := extractAdminKeyFromHeader(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if accessKey != environment.ServerAdminKey {
				http.Error(w, errors.New("invalid admin key").Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func extractAdminKeyFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header missing")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("authorization header format must be Bearer <token>")
	}

	accessKey := parts[1]
	return accessKey, nil
}
