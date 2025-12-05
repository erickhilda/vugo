package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/erickhilda/vugo/internal/database/queries"
	"github.com/erickhilda/vugo/internal/handlers"
	authMiddleware "github.com/erickhilda/vugo/internal/middleware"
	"github.com/erickhilda/vugo/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Server wraps the HTTP server and dependencies
type Server struct {
	db           *sql.DB
	router       *chi.Mux
	queries      *queries.Queries
	authService  *services.AuthService
	authHandlers *handlers.AuthHandlers
	authMW       *authMiddleware.AuthMiddleware
}

// New creates a new server instance
func New(db *sql.DB) *Server {
	s := &Server{
		db:     db,
		router: chi.NewRouter(),
	}

	// Initialize queries
	s.queries = queries.New(db)

	// Initialize services
	s.authService = services.NewAuthService(s.queries)

	// Initialize middleware
	s.authMW = authMiddleware.NewAuthMiddleware(s.authService)

	// Initialize handlers with templates
	templatesFS := os.DirFS("templates")
	s.authHandlers = handlers.NewAuthHandlers(s.authService, templatesFS)

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

	// Auth routes (public)
	s.router.Route("/", func(r chi.Router) {
		// Optional auth middleware to load user if logged in
		r.Use(s.authMW.OptionalAuth)

		// Landing page
		r.Get("/", s.handleLanding)

		// Auth pages
		r.Get("/login", s.authHandlers.ShowLoginPage)
		r.Post("/login", s.authHandlers.HandleLogin)
		r.Get("/register", s.authHandlers.ShowRegisterPage)
		r.Post("/register", s.authHandlers.HandleRegister)

		// Protected routes
		r.Group(func(r chi.Router) {
			r.Use(s.authMW.RequireAuth)
			r.Get("/dashboard", s.handleDashboard)
			r.Post("/logout", s.authHandlers.HandleLogout)
		})
	})

	// API routes
	s.router.Route("/api", func(r chi.Router) {
		// API routes will be added here in future milestones
	})
}

// handleLanding renders the landing page
func (s *Server) handleLanding(w http.ResponseWriter, r *http.Request) {
	data := handlers.TemplateData{
		Title: "Home",
		User:  nil,
	}

	// Get user from context if available
	if user, ok := authMiddleware.GetUserFromContext(r.Context()); ok {
		data.User = user
	}

	if err := s.authHandlers.RenderTemplate(w, "templates/pages/landing.html", data); err != nil {
		log.Printf("Error rendering landing page: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// handleDashboard renders the dashboard page
func (s *Server) handleDashboard(w http.ResponseWriter, r *http.Request) {
	user, _ := authMiddleware.GetUserFromContext(r.Context())

	data := handlers.TemplateData{
		Title: "Dashboard",
		User:  user,
	}

	if err := s.authHandlers.RenderTemplate(w, "templates/pages/dashboard.html", data); err != nil {
		log.Printf("Error rendering dashboard page: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// ServeHTTP implements http.Handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
