-- name: CreatePayment :one
INSERT INTO payments (
    order_id,
    amount,
    payment_status
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetPayment :one
SELECT * FROM payments WHERE id = $1;

-- name: GetAllPayments :many
SELECT * FROM payments;

-- name: UpdatePayment :one
UPDATE payments
SET
    order_id = $2,
    amount = $3,
    payment_status = $4
WHERE id = $1
RETURNING *;

-- name: DeletePayment :one
DELETE FROM payments WHERE id = $1
RETURNING *;
