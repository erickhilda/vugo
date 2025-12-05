-- name: CreateLabel :one
INSERT INTO labels (
    project_id, name, color
) VALUES (
    ?, ?, ?
)
RETURNING *;

-- name: GetLabel :one
SELECT * FROM labels
WHERE id = ? LIMIT 1;

-- name: ListLabelsByProject :many
SELECT * FROM labels
WHERE project_id = ?
ORDER BY name ASC;

-- name: UpdateLabel :one
UPDATE labels
SET 
    name = ?,
    color = ?
WHERE id = ?
RETURNING *;

-- name: DeleteLabel :exec
DELETE FROM labels
WHERE id = ?;

