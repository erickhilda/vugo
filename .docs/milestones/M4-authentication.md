# Milestone 4: Authentication System

**Status**: ⏳ Pending  
**Date Started**: TBD  
**Date Completed**: TBD

## Overview

Milestone ini fokus pada implementasi sistem autentikasi lengkap termasuk user registration, login, logout, dan session management. Ini adalah foundation yang diperlukan sebelum bisa mengakses fitur-fitur aplikasi yang memerlukan autentikasi.

## Objectives

1. Implementasi user registration dengan validation
2. Implementasi login/logout flow
3. Setup session management dengan secure cookies
4. Password hashing menggunakan bcrypt
5. Auth middleware untuk protected routes
6. Login dan Register pages dengan Vue forms

## Deliverables

### User Registration

- Form registration dengan validasi
- Email uniqueness check
- Password strength requirements
- Password hashing dengan bcrypt
- Auto-login setelah registration

### Login System

- Login form dengan email/password
- Session creation setelah login berhasil
- Remember me functionality (optional)
- Error handling untuk invalid credentials

### Session Management

- Secure cookie-based sessions
- Session expiration handling
- Session cleanup untuk expired sessions
- Logout dengan session deletion

### Auth Middleware

- Middleware untuk protected routes
- Redirect ke login jika tidak authenticated
- User context injection ke request

### Pages

- Login page dengan Vue form
- Register page dengan Vue form
- Error messages handling
- Success redirects

## Files to Create

```
vugo/
├── internal/
│   ├── handlers/
│   │   └── auth.go              # Auth handlers (register, login, logout)
│   ├── services/
│   │   └── auth.go              # Auth business logic
│   └── middleware/
│       └── auth.go              # Auth middleware (update dari placeholder)
├── templates/
│   └── pages/
│       ├── login.html           # Login page dengan Vue form
│       └── register.html         # Register page dengan Vue form
└── db/queries/
    ├── users.sql                # Already exists, verify queries
    └── sessions.sql             # Already exists, verify queries
```

## Technical Implementation

### Password Hashing

```go
// Using golang.org/x/crypto/bcrypt
import "golang.org/x/crypto/bcrypt"

// Hash password
hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// Verify password
err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
```

### Session Management

- Session ID: Generate random UUID atau secure random string
- Session storage: Database (sessions table)
- Cookie settings:
  - HttpOnly: true
  - Secure: true (production)
  - SameSite: Lax
  - MaxAge: 7 days (atau configurable)

### Auth Middleware

```go
func RequireAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get session from cookie
        // Verify session exists and not expired
        // Get user from session
        // Inject user to request context
        // Call next handler
    })
}
```

### Routes Structure

```
POST   /register          - User registration
POST   /login             - User login
POST   /logout            - User logout
GET    /login             - Login page
GET    /register          - Register page
```

## Database Queries Needed

### Users (Already exists in db/queries/users.sql)

- ✅ GetUserByEmail
- ✅ CreateUser
- ✅ GetUser
- ✅ UpdateUserPassword

### Sessions (Already exists in db/queries/sessions.sql)

- ✅ CreateSession
- ✅ GetSession
- ✅ DeleteSession
- ✅ DeleteExpiredSessions
- ✅ DeleteUserSessions

## Frontend Implementation

### Login Form (Vue)

```html
<form @submit.prevent="handleLogin">
  <input v-model="email" type="email" required />
  <input v-model="password" type="password" required />
  <button type="submit">Login</button>
  <div v-if="error" class="error">{{ error }}</div>
</form>
```

### Register Form (Vue)

```html
<form @submit.prevent="handleRegister">
  <input v-model="name" type="text" required />
  <input v-model="email" type="email" required />
  <input v-model="password" type="password" required />
  <input v-model="confirmPassword" type="password" required />
  <button type="submit">Register</button>
  <div v-if="error" class="error">{{ error }}</div>
</form>
```

## Validation Rules

### Registration

- Name: Required, min 2 characters, max 100 characters
- Email: Required, valid email format, unique
- Password: Required, min 8 characters, at least one letter and one number
- Confirm Password: Must match password

### Login

- Email: Required, valid email format
- Password: Required

## Security Considerations

1. **Password Security**

   - Never store plain text passwords
   - Use bcrypt with appropriate cost factor
   - Validate password strength

2. **Session Security**

   - Use secure, random session IDs
   - Set appropriate cookie flags (HttpOnly, Secure, SameSite)
   - Implement session expiration
   - Clean up expired sessions regularly

3. **CSRF Protection**

   - Consider adding CSRF tokens for forms (future enhancement)

4. **Rate Limiting**
   - Consider rate limiting for login attempts (future enhancement)

## Testing Checklist

- [ ] User can register with valid data
- [ ] Registration fails with duplicate email
- [ ] Registration fails with invalid data
- [ ] User can login with correct credentials
- [ ] Login fails with incorrect credentials
- [ ] Session is created after login
- [ ] Session persists across requests
- [ ] User can logout
- [ ] Session is deleted after logout
- [ ] Protected routes redirect to login if not authenticated
- [ ] Authenticated users can access protected routes
- [ ] Expired sessions are cleaned up

## Dependencies

- `golang.org/x/crypto/bcrypt` - Password hashing (already in go.mod)
- Session management - Custom implementation using database

## Next Steps

Setelah milestone ini selesai, langkah selanjutnya adalah:

1. **Milestone 5: Projects CRUD** - Implement project management dengan auth protection
2. Update routes untuk menggunakan auth middleware
3. Add user context ke semua protected handlers

## Notes

- Session management menggunakan database untuk simplicity
- Future enhancement: bisa migrate ke Redis untuk better performance
- Password requirements bisa dibuat configurable
- Remember me functionality bisa ditambahkan sebagai optional feature

## References

- [OWASP Password Storage Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)
- [Go bcrypt Documentation](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [Chi Middleware Patterns](https://github.com/go-chi/chi#middlewares)
