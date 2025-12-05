-- name: CreateActivity :one
INSERT INTO activities (
    project_id, user_id, task_id, action, details
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetActivity :one
SELECT 
    a.*,
    u.id as user_id,
    u.name as user_name,
    u.email as user_email,
    u.avatar_url as user_avatar_url
FROM activities a
JOIN users u ON a.user_id = u.id
WHERE a.id = ? LIMIT 1;

-- name: ListActivitiesByProject :many
SELECT 
    a.*,
    u.id as user_id,
    u.name as user_name,
    u.email as user_email,
    u.avatar_url as user_avatar_url
FROM activities a
JOIN users u ON a.user_id = u.id
WHERE a.project_id = ?
ORDER BY a.created_at DESC
LIMIT ? OFFSET ?;

-- name: ListActivitiesByTask :many
SELECT 
    a.*,
    u.id as user_id,
    u.name as user_name,
    u.email as user_email,
    u.avatar_url as user_avatar_url
FROM activities a
JOIN users u ON a.user_id = u.id
WHERE a.task_id = ?
ORDER BY a.created_at DESC;

