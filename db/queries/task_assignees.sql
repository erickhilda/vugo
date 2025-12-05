-- name: AssignTaskToUser :one
INSERT INTO task_assignees (
    task_id, user_id
) VALUES (
    ?, ?
)
RETURNING *;

-- name: GetTaskAssignees :many
SELECT 
    ta.*,
    u.id as user_id,
    u.name as user_name,
    u.email as user_email,
    u.avatar_url as user_avatar_url
FROM task_assignees ta
JOIN users u ON ta.user_id = u.id
WHERE ta.task_id = ?
ORDER BY ta.assigned_at ASC;

-- name: UnassignTaskFromUser :exec
DELETE FROM task_assignees
WHERE task_id = ? AND user_id = ?;

-- name: IsTaskAssignedToUser :one
SELECT EXISTS(
    SELECT 1 FROM task_assignees
    WHERE task_id = ? AND user_id = ?
) as is_assigned;

