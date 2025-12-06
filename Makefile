.PHONY: help dev build run test clean install migrate sqlc frontend-install frontend-dev frontend-build frontend-clean build-all

# Default target
help:
	@echo "VuGo - Project Management App"
	@echo ""
	@echo "Backend commands:"
	@echo "  make dev          - Run backend development server with hot reload (requires air)"
	@echo "  make build        - Build the backend application"
	@echo "  make run          - Run the backend application"
	@echo "  make test         - Run backend tests"
	@echo "  make clean        - Clean backend build artifacts"
	@echo "  make install      - Install backend dependencies"
	@echo ""
	@echo "Frontend commands:"
	@echo "  make frontend-install - Install frontend dependencies (requires pnpm)"
	@echo "  make frontend-dev     - Run frontend dev server"
	@echo "  make frontend-build   - Build frontend for production"
	@echo "  make frontend-clean   - Clean frontend build artifacts"
	@echo ""
	@echo "Full-stack commands:"
	@echo "  make build-all    - Build both backend and frontend"
	@echo ""
	@echo "Database commands:"
	@echo "  make migrate-up   - Run database migrations up"
	@echo "  make migrate-down - Run database migrations down"
	@echo "  make migrate-make - Create a new migration file"
	@echo ""
	@echo "Code generation:"
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

# Frontend commands
frontend-install:
	@echo "Installing frontend dependencies..."
	@cd frontend && pnpm install
	@echo "Frontend dependencies installed"

frontend-dev:
	@echo "Starting frontend dev server..."
	@cd frontend && pnpm dev

frontend-build:
	@echo "Building frontend for production..."
	@cd frontend && pnpm build
	@echo "Frontend build complete: frontend/dist/"

frontend-clean:
	@echo "Cleaning frontend build artifacts..."
	@rm -rf frontend/dist
	@rm -rf frontend/node_modules/.vite
	@echo "Frontend clean complete"

# Build both backend and frontend
build-all: frontend-build build
	@echo "Full build complete!"
	@echo "  - Backend: bin/vugo"
	@echo "  - Frontend: frontend/dist/"

