-- name: CreateOrder :one
INSERT INTO orders (
    book_id,
    user_id,
    order_no,
    quantity,
    total_price,
    status
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING id;

-- name: GetOrder :one
SELECT * FROM orders WHERE id = $1;

-- name: GetAllOrders :many
SELECT * FROM orders;

-- name: UpdateOrder :one
UPDATE orders
SET
    book_id = $2,
    user_id = $3,
    order_no = $4,
    quantity = $5,
    total_price = $6,
    status = $7
WHERE id = $1
RETURNING id;

-- name: DeleteOrder :one
DELETE FROM orders WHERE id = $1
RETURNING id;
