-- name: CreateCategoryImage :one
INSERT INTO categories_images (
    category_id,
    image
) VALUES (
    $1,
    $2
) RETURNING id;

-- name: GetCategoryImage :one
SELECT * FROM categories_images WHERE id = $1;

-- name: GetAllCategoryImages :many
SELECT * FROM categories_images;

-- name: UpdateCategoryImage :one
UPDATE categories_images
SET
    category_id = $2,
    image = $3
WHERE id = $1
RETURNING id;

-- name: DeleteCategoryImage :one
DELETE FROM categories_images WHERE id = $1
RETURNING id;
