-- name: CreatePermission :one
INSERT INTO permissions (
    name,
    permission
) VALUES (
    $1,
    $2
) RETURNING id;

-- name: GetPermission :one
SELECT * FROM permissions WHERE id = $1;

-- name: GetAllPermissions :many
SELECT * FROM permissions;

-- name: UpdatePermission :one
UPDATE permissions
SET
    name = $2,
    permission = $3
WHERE id = $1
RETURNING id;

-- name: DeletePermission :one
DELETE FROM permissions WHERE id = $1
RETURNING id;
