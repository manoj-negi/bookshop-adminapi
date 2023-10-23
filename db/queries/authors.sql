-- name: CreateAuthor :one
INSERT INTO authors (
    name,
    is_deleted
) VALUES (
    $1,
    $2
) RETURNING *;

-- name: GetAuthor :one
SELECT * FROM authors WHERE id = $1;

-- name: GetAllAuthors :many
SELECT * FROM authors;

-- name: UpdateAuthor :one
UPDATE authors
SET
    name = $2,
    is_deleted = $3
WHERE id = $1
RETURNING *;

-- name: DeleteAuthor :one
DELETE FROM authors WHERE id = $1
RETURNING *;
