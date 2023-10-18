-- name: CreateBook :one
INSERT INTO books (
    title,
    author_id,
    publication_date,
    price,
    stock_quantity
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING id;

-- name: GetBook :one
SELECT * FROM books WHERE id = $1;

-- name: GetAllBooks :many
SELECT * FROM books;

-- name: UpdateBook :one
UPDATE books
SET
    title = $2,
    author_id = $3,
    publication_date = $4,
    price = $5,
    stock_quantity = $6
WHERE id = $1
RETURNING id;

-- name: DeleteBook :one
DELETE FROM books WHERE id = $1
RETURNING id;
