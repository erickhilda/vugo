# VuGo Development Roadmap

## Project Overview

**VuGo** adalah aplikasi Project & Task Management (Kanban board) yang dibangun dengan:

- **Backend**: Go + Chi + sqlc + SQLite
- **Frontend**: Vue 3 (CDN) + Tailwind CSS + SortableJS

## Milestone Status

| Milestone                                                                      | Status       | Description                                                           |
| ------------------------------------------------------------------------------ | ------------ | --------------------------------------------------------------------- |
| [M0: Project Setup & Documentation](#milestone-0-project-setup--documentation) | ✅ Completed | Setup struktur folder, dependency management, dan dokumentasi roadmap |
| [M1: Database Foundation](#milestone-1-database-foundation)                    | ✅ Completed | Setup database schema, migrations, dan sqlc configuration             |
| [M2: HTTP Server & Routing](#milestone-2-http-server--routing)                 | ✅ Completed | Setup HTTP server, middleware, dan route structure                    |
| [M3: Template Engine & Base Layout](#milestone-3-template-engine--base-layout) | ✅ Completed | Setup Go templates dengan Vue integration                             |
| [M4: Authentication System](#milestone-4-authentication-system)                | ⏳ Pending   | Implementasi user registration, login, dan session management         |
| [M5: Projects CRUD](#milestone-5-projects-crud)                                | ⏳ Pending   | Implementasi manajemen projects                                       |
| [M6: Boards & Columns](#milestone-6-boards--columns)                           | ⏳ Pending   | Implementasi Kanban boards dan columns                                |
| [M7: Tasks Core](#milestone-7-tasks-core)                                      | ⏳ Pending   | Implementasi task cards dengan full functionality                     |
| [M8: Labels, Checklists & Comments](#milestone-8-labels-checklists--comments)  | ⏳ Pending   | Fitur tambahan untuk tasks                                            |
| [M9: Dashboard & Search](#milestone-9-dashboard--search)                       | ⏳ Pending   | User dashboard dan search functionality                               |
| [M10: Polish & Production Ready](#milestone-10-polish--production-ready)       | ⏳ Pending   | Final polish dan production readiness                                 |

---

## Milestone 0: Project Setup & Documentation

**Status**: ✅ Completed  
**Date**: 2025-01-27

### Deliverables

- ✅ Struktur direktori proyek
- ✅ `go.mod` dengan dependencies
- ✅ Makefile untuk common tasks
- ✅ Folder `.docs/` untuk dokumentasi roadmap
- ✅ README.md dengan instruksi setup

### Files Created

- `go.mod`
- `Makefile`
- `README.md`
- `.docs/ROADMAP.md`
- `.docs/CHANGELOG.md`
- `.docs/milestones/M0-setup.md`

### Notes

Project structure telah dibuat dengan pendekatan horizontal layer. Semua dependencies dasar sudah ditambahkan ke `go.mod`. Makefile menyediakan commands untuk development workflow.

---

## Milestone 1: Database Foundation

**Status**: ✅ Completed  
**Date**: 2025-01-27

### Deliverables

- SQLite database setup
- Schema migrations (semua tabel)
- sqlc configuration dan generated code
- Database connection helper

### Files to Create

- `sqlc.yaml`
- `db/schema.sql`
- `db/migrations/*.sql`
- `db/queries/*.sql` (per entity)
- `internal/database/db.go`
- Generated: `internal/database/sqlc/`

### Database Schema

Akan mencakup tabel-tabel untuk:

- Users & Sessions
- Projects & Project Members
- Boards & Columns
- Tasks & Task Assignees
- Labels & Task Labels
- Checklist Items
- Comments
- Activity Log

---

## Milestone 2: HTTP Server & Routing

**Status**: ✅ Completed  
**Date**: 2025-01-27

### Deliverables

- ✅ Chi router setup
- ✅ Middleware (logging, recovery, CORS)
- ✅ Route grouping (pages vs API)
- ✅ Static file serving
- ✅ Health check endpoint

### Files Created

- `cmd/server/main.go`
- `internal/server/server.go`
- `internal/server/routes.go`
- `internal/middleware/auth.go`

### Notes

Server sudah setup dengan Chi router, middleware untuk logging, recovery, dan CORS. Health check endpoint sudah tersedia di `/health`. Route structure sudah disiapkan untuk API dan pages yang akan ditambahkan di milestone berikutnya.

### Deliverables

- Chi router setup
- Middleware (logging, recovery, CORS)
- Route grouping (pages vs API)
- Static file serving
- Health check endpoint

### Files to Create

- `cmd/server/main.go`
- `internal/server/server.go`
- `internal/server/routes.go`
- `internal/middleware/*.go`

---

## Milestone 3: Template Engine & Base Layout

**Status**: ✅ Completed  
**Date**: 2025-01-27

### Deliverables

- ✅ Template parsing dan rendering
- ✅ Base layout dengan Vue CDN
- ✅ Template helpers/functions
- ✅ Error pages (404, 500)
- ✅ Landing page (static)

### Files Created

- `internal/templates/templates.go`
- `templates/layouts/base.html`
- `templates/layouts/app.html`
- `templates/pages/landing.html`
- `templates/errors/404.html`
- `templates/errors/500.html`
- `static/css/app.css`

### Notes

Template engine sudah setup dengan support untuk Vue 3 (CDN), Tailwind CSS, VueUse, Axios, dan SortableJS. Base layout dan app layout sudah dibuat. Landing page dan error pages sudah tersedia. Custom CSS untuk task cards, drag-and-drop, dan priority colors sudah ditambahkan.

### Deliverables

- Template parsing dan rendering
- Base layout dengan Vue CDN
- Template helpers/functions
- Error pages (404, 500)
- Landing page (static)

### Files to Create

- `internal/templates/templates.go`
- `templates/layouts/base.html`
- `templates/layouts/app.html`
- `templates/pages/landing.html`
- `templates/errors/404.html`
- `templates/errors/500.html`
- `static/css/app.css`

---

## Milestone 4: Authentication System

**Status**: ⏳ Pending

### Deliverables

- User registration dengan validation
- Login/Logout flow
- Session management (secure cookies)
- Password hashing (bcrypt)
- Auth middleware (protected routes)
- Login/Register pages

### Files to Create

- `internal/handlers/auth.go`
- `internal/services/auth.go`
- `internal/middleware/auth.go`
- `db/queries/users.sql`
- `db/queries/sessions.sql`
- `templates/pages/login.html`
- `templates/pages/register.html`

---

## Milestone 5: Projects CRUD

**Status**: ⏳ Pending

### Deliverables

- Create/Read/Update/Delete projects
- Project list page
- Project detail page
- Project settings
- API endpoints untuk Vue reactivity

### Files to Create

- `internal/handlers/projects.go`
- `internal/services/projects.go`
- `db/queries/projects.sql`
- `templates/pages/projects/index.html`
- `templates/pages/projects/show.html`
- `templates/pages/projects/settings.html`

---

## Milestone 6: Boards & Columns

**Status**: ⏳ Pending

### Deliverables

- Board management per project
- Column CRUD
- Drag & drop column reordering
- WIP limits
- Board view page

### Files to Create

- `internal/handlers/boards.go`
- `internal/services/boards.go`
- `db/queries/boards.sql`
- `db/queries/columns.sql`
- `templates/pages/boards/show.html`
- `templates/components/column.html`

---

## Milestone 7: Tasks Core

**Status**: ⏳ Pending

### Deliverables

- Task CRUD
- Drag & drop tasks between columns
- Task detail modal
- Priority & due dates
- Task assignees

### Files to Create

- `internal/handlers/tasks.go`
- `internal/services/tasks.go`
- `db/queries/tasks.sql`
- `db/queries/task_assignees.sql`
- `templates/components/task-card.html`
- `templates/components/task-modal.html`

---

## Milestone 8: Labels, Checklists & Comments

**Status**: ⏳ Pending

### Deliverables

- Label management
- Task labels
- Checklist items
- Comments thread
- Activity logging

### Files to Create

- `internal/handlers/labels.go`
- `internal/handlers/comments.go`
- `db/queries/labels.sql`
- `db/queries/checklist_items.sql`
- `db/queries/comments.sql`
- `db/queries/activities.sql`
- `templates/components/checklist.html`
- `templates/components/comments.html`

---

## Milestone 9: Dashboard & Search

**Status**: ⏳ Pending

### Deliverables

- Dashboard dengan overview
- My tasks view
- Upcoming deadlines
- Activity feed
- Global search

### Files to Create

- `internal/handlers/dashboard.go`
- `internal/handlers/search.go`
- `templates/pages/dashboard.html`
- `templates/components/activity-feed.html`

---

## Milestone 10: Polish & Production Ready

**Status**: ⏳ Pending

### Deliverables

- Dark mode toggle
- Responsive design fixes
- Performance optimization
- Error handling improvements
- Environment configuration
- Docker setup (optional)

### Files to Create/Update

- `config/config.go`
- `Dockerfile`
- `docker-compose.yml`
- Various CSS/template fixes

---

## Progress Tracking

Setiap milestone akan diupdate statusnya di file ini setelah selesai dikerjakan. Detail implementasi setiap milestone dapat dilihat di folder `.docs/milestones/`.

---

## Notes

- Menggunakan pendekatan **horizontal layer** - setup foundation terlebih dahulu, kemudian fitur-fitur secara bertahap
- Setiap milestone adalah deliverable yang bisa dijalankan dan ditest
- Dokumentasi lengkap setiap milestone ada di `.docs/milestones/MX-name.md`
