package middlewares

import (
	"net/http"
	"server/pkg/constants"
)

// Middleware to check if a user has access to a specific exam
func RequireExamAccessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Step 1: Extract the user ID (from session, token, etc.)
		userID := r.Context().Value(constants.UserIDKey).(string)

		// Step 2: Extract the exam ID from the URL or request context
		examID := r.URL.Query().Get("exam_id") // or extract from URL params

		// Step 3: Check if the user has access to the exam
		hasAccess, err := CanUserAccessExam(userID, examID)
		if err != nil {
			http.Error(w, "Failed to check exam access", http.StatusInternalServerError)
			return
		}

		if !hasAccess {
			http.Error(w, "You do not have access to this exam", http.StatusForbidden)
			return
		}

		// Step 4: If access is allowed, call the next handler
		next.ServeHTTP(w, r)
	})
}

// Example function to check if a user can access an exam
func CanUserAccessExam(userID string, examID string) (bool, error) {
	// Check in the database if the user has an active subscription for the exam
	// This is where you'd query your database or cache to verify permissions

	// Example: Assume user has access to the exam
	return true, nil
}
