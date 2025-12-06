package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// SPAHandler serves the Single Page Application
type SPAHandler struct {
	staticPath string
	indexPath  string
}

// NewSPAHandler creates a new SPA handler
func NewSPAHandler(staticPath string) *SPAHandler {
	return &SPAHandler{
		staticPath: staticPath,
		indexPath:  filepath.Join(staticPath, "index.html"),
	}
}

// ServeHTTP implements the http.Handler interface
func (h *SPAHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get the absolute path to prevent directory traversal
	path := filepath.Join(h.staticPath, r.URL.Path)

	// Check if the file exists
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		// File does not exist or is a directory, serve index.html for SPA routing
		http.ServeFile(w, r, h.indexPath)
		return
	}

	if err != nil {
		// If there was an error, serve index.html
		log.Printf("Error checking file %s: %v", path, err)
		http.ServeFile(w, r, h.indexPath)
		return
	}

	// Serve the file
	http.ServeFile(w, r, path)
}

// ServeStatic serves static files with proper fallback for SPA
func ServeStatic(staticDir string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip API routes
		if strings.HasPrefix(r.URL.Path, "/api") {
			http.NotFound(w, r)
			return
		}

		// Construct the file path
		filePath := filepath.Join(staticDir, r.URL.Path)

		// Check if file exists
		info, err := os.Stat(filePath)
		if err != nil {
			// File doesn't exist, serve index.html for SPA routing
			indexPath := filepath.Join(staticDir, "index.html")
			http.ServeFile(w, r, indexPath)
			return
		}

		// If it's a directory, try to serve index.html
		if info.IsDir() {
			indexPath := filepath.Join(filePath, "index.html")
			if _, err := os.Stat(indexPath); err == nil {
				http.ServeFile(w, r, indexPath)
				return
			}
			// No index.html in directory, serve main index.html
			mainIndexPath := filepath.Join(staticDir, "index.html")
			http.ServeFile(w, r, mainIndexPath)
			return
		}

		// File exists, serve it
		http.ServeFile(w, r, filePath)
	})
}
