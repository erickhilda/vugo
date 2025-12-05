.PHONY: help dev build run test clean install migrate sqlc

# Default target
help:
	@echo "VuGo - Project Management App"
	@echo ""
	@echo "Available commands:"
	@echo "  make dev          - Run development server with hot reload (requires air)"
	@echo "  make build        - Build the application"
	@echo "  make run          - Run the application"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make install      - Install dependencies"
	@echo "  make migrate-up   - Run database migrations up"
	@echo "  make migrate-down - Run database migrations down"
	@echo "  make migrate-make   - Create a new migration file"
	@echo "  make sqlc         - Generate sqlc code from queries"
	@echo "  make sqlc-validate - Validate sqlc queries"

# Development server with hot reload
dev:
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "Air not installed. Install with: go install github.com/air-verse/air@latest"; \
		go run cmd/server/main.go; \
	fi

# Build the application
build:
	@echo "Building VuGo..."
	@go build -o bin/vugo cmd/server/main.go
	@echo "Build complete: bin/vugo"

# Run the application
run: build
	@./bin/vugo

# Run tests
test:
	@go test -v ./...

# Clean build artifacts
clean:
	@rm -rf bin/
	@rm -f vugo
	@echo "Clean complete"

# Install dependencies
install:
	@echo "Installing Go dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies installed"

# Database migrations
migrate-up:
	@mkdir -p data
	@if command -v migrate > /dev/null; then \
		migrate -path db/migrations -database "sqlite3://data/vugo.db" up; \
	else \
		echo "golang-migrate not installed. Install with: go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; \
	fi

migrate-down:
	@mkdir -p data
	@if command -v migrate > /dev/null; then \
		migrate -path db/migrations -database "sqlite3://data/vugo.db" down; \
	else \
		echo "golang-migrate not installed."; \
	fi

migrate-make:
	@echo "Creating new migration file..."
	@migrate create -ext sql -dir db/migrations -seq $(name)

# SQLC code generation
sqlc:
	@if command -v sqlc > /dev/null; then \
		sqlc generate; \
	else \
		echo "sqlc not installed. Install with: go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest"; \
	fi

sqlc-validate:
	@if command -v sqlc > /dev/null; then \
		sqlc vet; \
	else \
		echo "sqlc not installed."; \
	fi

