# VuGo

A modern Project & Task Management application built with Go and Vue.js.

## Overview

VuGo is a Kanban-style project management tool that combines the power of Go backend with Vue.js frontend (CDN approach). It provides a simple yet powerful way to manage projects, tasks, and collaborate with your team.

## Tech Stack

### Backend

- **Go 1.22+** - Programming language
- **Chi** - HTTP router
- **sqlc** - Type-safe SQL code generation
- **SQLite** - Database
- **golang-migrate** - Database migrations

### Frontend

- **Vue 3** (CDN) - Progressive JavaScript framework
- **Tailwind CSS** (CDN) - Utility-first CSS framework
- **SortableJS** - Drag and drop functionality
- **VueUse** (CDN) - Vue composition utilities

## Project Structure

```
vugo/
├── .docs/              # Documentation and roadmap
├── cmd/
│   └── server/        # Application entry point
├── config/            # Configuration management
├── db/
│   ├── migrations/    # Database migrations
│   └── queries/       # SQL queries for sqlc
├── internal/
│   ├── database/      # Database connection & sqlc generated code
│   ├── handlers/      # HTTP handlers
│   ├── middleware/    # HTTP middleware
│   ├── server/        # Server setup
│   ├── services/      # Business logic
│   └── templates/     # Template helpers
├── static/            # Static assets (CSS, JS)
├── templates/         # Go HTML templates
├── go.mod
├── Makefile
└── README.md
```

## Prerequisites

- Go 1.22 or higher
- Make (optional, for using Makefile commands)
- sqlc (for code generation)
- golang-migrate (for database migrations)

## Installation

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd vugo
   ```

2. **Install dependencies**

   ```bash
   make install
   # or manually:
   go mod download
   ```

3. **Install tools (optional but recommended)**

   ```bash
   # Install sqlc
   go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

   # Install golang-migrate
   go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

   # Install Air for hot reload (optional)
   go install github.com/air-verse/air@latest
   ```

4. **Run database migrations**

   ```bash
   make migrate-up
   ```

5. **Generate sqlc code**
   ```bash
   make sqlc
   ```

## Development

### Running the Server

```bash
# Development mode with hot reload (requires Air)
make dev

# Or run directly
go run cmd/server/main.go

# Or build and run
make build
make run
```

The server will start on `http://localhost:8080` by default.

### Available Commands

- `make dev` - Run development server with hot reload
- `make build` - Build the application
- `make run` - Run the built application
- `make test` - Run tests
- `make clean` - Clean build artifacts
- `make install` - Install Go dependencies
- `make migrate-up` - Run database migrations
- `make migrate-down` - Rollback database migrations
- `make sqlc` - Generate sqlc code from queries
- `make sqlc-validate` - Validate sqlc queries

## Roadmap

See [.docs/ROADMAP.md](.docs/ROADMAP.md) for the complete development roadmap and milestones.

## Contributing

This is a personal project, but contributions and suggestions are welcome!

## License

MIT License
