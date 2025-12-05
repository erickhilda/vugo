package middleware

import (
	"net/http"
)

// RequireAuth is a middleware that checks if user is authenticated
// This will be implemented in Milestone 4 when we add authentication
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement authentication check in Milestone 4
		// For now, just pass through
		next.ServeHTTP(w, r)
	})
}

// RequireProjectAccess checks if user has access to a project
// This will be implemented in Milestone 5 when we add projects
func RequireProjectAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement project access check in Milestone 5
		next.ServeHTTP(w, r)
	})
}
