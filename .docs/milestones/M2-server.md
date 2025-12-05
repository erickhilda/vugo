# Milestone 2: HTTP Server & Routing

**Status**: ✅ Completed  
**Date Started**: 2025-01-27  
**Date Completed**: 2025-01-27

## Overview

Milestone ini fokus pada setup HTTP server dengan Chi router, middleware, dan route structure yang akan digunakan untuk semua fitur aplikasi.

## Objectives

1. Setup Chi router sebagai HTTP router
2. Implement middleware untuk logging, recovery, dan CORS
3. Setup route grouping untuk pages dan API
4. Setup static file serving
5. Buat health check endpoint

## Deliverables

### ✅ Completed

1. **HTTP Server**

   - `cmd/server/main.go` - Entry point aplikasi
   - Support untuk environment variables (DATABASE_PATH, PORT)
   - Default values untuk development

2. **Server Setup**

   - `internal/server/server.go` - Server struct dan initialization
   - Chi router setup
   - Middleware configuration
   - Route setup

3. **Middleware**

   - Request ID untuk tracing
   - Logger untuk request logging
   - Recoverer untuk panic recovery
   - CORS handler dengan konfigurasi untuk development

4. **Routes**

   - Health check endpoint (`/health`)
   - Static file serving (`/static/*`)
   - API routes group (`/api`) - siap untuk handlers
   - Page routes group (`/`) - siap untuk page handlers

5. **Auth Middleware Placeholder**
   - `internal/middleware/auth.go` - Placeholder untuk auth middleware
   - Akan diimplementasikan di Milestone 4

## Files Created

```
vugo/
├── cmd/
│   └── server/
│       └── main.go
└── internal/
    ├── server/
    │   ├── server.go
    │   └── routes.go
    └── middleware/
        └── auth.go
```

## Technical Decisions

1. **Chi Router**: Dipilih karena ringan, idiomatic untuk Go, dan memiliki middleware yang bagus.

2. **Middleware Order**:

   - RequestID → Logger → Recoverer → CORS
   - RequestID harus pertama untuk tracing
   - Recoverer harus sebelum handler untuk catch panics

3. **CORS Configuration**:

   - Allowed origins untuk localhost (development)
   - Support untuk credentials (cookies)
   - Max age 300 seconds

4. **Route Structure**:

   - `/health` - Health check
   - `/static/*` - Static files
   - `/api/*` - API endpoints
   - `/*` - Page routes

5. **Environment Variables**:
   - `DATABASE_PATH` - Path ke database file (default: `data/vugo.db`)
   - `PORT` - Server port (default: `8080`)

## Next Steps

Setelah milestone ini selesai, langkah selanjutnya adalah:

1. **Milestone 3: Template Engine & Base Layout** - Setup templates dengan Vue integration
2. **Milestone 4: Authentication** - Implement auth handlers dan middleware

## Testing

Untuk test server:

```bash
# Build and run
make build
make run

# Or run directly
go run cmd/server/main.go

# Test health check
curl http://localhost:8080/health
```

Expected response:

```json
{ "status": "ok" }
```

## Notes

- Server sudah siap untuk menambahkan handlers di milestone berikutnya
- Middleware sudah dikonfigurasi dengan baik untuk development
- Route structure sudah disiapkan untuk scalability
- Auth middleware placeholder sudah dibuat untuk implementasi nanti

## References

- [Chi Router Documentation](https://github.com/go-chi/chi)
- [Chi Middleware](https://github.com/go-chi/chi/tree/master/middleware)
