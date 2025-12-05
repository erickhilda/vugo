-- name: CreateComment :one
INSERT INTO comments (
    task_id, user_id, content
) VALUES (
    ?, ?, ?
)
RETURNING *;

-- name: GetComment :one
SELECT * FROM comments
WHERE id = ? LIMIT 1;

-- name: GetCommentWithUser :one
SELECT 
    c.*,
    u.id as user_id,
    u.name as user_name,
    u.email as user_email,
    u.avatar_url as user_avatar_url
FROM comments c
JOIN users u ON c.user_id = u.id
WHERE c.id = ? LIMIT 1;

-- name: ListCommentsByTask :many
SELECT 
    c.*,
    u.id as user_id,
    u.name as user_name,
    u.email as user_email,
    u.avatar_url as user_avatar_url
FROM comments c
JOIN users u ON c.user_id = u.id
WHERE c.task_id = ?
ORDER BY c.created_at ASC;

-- name: UpdateComment :one
UPDATE comments
SET 
    content = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ? AND user_id = ?
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = ? AND user_id = ?;

