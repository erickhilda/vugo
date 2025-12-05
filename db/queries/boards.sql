-- name: CreateBoard :one
INSERT INTO boards (
    project_id, name, position
) VALUES (
    ?, ?, ?
)
RETURNING *;

-- name: GetBoard :one
SELECT * FROM boards
WHERE id = ? LIMIT 1;

-- name: ListBoardsByProject :many
SELECT * FROM boards
WHERE project_id = ?
ORDER BY position ASC, created_at ASC;

-- name: UpdateBoard :one
UPDATE boards
SET 
    name = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: UpdateBoardPosition :exec
UPDATE boards
SET 
    position = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteBoard :exec
DELETE FROM boards
WHERE id = ?;

