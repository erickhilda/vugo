package middleware

import (
	"context"
	"net/http"

	"github.com/erickhilda/vugo/internal/database/queries"
	"github.com/erickhilda/vugo/internal/services"
)

// contextKey is a type for context keys
type contextKey string

const (
	// UserContextKey is the key for storing user in context
	UserContextKey contextKey = "user"
)

// AuthMiddleware wraps the auth service
type AuthMiddleware struct {
	authService *services.AuthService
}

// NewAuthMiddleware creates a new auth middleware
func NewAuthMiddleware(authService *services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

// RequireAuth is a middleware that checks if user is authenticated
func (m *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get session ID from cookie
		cookie, err := r.Cookie("session_id")
		if err != nil {
			// No session cookie, redirect to login
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Get user from session
		user, err := m.authService.GetUserBySession(r.Context(), cookie.Value)
		if err != nil {
			// Invalid or expired session, redirect to login
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Inject user into context
		ctx := context.WithValue(r.Context(), UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// OptionalAuth is a middleware that optionally loads user if session exists
func (m *AuthMiddleware) OptionalAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get session ID from cookie
		cookie, err := r.Cookie("session_id")
		if err != nil {
			// No session, continue without user
			next.ServeHTTP(w, r)
			return
		}

		// Get user from session (ignore errors)
		user, err := m.authService.GetUserBySession(r.Context(), cookie.Value)
		if err != nil {
			// Invalid session, continue without user
			next.ServeHTTP(w, r)
			return
		}

		// Inject user into context
		ctx := context.WithValue(r.Context(), UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserFromContext retrieves the user from request context
func GetUserFromContext(ctx context.Context) (*queries.User, bool) {
	user, ok := ctx.Value(UserContextKey).(*queries.User)
	return user, ok
}
