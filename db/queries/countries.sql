-- name: CreateCountry :one
INSERT INTO countries (
    iso2,
    short_name,
    long_name,
    numcode,
    calling_code,
    cctld
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING id;

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
    cctld = $7
WHERE id = $1
RETURNING id;

-- name: DeleteCountry :one
DELETE FROM countries WHERE id = $1
RETURNING id;
