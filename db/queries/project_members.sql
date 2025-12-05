-- name: AddProjectMember :one
INSERT INTO project_members (
    project_id, user_id, role
) VALUES (
    ?, ?, ?
)
RETURNING *;

-- name: GetProjectMember :one
SELECT * FROM project_members
WHERE project_id = ? AND user_id = ? LIMIT 1;

-- name: ListProjectMembers :many
SELECT 
    pm.*,
    u.id as user_id,
    u.name as user_name,
    u.email as user_email,
    u.avatar_url as user_avatar_url
FROM project_members pm
JOIN users u ON pm.user_id = u.id
WHERE pm.project_id = ?
ORDER BY pm.joined_at ASC;

-- name: UpdateProjectMemberRole :exec
UPDATE project_members
SET role = ?
WHERE project_id = ? AND user_id = ?;

-- name: RemoveProjectMember :exec
DELETE FROM project_members
WHERE project_id = ? AND user_id = ?;

-- name: IsProjectMember :one
SELECT EXISTS(
    SELECT 1 FROM project_members
    WHERE project_id = ? AND user_id = ?
) as is_member;

