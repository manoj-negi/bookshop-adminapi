// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    first_name,
    last_name,
    gender,
    dob,
    address,
    city,
    state,
    country_id,
    mobile_no,
    username,
    email,
    password,
    role_id
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13
) RETURNING id
`

type CreateUserParams struct {
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Gender    GenderEnum  `json:"gender"`
	Dob       pgtype.Date `json:"dob"`
	Address   string      `json:"address"`
	City      string      `json:"city"`
	State     string      `json:"state"`
	CountryID int32       `json:"country_id"`
	MobileNo  string      `json:"mobile_no"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	RoleID    int32       `json:"role_id"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int32, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Gender,
		arg.Dob,
		arg.Address,
		arg.City,
		arg.State,
		arg.CountryID,
		arg.MobileNo,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.RoleID,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM users WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteUser, id)
	err := row.Scan(&id)
	return id, err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, first_name, last_name, gender, dob, address, city, state, country_id, mobile_no, username, email, password, role_id, is_deleted, created_at, updated_at FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Gender,
			&i.Dob,
			&i.Address,
			&i.City,
			&i.State,
			&i.CountryID,
			&i.MobileNo,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.RoleID,
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

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, gender, dob, address, city, state, country_id, mobile_no, username, email, password, role_id, is_deleted, created_at, updated_at FROM users WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Gender,
		&i.Dob,
		&i.Address,
		&i.City,
		&i.State,
		&i.CountryID,
		&i.MobileNo,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.RoleID,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
    first_name = $2,
    last_name = $3,
    gender = $4,
    dob = $5,
    address = $6,
    city = $7,
    state = $8,
    country_id = $9,
    mobile_no = $10,
    username = $11,
    email = $12,
    password = $13,
    role_id = $14
WHERE id = $1
RETURNING id
`

type UpdateUserParams struct {
	ID        int32       `json:"id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Gender    GenderEnum  `json:"gender"`
	Dob       pgtype.Date `json:"dob"`
	Address   string      `json:"address"`
	City      string      `json:"city"`
	State     string      `json:"state"`
	CountryID int32       `json:"country_id"`
	MobileNo  string      `json:"mobile_no"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	RoleID    int32       `json:"role_id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Gender,
		arg.Dob,
		arg.Address,
		arg.City,
		arg.State,
		arg.CountryID,
		arg.MobileNo,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.RoleID,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}