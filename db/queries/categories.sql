-- name: CreateCategory :one
INSERT INTO categories (
    name,
    is_special
) VALUES (
    $1,
    $2
) RETURNING id;

-- name: GetCategory :one
SELECT * FROM categories WHERE id = $1;

-- name: GetAllCategories :many
SELECT * FROM categories;

-- name: UpdateCategory :one
UPDATE categories
SET
    name = $2,
    is_special = $3
WHERE id = $1
RETURNING id;

-- name: DeleteCategory :one
DELETE FROM categories WHERE id = $1
RETURNING id;
