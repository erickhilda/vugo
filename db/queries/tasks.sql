-- name: CreateTask :one
INSERT INTO tasks (
    column_id, created_by, title, description, position, priority, due_date
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks
WHERE id = ? LIMIT 1;

-- name: GetTaskWithCreator :one
SELECT 
    t.*,
    u.id as creator_id,
    u.name as creator_name,
    u.email as creator_email,
    u.avatar_url as creator_avatar_url
FROM tasks t
JOIN users u ON t.created_by = u.id
WHERE t.id = ? LIMIT 1;

-- name: ListTasksByColumn :many
SELECT * FROM tasks
WHERE column_id = ? AND completed_at IS NULL
ORDER BY position ASC, created_at ASC;

-- name: ListTasksByUser :many
SELECT t.* FROM tasks t
JOIN task_assignees ta ON t.id = ta.task_id
WHERE ta.user_id = ? AND t.completed_at IS NULL
ORDER BY t.due_date ASC, t.created_at DESC;

-- name: ListTasksByProject :many
SELECT t.* FROM tasks t
JOIN columns c ON t.column_id = c.id
JOIN boards b ON c.board_id = b.id
WHERE b.project_id = ? AND t.completed_at IS NULL
ORDER BY t.due_date ASC, t.created_at DESC;

-- name: UpdateTask :one
UPDATE tasks
SET 
    title = ?,
    description = ?,
    priority = ?,
    due_date = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: MoveTask :one
UPDATE tasks
SET 
    column_id = ?,
    position = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: UpdateTaskPosition :exec
UPDATE tasks
SET 
    position = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: CompleteTask :exec
UPDATE tasks
SET 
    completed_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UncompleteTask :exec
UPDATE tasks
SET 
    completed_at = NULL,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?;

