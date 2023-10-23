-- name: CreateCategoryImage :one
INSERT INTO categories_images (
    category_id,
    image,
    is_deleted
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetCategoryImage :one
SELECT * FROM categories_images WHERE id = $1;

-- name: GetAllCategoryImages :many
SELECT * FROM categories_images;

-- name: UpdateCategoryImage :one
UPDATE categories_images
SET
    category_id =  CASE
    WHEN @set_category_id::boolean = TRUE THEN @category_id
    ELSE category_id
    END,
    image =  CASE
    WHEN @set_image::boolean = TRUE THEN @image
    ELSE image
    END,
    is_deleted =  CASE
    WHEN @set_is_deleted::boolean = TRUE THEN @is_deleted
    ELSE is_deleted
    END
WHERE id = @id
RETURNING *;

-- name: DeleteCategoryImage :one
DELETE FROM categories_images WHERE id = $1
RETURNING *;
