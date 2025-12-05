package main

import (
	"log"
	"net/http"
	"os"

	"github.com/erickhilda/vugo/internal/database"
	"github.com/erickhilda/vugo/internal/server"
)

func main() {
	// Get database path from environment or use default
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "data/vugo.db"
	}

	// Initialize database
	db, err := database.New(dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Get server port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize server
	srv := server.New(db.DB)

	// Start server
	log.Printf("Server starting on :%s", port)
	if err := http.ListenAndServe(":"+port, srv); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
