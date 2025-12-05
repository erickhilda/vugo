-- name: AddTaskLabel :one
INSERT INTO task_labels (
    task_id, label_id
) VALUES (
    ?, ?
)
RETURNING *;

-- name: GetTaskLabels :many
SELECT 
    tl.*,
    l.id as label_id,
    l.name as label_name,
    l.color as label_color
FROM task_labels tl
JOIN labels l ON tl.label_id = l.id
WHERE tl.task_id = ?
ORDER BY l.name ASC;

-- name: RemoveTaskLabel :exec
DELETE FROM task_labels
WHERE task_id = ? AND label_id = ?;

-- name: HasTaskLabel :one
SELECT EXISTS(
    SELECT 1 FROM task_labels
    WHERE task_id = ? AND label_id = ?
) as has_label;

