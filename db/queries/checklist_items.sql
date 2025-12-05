-- name: CreateChecklistItem :one
INSERT INTO checklist_items (
    task_id, content, position
) VALUES (
    ?, ?, ?
)
RETURNING *;

-- name: GetChecklistItem :one
SELECT * FROM checklist_items
WHERE id = ? LIMIT 1;

-- name: ListChecklistItemsByTask :many
SELECT * FROM checklist_items
WHERE task_id = ?
ORDER BY position ASC, created_at ASC;

-- name: UpdateChecklistItem :one
UPDATE checklist_items
SET 
    content = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: ToggleChecklistItem :one
UPDATE checklist_items
SET 
    completed = NOT completed,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: UpdateChecklistItemPosition :exec
UPDATE checklist_items
SET 
    position = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteChecklistItem :exec
DELETE FROM checklist_items
WHERE id = ?;

