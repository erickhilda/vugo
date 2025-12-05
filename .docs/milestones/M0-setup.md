# Milestone 0: Project Setup & Documentation

**Status**: ✅ Completed  
**Date Started**: 2025-01-27  
**Date Completed**: 2025-01-27

## Overview

Milestone ini fokus pada setup awal project structure, dependency management, dan dokumentasi roadmap. Ini adalah foundation yang diperlukan sebelum mulai development fitur-fitur aplikasi.

## Objectives

1. Membuat struktur direktori yang sesuai dengan best practices Go project
2. Setup Go module dengan dependencies yang diperlukan
3. Membuat Makefile untuk memudahkan development workflow
4. Membuat dokumentasi roadmap dan changelog
5. Membuat README dengan instruksi setup yang jelas

## Deliverables

### ✅ Completed

1. **Project Structure**

   - Direktori utama sudah dibuat sesuai dengan struktur yang direncanakan
   - Folder `.docs/` untuk dokumentasi
   - Folder `cmd/`, `internal/`, `db/`, `templates/`, `static/` sudah direncanakan

2. **Go Module Setup**

   - `go.mod` dibuat dengan module name `github.com/zou/vugo`
   - Dependencies yang ditambahkan:
     - `github.com/go-chi/chi/v5` - HTTP router
     - `github.com/go-chi/cors` - CORS middleware
     - `github.com/golang-migrate/migrate/v4` - Database migrations
     - `github.com/mattn/go-sqlite3` - SQLite driver
     - `golang.org/x/crypto` - Password hashing (bcrypt)

3. **Makefile**

   - Commands untuk development workflow:
     - `make dev` - Run dengan hot reload (menggunakan Air)
     - `make build` - Build aplikasi
     - `make run` - Run aplikasi yang sudah di-build
     - `make test` - Run tests
     - `make clean` - Clean build artifacts
     - `make install` - Install dependencies
     - `make migrate-up` / `make migrate-down` - Database migrations
     - `make sqlc` - Generate sqlc code
     - `make sqlc-validate` - Validate sqlc queries

4. **Documentation**
   - `README.md` - Overview, installation, dan usage instructions
   - `.docs/ROADMAP.md` - Roadmap lengkap dengan status setiap milestone
   - `.docs/CHANGELOG.md` - Changelog untuk tracking perubahan
   - `.docs/milestones/M0-setup.md` - Dokumentasi milestone ini

## Files Created

```
vugo/
├── go.mod
├── Makefile
├── README.md
└── .docs/
    ├── ROADMAP.md
    ├── CHANGELOG.md
    └── milestones/
        └── M0-setup.md
```

## Technical Decisions

1. **Module Name**: Menggunakan `github.com/zou/vugo` sebagai module name. Bisa diubah sesuai dengan repository yang sebenarnya nanti.

2. **Dependencies**: Memilih Chi sebagai router karena ringan dan idiomatic untuk Go. SQLite dipilih untuk kemudahan development dan deployment.

3. **Makefile**: Menggunakan Makefile untuk standardize development commands. Setiap developer bisa langsung menggunakan commands yang sama.

4. **Documentation Structure**: Memisahkan dokumentasi ke folder `.docs/` agar tidak tercampur dengan kode aplikasi. Setiap milestone akan didokumentasikan di folder `milestones/`.

## Next Steps

Setelah milestone ini selesai, langkah selanjutnya adalah:

1. **Milestone 1: Database Foundation**

   - Setup database schema
   - Setup sqlc configuration
   - Buat migrations
   - Generate sqlc code

2. **Milestone 2: HTTP Server & Routing**
   - Setup Chi router
   - Implement middleware
   - Setup route structure

## Notes

- Semua file yang dibuat sudah sesuai dengan struktur yang direncanakan di roadmap
- Dependencies yang ditambahkan adalah minimal yang diperlukan untuk foundation
- Makefile commands sudah include error handling untuk tools yang belum terinstall
- Dokumentasi dibuat dalam bahasa Indonesia untuk kemudahan pemahaman

## Testing

Milestone ini tidak memerlukan testing karena hanya setup project structure. Testing akan dimulai dari milestone berikutnya setelah ada kode yang bisa ditest.

## References

- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Chi Router Documentation](https://github.com/go-chi/chi)
- [sqlc Documentation](https://docs.sqlc.dev/)
- [golang-migrate Documentation](https://github.com/golang-migrate/migrate)
