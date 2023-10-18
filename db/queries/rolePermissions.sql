-- name: CreateRolePermission :one
INSERT INTO roles_permissions (
    role_id,
    permission_id
) VALUES (
    $1,
    $2
) RETURNING id;

-- name: GetRolePermission :one
SELECT * FROM roles_permissions WHERE id = $1;

-- name: GetAllRolePermissions :many
SELECT * FROM roles_permissions;

-- name: UpdateRolePermission :one
UPDATE roles_permissions
SET
    role_id = $2,
    permission_id = $3
WHERE id = $1
RETURNING id;

-- name: DeleteRolePermission :one
DELETE FROM roles_permissions WHERE id = $1
RETURNING id;
