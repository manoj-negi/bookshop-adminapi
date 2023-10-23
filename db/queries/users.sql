-- name: CreateUser :one
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
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: UpdateUser :one
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
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users WHERE id = $1
RETURNING *;
