# Changelog

All notable changes to VuGo will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Initial project structure
- Go module setup with dependencies (Chi, sqlc, SQLite, golang-migrate)
- Makefile with common development commands
- Documentation structure (ROADMAP.md, CHANGELOG.md)
- README.md with setup instructions
- Database schema with all tables (users, projects, boards, tasks, etc.)
- SQL migrations using golang-migrate
- sqlc configuration and queries for all entities
- Database connection helper
- HTTP server with Chi router
- Middleware (logging, recovery, CORS)
- Template engine with Vue 3 integration
- Base layouts (base.html, app.html)
- Landing page and error pages (404, 500)
- Custom CSS styles

### Changed

### Deprecated

### Removed

### Fixed

### Security

---

## [0.1.0] - 2025-01-27

### Added

- Milestone 0: Project Setup & Documentation
  - Project directory structure
  - `go.mod` with core dependencies
  - `Makefile` for development workflow
  - Documentation folder structure
  - README with installation and usage instructions
