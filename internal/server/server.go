package server

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Server wraps the HTTP server and dependencies
type Server struct {
	db     *sql.DB
	router *chi.Mux
}

// New creates a new server instance
func New(db *sql.DB) *Server {
	s := &Server{
		db:     db,
		router: chi.NewRouter(),
	}

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

	// Static files
	s.router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// API routes
	s.router.Route("/api", func(r chi.Router) {
		// API routes will be added here in future milestones
	})

	// Page routes
	s.router.Route("/", func(r chi.Router) {
		// Page routes will be added here in future milestones
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("VuGo - Project Management App\n\nAPI is running. Pages coming soon..."))
		})
	})
}

// ServeHTTP implements http.Handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
