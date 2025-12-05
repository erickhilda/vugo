package server

import (
	"net/http"
)

// setupRoutes is called from server.go
// This file will contain route definitions as we add handlers

// Example route groups (to be implemented in future milestones):
//
// func (s *Server) setupAuthRoutes(r chi.Router) {
//     r.Post("/register", s.handleRegister)
//     r.Post("/login", s.handleLogin)
//     r.Post("/logout", s.handleLogout)
// }
//
// func (s *Server) setupProjectRoutes(r chi.Router) {
//     r.Get("/", s.handleListProjects)
//     r.Post("/", s.handleCreateProject)
//     r.Get("/{id}", s.handleGetProject)
//     r.Put("/{id}", s.handleUpdateProject)
//     r.Delete("/{id}", s.handleDeleteProject)
// }

// handleNotFound handles 404 errors
func (s *Server) handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - Page Not Found"))
}

// handleMethodNotAllowed handles 405 errors
func (s *Server) handleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("405 - Method Not Allowed"))
}
