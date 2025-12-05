-- name: CreateColumn :one
INSERT INTO columns (
    board_id, name, position, color, wip_limit
) VALUES (
    ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetColumn :one
SELECT * FROM columns
WHERE id = ? LIMIT 1;

-- name: ListColumnsByBoard :many
SELECT * FROM columns
WHERE board_id = ?
ORDER BY position ASC, created_at ASC;

-- name: UpdateColumn :one
UPDATE columns
SET 
    name = ?,
    color = ?,
    wip_limit = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: UpdateColumnPosition :exec
UPDATE columns
SET 
    position = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteColumn :exec
DELETE FROM columns
WHERE id = ?;

