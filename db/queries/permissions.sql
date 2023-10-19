-- name: CreatePermission :one
INSERT INTO permissions (
    name,
    permission,
    is_deleted
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetPermission :one
SELECT * FROM permissions WHERE id = $1;

-- name: GetAllPermissions :many
SELECT * FROM permissions;

-- name: UpdatePermission :one
UPDATE permissions
SET
    name = $2,
    permission = $3,
    is_deleted = $4
WHERE id = $1
RETURNING *;

-- name: DeletePermission :one
DELETE FROM permissions WHERE id = $1
RETURNING *;
