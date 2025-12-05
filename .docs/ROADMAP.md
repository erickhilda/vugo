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
- Login/Register pages dengan Vue forms

### Files to Create

- `internal/handlers/auth.go` - Auth handlers (register, login, logout)
- `internal/services/auth.go` - Auth business logic
- `internal/middleware/auth.go` - Auth middleware (update dari placeholder)
- `templates/pages/login.html` - Login page dengan Vue form
- `templates/pages/register.html` - Register page dengan Vue form

### Database Queries

- ✅ Users queries (already exists)
- ✅ Sessions queries (already exists)

### Key Features

- Secure password hashing dengan bcrypt
- Cookie-based session management
- Protected routes dengan auth middleware
- Form validation di frontend dan backend
- Error handling dan user feedback

### See Also

Detail implementasi: [M4 Authentication Documentation](milestones/M4-authentication.md)

---

## Milestone 5: Projects CRUD

**Status**: ⏳ Pending

### Deliverables

- Create/Read/Update/Delete projects
- Project list page dengan Vue components
- Project detail page
- Project settings page
- API endpoints untuk Vue reactivity
- Archive/Unarchive functionality

### Files to Create

- `internal/handlers/projects.go` - Project handlers
- `internal/services/projects.go` - Project business logic
- `templates/pages/projects/index.html` - Projects list
- `templates/pages/projects/show.html` - Project detail
- `templates/pages/projects/new.html` - Create project
- `templates/pages/projects/settings.html` - Project settings

### Database Queries

- ✅ Projects queries (already exists)

### Key Features

- Project CRUD dengan ownership validation
- Project color untuk visual identification
- Archive/Unarchive untuk soft delete
- Project member management (basic)
- Activity logging untuk project actions

### See Also

Detail implementasi: [M5 Projects Documentation](milestones/M5-projects.md)

---

## Milestone 6: Boards & Columns

**Status**: ⏳ Pending

### Deliverables

- Board management per project
- Column CRUD operations
- Drag & drop column reordering dengan SortableJS
- WIP limits untuk columns
- Board view page dengan Kanban layout

### Files to Create

- `internal/handlers/boards.go` - Board & column handlers
- `internal/services/boards.go` - Board & column business logic
- `templates/pages/boards/show.html` - Board view
- `templates/components/column.html` - Column component

### Database Queries

- ✅ Boards queries (already exists)
- ✅ Columns queries (already exists)

### Key Features

- Multiple boards per project
- Kanban board layout
- Drag & drop untuk column reordering
- WIP limit dengan visual indicators
- Column color customization

### See Also

Detail implementasi: [M6 Boards Documentation](milestones/M6-boards.md)

---

## Milestone 7: Tasks Core

**Status**: ⏳ Pending

### Deliverables

- Task CRUD operations
- Drag & drop tasks between columns dengan SortableJS
- Task detail modal dengan full information
- Priority levels (low, medium, high, urgent)
- Due dates dengan reminders
- Task assignees (multiple users)

### Files to Create

- `internal/handlers/tasks.go` - Task handlers
- `internal/services/tasks.go` - Task business logic
- `templates/components/task-card.html` - Task card component
- `templates/components/task-modal.html` - Task detail modal

### Database Queries

- ✅ Tasks queries (already exists)
- ✅ Task assignees queries (already exists)

### Key Features

- Full task CRUD dengan validation
- Drag & drop between columns
- Priority color coding
- Due date management
- Multiple assignees support

### See Also

Detail implementasi: [M7 Tasks Documentation](milestones/M7-tasks.md)

---

## Milestone 8: Labels, Checklists & Comments

**Status**: ⏳ Pending

### Deliverables

- Label management per project
- Task labeling (many-to-many)
- Checklist items dengan completion tracking
- Comments thread untuk tasks
- Activity logging untuk all actions

### Files to Create

- `internal/handlers/labels.go` - Label handlers
- `internal/handlers/comments.go` - Comment handlers
- `internal/services/labels.go` - Label business logic
- `internal/services/comments.go` - Comment business logic
- `templates/components/checklist.html` - Checklist component
- `templates/components/comments.html` - Comments component

### Database Queries

- ✅ Labels queries (already exists)
- ✅ Task labels queries (already exists)
- ✅ Checklist items queries (already exists)
- ✅ Comments queries (already exists)
- ✅ Activities queries (already exists)

### Key Features

- Project-level labels dengan color coding
- Checklist dengan progress tracking
- Comments dengan user info
- Activity log untuk audit trail

### See Also

Detail implementasi: [M8 Extras Documentation](milestones/M8-extras.md)

---

## Milestone 9: Dashboard & Search

**Status**: ⏳ Pending

### Deliverables

- User dashboard dengan overview statistics
- My Tasks view (tasks assigned to user)
- Upcoming deadlines dengan calendar view
- Activity feed dengan filtering
- Global search untuk projects, tasks, comments

### Files to Create

- `internal/handlers/dashboard.go` - Dashboard handlers
- `internal/handlers/search.go` - Search handlers
- `internal/services/dashboard.go` - Dashboard business logic
- `internal/services/search.go` - Search business logic
- `templates/pages/dashboard.html` - Dashboard page
- `templates/components/activity-feed.html` - Activity feed component

### Key Features

- Stats overview dengan cards
- My Tasks dengan filters
- Upcoming deadlines dengan highlighting
- Activity feed dengan pagination
- Full-text search across content

### See Also

Detail implementasi: [M9 Dashboard Documentation](milestones/M9-dashboard.md)

---

## Milestone 10: Polish & Production Ready

**Status**: ⏳ Pending

### Deliverables

- Dark mode toggle dengan persistence
- Responsive design untuk mobile/tablet/desktop
- Performance optimization (queries, caching)
- Error handling improvements
- Environment configuration management
- Docker setup untuk deployment
- Documentation updates

### Files to Create/Update

- `config/config.go` - Configuration management
- `Dockerfile` - Docker image
- `docker-compose.yml` - Docker compose
- `.env.example` - Environment variables example
- `docs/DEPLOYMENT.md` - Deployment guide
- Various CSS/template fixes

### Key Features

- Theme switching (dark/light)
- Mobile-first responsive design
- Optimized database queries
- User-friendly error messages
- Production-ready configuration
- Docker deployment

### See Also

Detail implementasi: [M10 Polish Documentation](milestones/M10-polish.md)

---

## Progress Tracking

Setiap milestone akan diupdate statusnya di file ini setelah selesai dikerjakan. Detail implementasi setiap milestone dapat dilihat di folder `.docs/milestones/`.

---

## Notes

- Menggunakan pendekatan **horizontal layer** - setup foundation terlebih dahulu, kemudian fitur-fitur secara bertahap
- Setiap milestone adalah deliverable yang bisa dijalankan dan ditest
- Dokumentasi lengkap setiap milestone ada di `.docs/milestones/MX-name.md`
