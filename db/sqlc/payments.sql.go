// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: payments.sql

package db

import (
	"context"
)

const createPayment = `-- name: CreatePayment :one
INSERT INTO payments (
    order_id,
    amount,
    payment_status
) VALUES (
    $1,
    $2,
    $3
) RETURNING id
`

type CreatePaymentParams struct {
	OrderID       int32             `json:"order_id"`
	Amount        int32             `json:"amount"`
	PaymentStatus PaymentStatusEnum `json:"payment_status"`
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (int32, error) {
	row := q.db.QueryRow(ctx, createPayment, arg.OrderID, arg.Amount, arg.PaymentStatus)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deletePayment = `-- name: DeletePayment :one
DELETE FROM payments WHERE id = $1
RETURNING id
`

func (q *Queries) DeletePayment(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deletePayment, id)
	err := row.Scan(&id)
	return id, err
}

const getAllPayments = `-- name: GetAllPayments :many
SELECT id, order_id, amount, payment_status, is_deleted, created_at, updated_at FROM payments
`

func (q *Queries) GetAllPayments(ctx context.Context) ([]Payment, error) {
	rows, err := q.db.Query(ctx, getAllPayments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Payment{}
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.Amount,
			&i.PaymentStatus,
			&i.IsDeleted,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPayment = `-- name: GetPayment :one
SELECT id, order_id, amount, payment_status, is_deleted, created_at, updated_at FROM payments WHERE id = $1
`

func (q *Queries) GetPayment(ctx context.Context, id int32) (Payment, error) {
	row := q.db.QueryRow(ctx, getPayment, id)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.Amount,
		&i.PaymentStatus,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePayment = `-- name: UpdatePayment :one
UPDATE payments
SET
    order_id = $2,
    amount = $3,
    payment_status = $4
WHERE id = $1
RETURNING id
`

type UpdatePaymentParams struct {
	ID            int32             `json:"id"`
	OrderID       int32             `json:"order_id"`
	Amount        int32             `json:"amount"`
	PaymentStatus PaymentStatusEnum `json:"payment_status"`
}

func (q *Queries) UpdatePayment(ctx context.Context, arg UpdatePaymentParams) (int32, error) {
	row := q.db.QueryRow(ctx, updatePayment,
		arg.ID,
		arg.OrderID,
		arg.Amount,
		arg.PaymentStatus,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}