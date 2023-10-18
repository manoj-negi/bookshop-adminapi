// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: authors.sql

package db

import (
	"context"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (
    name
) VALUES (
    $1
) RETURNING id
`

func (q *Queries) CreateAuthor(ctx context.Context, name string) (int32, error) {
	row := q.db.QueryRow(ctx, createAuthor, name)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteAuthor = `-- name: DeleteAuthor :one
DELETE FROM authors WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteAuthor, id)
	err := row.Scan(&id)
	return id, err
}

const getAllAuthors = `-- name: GetAllAuthors :many
SELECT id, name, is_deleted, created_at, updated_at FROM authors
`

func (q *Queries) GetAllAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.Query(ctx, getAllAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Author{}
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
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

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, is_deleted, created_at, updated_at FROM authors WHERE id = $1
`

func (q *Queries) GetAuthor(ctx context.Context, id int32) (Author, error) {
	row := q.db.QueryRow(ctx, getAuthor, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateAuthor = `-- name: UpdateAuthor :one
UPDATE authors
SET
    name = $2
WHERE id = $1
RETURNING id
`

type UpdateAuthorParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateAuthor, arg.ID, arg.Name)
	var id int32
	err := row.Scan(&id)
	return id, err
}