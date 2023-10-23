-- name: CreateCountry :one
INSERT INTO countries (
    iso2,
    short_name,
    long_name,
    numcode,
    calling_code,
    cctld,
    is_deleted
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
) RETURNING *;

-- name: GetCountry :one
SELECT * FROM countries WHERE id = $1;

-- name: GetAllCountries :many
SELECT * FROM countries;

-- name: UpdateCountry :one
UPDATE countries
SET
    iso2 = $2,
    short_name = $3,
    long_name = $4,
    numcode = $5,
    calling_code = $6,
    cctld = $7,
    is_deleted = $8
WHERE id = $1
RETURNING *;

-- name: DeleteCountry :one
DELETE FROM countries WHERE id = $1
RETURNING *;
