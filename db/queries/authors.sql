-- name: CreateAuthor :one
INSERT INTO authors (
    name
) VALUES (
    $1
) RETURNING id;

-- name: GetAuthor :one
SELECT * FROM authors WHERE id = $1;

-- name: GetAllAuthors :many
SELECT * FROM authors;

-- name: UpdateAuthor :one
UPDATE authors
SET
    name = $2
WHERE id = $1
RETURNING id;

-- name: DeleteAuthor :one
DELETE FROM authors WHERE id = $1
RETURNING id;
