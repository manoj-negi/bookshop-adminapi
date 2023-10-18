-- name: CreateRole :one
INSERT INTO roles (
    name,
    description
) VALUES (
    $1,
    $2
) RETURNING id;

-- name: GetRole :one
SELECT * FROM roles WHERE id = $1;

-- name: GetAllRoles :many
SELECT * FROM roles;

-- name: UpdateRole :one
UPDATE roles
SET
    name = $2,
    description = $3
WHERE id = $1
RETURNING id;

-- name: DeleteRole :one
DELETE FROM roles WHERE id = $1
RETURNING id;
