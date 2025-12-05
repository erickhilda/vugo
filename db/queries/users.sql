-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    email, password_hash, name, avatar_url
) VALUES (
    ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET 
    name = ?,
    avatar_url = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE users
SET 
    password_hash = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

