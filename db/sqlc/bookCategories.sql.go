// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: bookCategories.sql

package db

import (
	"context"
)

const createBookCategory = `-- name: CreateBookCategory :one
INSERT INTO books_categories (
    book_id,
    category_id
) VALUES (
    $1,
    $2
) RETURNING id
`

type CreateBookCategoryParams struct {
	BookID     int32 `json:"book_id"`
	CategoryID int32 `json:"category_id"`
}

func (q *Queries) CreateBookCategory(ctx context.Context, arg CreateBookCategoryParams) (int32, error) {
	row := q.db.QueryRow(ctx, createBookCategory, arg.BookID, arg.CategoryID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteBookCategory = `-- name: DeleteBookCategory :one
DELETE FROM books_categories WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteBookCategory(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteBookCategory, id)
	err := row.Scan(&id)
	return id, err
}

const getAllBookCategories = `-- name: GetAllBookCategories :many
SELECT id, book_id, category_id, is_deleted, created_at, updated_at FROM books_categories
`

func (q *Queries) GetAllBookCategories(ctx context.Context) ([]BooksCategory, error) {
	rows, err := q.db.Query(ctx, getAllBookCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BooksCategory{}
	for rows.Next() {
		var i BooksCategory
		if err := rows.Scan(
			&i.ID,
			&i.BookID,
			&i.CategoryID,
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

const getBookCategory = `-- name: GetBookCategory :one
SELECT id, book_id, category_id, is_deleted, created_at, updated_at FROM books_categories WHERE id = $1
`

func (q *Queries) GetBookCategory(ctx context.Context, id int32) (BooksCategory, error) {
	row := q.db.QueryRow(ctx, getBookCategory, id)
	var i BooksCategory
	err := row.Scan(
		&i.ID,
		&i.BookID,
		&i.CategoryID,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateBookCategory = `-- name: UpdateBookCategory :one
UPDATE books_categories
SET
    book_id = $2,
    category_id = $3
WHERE id = $1
RETURNING id
`

type UpdateBookCategoryParams struct {
	ID         int32 `json:"id"`
	BookID     int32 `json:"book_id"`
	CategoryID int32 `json:"category_id"`
}

func (q *Queries) UpdateBookCategory(ctx context.Context, arg UpdateBookCategoryParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateBookCategory, arg.ID, arg.BookID, arg.CategoryID)
	var id int32
	err := row.Scan(&id)
	return id, err
}