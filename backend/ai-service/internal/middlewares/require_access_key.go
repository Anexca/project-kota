package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"ai-service/pkg/config"
)

func RequireAccessKeyMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			environment, err := config.LoadEnvironment()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			accessKey, err := extractAccessKeyFromHeader(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if accessKey != environment.ServerAccessKey {
				http.Error(w, errors.New("invalid access key").Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func extractAccessKeyFromHeader(r *http.Request) (string, error) {
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
