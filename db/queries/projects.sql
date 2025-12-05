-- name: CreateProject :one
INSERT INTO projects (
    owner_id, name, description, color
) VALUES (
    ?, ?, ?, ?
)
RETURNING *;

-- name: GetProject :one
SELECT * FROM projects
WHERE id = ? LIMIT 1;

-- name: GetProjectWithOwner :one
SELECT 
    p.*,
    u.id as owner_id,
    u.name as owner_name,
    u.email as owner_email
FROM projects p
JOIN users u ON p.owner_id = u.id
WHERE p.id = ? LIMIT 1;

-- name: ListProjectsByOwner :many
SELECT * FROM projects
WHERE owner_id = ? AND archived = FALSE
ORDER BY created_at DESC;

-- name: ListProjectsByMember :many
SELECT p.* FROM projects p
JOIN project_members pm ON p.id = pm.project_id
WHERE pm.user_id = ? AND p.archived = FALSE
ORDER BY p.created_at DESC;

-- name: UpdateProject :one
UPDATE projects
SET 
    name = ?,
    description = ?,
    color = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ? AND owner_id = ?
RETURNING *;

-- name: ArchiveProject :exec
UPDATE projects
SET 
    archived = TRUE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ? AND owner_id = ?;

-- name: UnarchiveProject :exec
UPDATE projects
SET 
    archived = FALSE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ? AND owner_id = ?;

-- name: DeleteProject :exec
DELETE FROM projects
WHERE id = ? AND owner_id = ?;

