-- name: CreateBookCategory :one
INSERT INTO books_categories (
    book_id,
    category_id
) VALUES (
    $1,
    $2
) RETURNING *;

-- name: GetBookCategory :one
SELECT * FROM books_categories WHERE id = $1;

-- name: GetAllBookCategories :many
SELECT * FROM books_categories;

-- name: UpdateBookCategory :one
UPDATE books_categories
SET
    book_id = $2,
    category_id = $3
WHERE id = $1
RETURNING *;

-- name: DeleteBookCategory :one
DELETE FROM books_categories WHERE id = $1
RETURNING *;
