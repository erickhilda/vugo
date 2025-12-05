package handlers

import (
	"bytes"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/erickhilda/vugo/internal/database/queries"
	"github.com/erickhilda/vugo/internal/middleware"
	"github.com/erickhilda/vugo/internal/services"
)

// AuthHandlers handles authentication routes
type AuthHandlers struct {
	authService *services.AuthService
	templates   *template.Template
}

// NewAuthHandlers creates a new auth handlers instance
func NewAuthHandlers(authService *services.AuthService, templatesFS fs.FS) *AuthHandlers {
	// Parse templates
	tmpl := template.New("").Funcs(templateFuncs())

	// Walk through templates directory
	err := fs.WalkDir(templatesFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Only process .html files
		if d.IsDir() || filepath.Ext(path) != ".html" {
			return nil
		}

		// Read template file
		data, err := fs.ReadFile(templatesFS, path)
		if err != nil {
			return err
		}

		// Parse template
		name := path
		if _, err := tmpl.New(name).Parse(string(data)); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Printf("Warning: failed to parse templates: %v", err)
	} else {
		log.Printf("Templates loaded successfully")
		// Log all template names for debugging
		for _, tmpl := range tmpl.Templates() {
			log.Printf("  - Template: %s", tmpl.Name())
		}
	}

	return &AuthHandlers{
		authService: authService,
		templates:   tmpl,
	}
}

// templateFuncs returns template helper functions
func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"sub": func(a, b int) int {
			return a - b
		},
		"eq": func(a, b interface{}) bool {
			return a == b
		},
		"ne": func(a, b interface{}) bool {
			return a != b
		},
		"gt": func(a, b int) bool {
			return a > b
		},
		"lt": func(a, b int) bool {
			return a < b
		},
	}
}

// TemplateData holds data for template rendering
type TemplateData struct {
	Title string
	User  *queries.User
	Error string
}

// RenderTemplate renders a template with data
func (h *AuthHandlers) RenderTemplate(w http.ResponseWriter, name string, data TemplateData) error {
	var buf bytes.Buffer
	if err := h.templates.ExecuteTemplate(&buf, name, data); err != nil {
		log.Printf("Template execution error for '%s': %v", name, err)
		// Log available templates for debugging
		log.Printf("Available templates:")
		for _, tmpl := range h.templates.Templates() {
			log.Printf("  - %s", tmpl.Name())
		}
		return err
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := buf.WriteTo(w)
	return err
}

// ShowLoginPage renders the login page
func (h *AuthHandlers) ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	// Check if user is already logged in
	_, ok := middleware.GetUserFromContext(r.Context())
	if ok {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	data := TemplateData{
		Title: "Login",
		User:  nil,
	}

	if err := h.RenderTemplate(w, "templates/pages/login.html", data); err != nil {
		log.Printf("Error rendering login page: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// ShowRegisterPage renders the register page
func (h *AuthHandlers) ShowRegisterPage(w http.ResponseWriter, r *http.Request) {
	// Check if user is already logged in
	_, ok := middleware.GetUserFromContext(r.Context())
	if ok {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	data := TemplateData{
		Title: "Register",
		User:  nil,
	}

	if err := h.RenderTemplate(w, "templates/pages/register.html", data); err != nil {
		log.Printf("Error rendering register page: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// HandleRegister processes user registration
func (h *AuthHandlers) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	// Validate
	if password != confirmPassword {
		data := TemplateData{
			Title: "Register",
			Error: "Passwords do not match",
		}
		h.RenderTemplate(w, "templates/pages/register.html", data)
		return
	}

	// Register user
	result, err := h.authService.Register(r.Context(), email, password, name)
	if err != nil {
		data := TemplateData{
			Title: "Register",
			Error: err.Error(),
		}
		h.RenderTemplate(w, "templates/pages/register.html", data)
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

	// Redirect to dashboard
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

// HandleLogin processes user login
func (h *AuthHandlers) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	// Login user
	result, err := h.authService.Login(r.Context(), email, password)
	if err != nil {
		data := TemplateData{
			Title: "Login",
			Error: err.Error(),
		}
		h.RenderTemplate(w, "templates/pages/login.html", data)
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

	// Redirect to dashboard
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

// HandleLogout processes user logout
func (h *AuthHandlers) HandleLogout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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

	// Redirect to home
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
