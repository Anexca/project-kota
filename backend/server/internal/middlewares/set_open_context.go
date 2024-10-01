package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"server/pkg/constants"
)

func SetOpenExamContext(isOpen bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), constants.OpenExamKey, fmt.Sprintf("%t", isOpen))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
