-- name: CreateRole :one
INSERT INTO roles (
    name,
    description,
    is_deleted
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetRole :one
SELECT * FROM roles WHERE id = $1;

-- name: GetAllRoles :many
SELECT * FROM roles;

-- name: UpdateRole :one
UPDATE roles
SET
    name = $2,
    description = $3,
    is_deleted = $4
WHERE id = $1
RETURNING *;

-- name: DeleteRole :one
DELETE FROM roles WHERE id = $1
RETURNING *;
