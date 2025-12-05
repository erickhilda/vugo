# Milestone 1: Database Foundation

**Status**: ✅ Completed  
**Date Started**: 2025-01-27  
**Date Completed**: 2025-01-27

## Overview

Milestone ini fokus pada setup database foundation dengan SQLite, schema migrations, dan sqlc configuration untuk type-safe SQL code generation.

## Objectives

1. Membuat database schema lengkap untuk semua entitas aplikasi
2. Setup sqlc configuration untuk code generation
3. Membuat migration files untuk database versioning
4. Membuat SQL queries untuk setiap entitas
5. Setup database connection helper

## Deliverables

### ✅ Completed

1. **Database Schema**

   - File `db/migrations` dengan semua tabel:
     - `users` - User accounts
     - `sessions` - Session management
     - `projects` - Projects
     - `project_members` - Project collaboration
     - `boards` - Kanban boards
     - `columns` - Board columns
     - `tasks` - Task cards
     - `task_assignees` - Task assignments
     - `labels` - Project labels
     - `task_labels` - Task labeling
     - `checklist_items` - Task checklists
     - `comments` - Task comments
     - `activities` - Activity log
   - Semua tabel memiliki indexes yang diperlukan untuk performa

2. **Migrations**

   - `db/migrations/000001_initial_schema.up.sql` - Create all tables
   - `db/migrations/000001_initial_schema.down.sql` - Drop all tables
   - Menggunakan golang-migrate format

3. **SQL Queries (sqlc)**

   - `db/queries/users.sql` - User CRUD operations
   - `db/queries/sessions.sql` - Session management
   - `db/queries/projects.sql` - Project operations
   - `db/queries/project_members.sql` - Project membership
   - `db/queries/boards.sql` - Board operations
   - `db/queries/columns.sql` - Column operations
   - `db/queries/tasks.sql` - Task operations
   - `db/queries/task_assignees.sql` - Task assignments
   - `db/queries/labels.sql` - Label management
   - `db/queries/task_labels.sql` - Task labeling
   - `db/queries/checklist_items.sql` - Checklist operations
   - `db/queries/comments.sql` - Comment operations
   - `db/queries/activities.sql` - Activity logging

4. **sqlc Configuration**

   - `sqlc.yaml` dengan konfigurasi untuk SQLite
   - Output directory: `internal/database/queries`
   - Package name: `sqlc`

5. **Database Connection**
   - `internal/database/db.go` - Database connection helper
   - Support untuk SQLite dengan foreign keys enabled
   - Connection pool configuration
   - WAL mode untuk better concurrency

## Files Created

```
vugo/
├── sqlc.yaml
├── db/
│   ├── migrations/
│   │   ├── 000001_initial_schema.up.sql
│   │   └── 000001_initial_schema.down.sql
│   └── queries/
│       ├── users.sql
│       ├── sessions.sql
│       ├── projects.sql
│       ├── project_members.sql
│       ├── boards.sql
│       ├── columns.sql
│       ├── tasks.sql
│       ├── task_assignees.sql
│       ├── labels.sql
│       ├── task_labels.sql
│       ├── checklist_items.sql
│       ├── comments.sql
│       └── activities.sql
└── internal/
    └── database/
        └── db.go
```

## Technical Decisions

1. **SQLite**: Dipilih karena kemudahan development dan deployment. Support untuk foreign keys dan WAL mode untuk concurrency.

2. **sqlc**: Menggunakan sqlc untuk type-safe SQL code generation. Tidak menggunakan ORM untuk lebih control dan performa.

3. **Migrations**: Menggunakan golang-migrate untuk versioning database schema. Format: `000001_initial_schema.up.sql` dan `.down.sql`.

4. **Indexes**: Semua foreign keys dan frequently queried columns sudah di-index untuk performa optimal.

5. **Naming Convention**:
   - Tables: plural (users, projects, tasks)
   - Columns: snake_case (user_id, created_at)
   - Queries: descriptive names (GetUser, ListProjectsByOwner)

## Next Steps

Setelah milestone ini selesai, langkah selanjutnya adalah:

1. **Generate sqlc code**: Run `make sqlc` untuk generate Go code dari SQL queries
2. **Milestone 2: HTTP Server & Routing** - Setup Chi router dan middleware

## Notes

- Semua queries menggunakan prepared statements untuk security
- Foreign keys enabled untuk data integrity
- WAL mode untuk better concurrency dengan SQLite
- Activity log menggunakan JSON untuk details field (flexible untuk berbagai action types)
- Position fields untuk drag-and-drop ordering

## Testing

Untuk test database setup:

```bash
# Generate sqlc code
make sqlc

# Run migrations
make migrate-up

# Verify database
sqlite3 data/vugo.db ".tables"
```

## References

- [sqlc Documentation](https://docs.sqlc.dev/)
- [golang-migrate Documentation](https://github.com/golang-migrate/migrate)
- [SQLite Documentation](https://www.sqlite.org/docs.html)
