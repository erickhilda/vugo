package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/erickhilda/vugo/internal/database/queries"
	"github.com/erickhilda/vugo/internal/middleware"
	"github.com/erickhilda/vugo/internal/services"
)

// APIAuthHandlers handles authentication API routes
type APIAuthHandlers struct {
	authService *services.AuthService
}

// NewAPIAuthHandlers creates a new API auth handlers instance
func NewAPIAuthHandlers(authService *services.AuthService) *APIAuthHandlers {
	return &APIAuthHandlers{
		authService: authService,
	}
}

// API Response types

// ApiSuccessResponse is a generic success response
type ApiSuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

// ApiErrorResponse is a generic error response
type ApiErrorResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   ErrorDetail `json:"error"`
}

// ErrorDetail contains error information
type ErrorDetail struct {
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Details interface{} `json:"details,omitempty"`
}

// Request/Response types

// LoginRequest represents a login request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterRequest represents a registration request
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// UserResponse represents a user in API responses
type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// AuthResponse represents authentication response data
type AuthResponse struct {
	User UserResponse `json:"user"`
}

// Helper functions

// sendSuccess sends a successful JSON response
func sendSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ApiSuccessResponse{
		Success: true,
		Data:    data,
		Error:   nil,
	})
}

// sendError sends an error JSON response
func sendError(w http.ResponseWriter, statusCode int, message, code string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ApiErrorResponse{
		Success: false,
		Data:    nil,
		Error: ErrorDetail{
			Message: message,
			Code:    code,
		},
	})
}

// userToResponse converts a database user to API response format
func userToResponse(user *queries.User) UserResponse {
	createdAt := ""
	if user.CreatedAt.Valid {
		createdAt = user.CreatedAt.Time.Format("2006-01-02T15:04:05Z")
	}

	updatedAt := ""
	if user.UpdatedAt.Valid {
		updatedAt = user.UpdatedAt.Time.Format("2006-01-02T15:04:05Z")
	}

	return UserResponse{
		ID:        fmt.Sprintf("%d", user.ID),
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

// API Handlers

// HandleLogin processes user login via JSON API
func (h *APIAuthHandlers) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, http.StatusMethodNotAllowed, "Method not allowed", "METHOD_NOT_ALLOWED")
		return
	}

	// Parse JSON request
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request body", "INVALID_REQUEST_BODY")
		return
	}

	// Validate
	if req.Email == "" || req.Password == "" {
		sendError(w, http.StatusBadRequest, "Email and password are required", "VALIDATION_ERROR")
		return
	}

	// Login user
	result, err := h.authService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		sendError(w, http.StatusUnauthorized, err.Error(), "AUTH_INVALID_CREDENTIALS")
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    result.Session.ID,
		Path:     "/",
		HttpOnly: true,
		Secure:   os.Getenv("ENV") == "production",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   7 * 24 * 60 * 60, // 7 days
	})

	// Send success response
	sendSuccess(w, AuthResponse{
		User: userToResponse(&result.User),
	})
}

// HandleRegister processes user registration via JSON API
func (h *APIAuthHandlers) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, http.StatusMethodNotAllowed, "Method not allowed", "METHOD_NOT_ALLOWED")
		return
	}

	// Parse JSON request
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request body", "INVALID_REQUEST_BODY")
		return
	}

	// Validate
	if req.Email == "" || req.Password == "" || req.Name == "" {
		sendError(w, http.StatusBadRequest, "Email, password, and name are required", "VALIDATION_ERROR")
		return
	}

	// Register user
	result, err := h.authService.Register(r.Context(), req.Email, req.Password, req.Name)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error(), "REGISTRATION_ERROR")
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    result.Session.ID,
		Path:     "/",
		HttpOnly: true,
		Secure:   os.Getenv("ENV") == "production",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   7 * 24 * 60 * 60, // 7 days
	})

	// Send success response
	sendSuccess(w, AuthResponse{
		User: userToResponse(&result.User),
	})
}

// HandleLogout processes user logout via JSON API
func (h *APIAuthHandlers) HandleLogout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, http.StatusMethodNotAllowed, "Method not allowed", "METHOD_NOT_ALLOWED")
		return
	}

	// Get session cookie
	cookie, err := r.Cookie("session_id")
	if err == nil {
		// Delete session
		_ = h.authService.Logout(r.Context(), cookie.Value)
	}

	// Delete cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   os.Getenv("ENV") == "production",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})

	// Send success response
	sendSuccess(w, map[string]string{
		"message": "Logged out successfully",
	})
}

// HandleMe returns the current user information
func (h *APIAuthHandlers) HandleMe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendError(w, http.StatusMethodNotAllowed, "Method not allowed", "METHOD_NOT_ALLOWED")
		return
	}

	// Get user from context
	user, ok := middleware.GetUserFromContext(r.Context())
	if !ok {
		sendError(w, http.StatusUnauthorized, "Not authenticated", "NOT_AUTHENTICATED")
		return
	}

	// Send success response
	sendSuccess(w, AuthResponse{
		User: userToResponse(user),
	})
}
