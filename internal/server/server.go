package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/erickhilda/vugo/internal/database/queries"
	"github.com/erickhilda/vugo/internal/handlers"
	"github.com/erickhilda/vugo/internal/handlers/api"
	authMiddleware "github.com/erickhilda/vugo/internal/middleware"
	"github.com/erickhilda/vugo/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Server wraps the HTTP server and dependencies
type Server struct {
	db              *sql.DB
	router          *chi.Mux
	authService     *services.AuthService
	apiAuthHandlers *api.APIAuthHandlers
	authMW          *authMiddleware.AuthMiddleware
}

// New creates a new server instance
func New(db *sql.DB) *Server {
	s := &Server{
		db:     db,
		router: chi.NewRouter(),
	}

	// Initialize services
	queries := queries.New(db)
	s.authService = services.NewAuthService(queries)

	// Initialize middleware
	s.authMW = authMiddleware.NewAuthMiddleware(s.authService)

	// Initialize API handlers
	s.apiAuthHandlers = api.NewAPIAuthHandlers(s.authService)

	s.setupMiddleware()
	s.setupRoutes()

	return s
}

// setupMiddleware configures middleware
func (s *Server) setupMiddleware() {
	// Request ID for tracing
	s.router.Use(middleware.RequestID)

	// Logger
	s.router.Use(middleware.Logger)

	// Recoverer for panic recovery
	s.router.Use(middleware.Recoverer)

	// CORS
	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
}

// setupRoutes configures all routes
func (s *Server) setupRoutes() {
	// Health check
	s.router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// API routes
	s.router.Route("/api", func(r chi.Router) {
		// Auth API routes (public)
		r.Post("/auth/login", s.apiAuthHandlers.HandleLogin)
		r.Post("/auth/register", s.apiAuthHandlers.HandleRegister)

		// Protected API routes
		r.Group(func(r chi.Router) {
			r.Use(s.authMW.RequireAuth)
			r.Get("/auth/me", s.apiAuthHandlers.HandleMe)
			r.Post("/auth/logout", s.apiAuthHandlers.HandleLogout)
		})
	})

	// SPA - Serve frontend static files (catch-all route, must be last)
	// Check if frontend build exists
	if _, err := os.Stat("frontend/dist/index.html"); err == nil {
		// Production mode - serve SPA
		log.Println("Serving SPA from frontend/dist")
		s.router.Handle("/*", handlers.ServeStatic("frontend/dist"))
	} else {
		// Development mode - show message
		log.Println("Frontend build not found. Run 'make frontend-build' or use 'make frontend-dev' for development")
		s.router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
				<!DOCTYPE html>
				<html>
				<head>
					<title>Vugo - Development Mode</title>
					<style>
						body {
							font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
							display: flex;
							justify-content: center;
							align-items: center;
							height: 100vh;
							margin: 0;
							background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
							color: white;
						}
						.container {
							text-align: center;
							padding: 2rem;
							background: rgba(255, 255, 255, 0.1);
							backdrop-filter: blur(10px);
							border-radius: 1rem;
							max-width: 600px;
						}
						h1 { margin-bottom: 1rem; }
						p { margin: 0.5rem 0; line-height: 1.6; }
						code {
							background: rgba(0, 0, 0, 0.2);
							padding: 0.2rem 0.5rem;
							border-radius: 0.25rem;
							font-family: 'Courier New', monospace;
						}
						.instructions {
							margin-top: 2rem;
							text-align: left;
							background: rgba(0, 0, 0, 0.2);
							padding: 1rem;
							border-radius: 0.5rem;
						}
						.instructions ol {
							margin: 0.5rem 0;
							padding-left: 1.5rem;
						}
						.instructions li {
							margin: 0.5rem 0;
						}
					</style>
				</head>
				<body>
					<div class="container">
						<h1>🚀 Vugo - Development Mode</h1>
						<p>Frontend build not found. The SPA frontend needs to be built or run in development mode.</p>
						
						<div class="instructions">
							<strong>For Development:</strong>
							<ol>
								<li>Terminal 1: <code>make dev</code> (Backend - runs on :8080)</li>
								<li>Terminal 2: <code>make frontend-dev</code> (Frontend - runs on :5173)</li>
								<li>Access: <code>http://localhost:5173</code></li>
							</ol>
							
							<strong>For Production:</strong>
							<ol>
								<li>Build: <code>make build-all</code></li>
								<li>Run: <code>./bin/vugo</code></li>
								<li>Access: <code>http://localhost:8080</code></li>
							</ol>
						</div>
						
						<p style="margin-top: 2rem; opacity: 0.8; font-size: 0.9rem;">
							API is accessible at <code>/api/*</code> endpoints
						</p>
					</div>
				</body>
				</html>
			`))
		})
	}
}

// ServeHTTP implements http.Handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
