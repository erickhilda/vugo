package services

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"time"

	"github.com/erickhilda/vugo/internal/database/queries"
	"golang.org/x/crypto/bcrypt"
)

// AuthService handles authentication business logic
type AuthService struct {
	queries *queries.Queries
}

// NewAuthService creates a new auth service
func NewAuthService(q *queries.Queries) *AuthService {
	return &AuthService{
		queries: q,
	}
}

// RegisterResult contains the result of registration
type RegisterResult struct {
	User    queries.User
	Session queries.Session
}

// Register creates a new user and session
func (s *AuthService) Register(ctx context.Context, email, password, name string) (*RegisterResult, error) {
	// Validate input
	if len(name) < 2 || len(name) > 100 {
		return nil, errors.New("name must be between 2 and 100 characters")
	}

	if len(password) < 8 {
		return nil, errors.New("password must be at least 8 characters")
	}

	// Check if email already exists
	_, err := s.queries.GetUserByEmail(ctx, email)
	if err == nil {
		return nil, errors.New("email already registered")
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user, err := s.queries.CreateUser(ctx, queries.CreateUserParams{
		Email:        email,
		PasswordHash: string(hashedPassword),
		Name:         name,
		AvatarUrl:    sql.NullString{Valid: false},
	})
	if err != nil {
		return nil, err
	}

	// Create session
	sessionID, err := s.GenerateSessionID()
	if err != nil {
		return nil, err
	}

	session, err := s.queries.CreateSession(ctx, queries.CreateSessionParams{
		ID:        sessionID,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour), // 7 days
	})
	if err != nil {
		return nil, err
	}

	return &RegisterResult{
		User:    user,
		Session: session,
	}, nil
}

// LoginResult contains the result of login
type LoginResult struct {
	User    queries.User
	Session queries.Session
}

// Login authenticates a user and creates a session
func (s *AuthService) Login(ctx context.Context, email, password string) (*LoginResult, error) {
	// Get user by email
	user, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Create session
	sessionID, err := s.GenerateSessionID()
	if err != nil {
		return nil, err
	}

	session, err := s.queries.CreateSession(ctx, queries.CreateSessionParams{
		ID:        sessionID,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour), // 7 days
	})
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		User:    user,
		Session: session,
	}, nil
}

// Logout deletes a session
func (s *AuthService) Logout(ctx context.Context, sessionID string) error {
	return s.queries.DeleteSession(ctx, sessionID)
}

// GetUserBySession validates a session and returns the user
func (s *AuthService) GetUserBySession(ctx context.Context, sessionID string) (*queries.User, error) {
	// Get session
	session, err := s.queries.GetSession(ctx, sessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("session not found")
		}
		return nil, err
	}

	// Check if session is expired
	if time.Now().After(session.ExpiresAt) {
		// Delete expired session
		_ = s.queries.DeleteSession(ctx, sessionID)
		return nil, errors.New("session expired")
	}

	// Get user
	user, err := s.queries.GetUser(ctx, session.UserID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GenerateSessionID generates a secure random session ID
func (s *AuthService) GenerateSessionID() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// CleanupExpiredSessions deletes expired sessions
func (s *AuthService) CleanupExpiredSessions(ctx context.Context) error {
	return s.queries.DeleteExpiredSessions(ctx)
}
