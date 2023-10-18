-- name: CreateOffer :one
INSERT INTO offers (
    book_id,
    discount_percentage,
    start_date,
    end_date
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING id;

-- name: GetOffer :one
SELECT * FROM offers WHERE id = $1;

-- name: GetAllOffers :many
SELECT * FROM offers;

-- name: UpdateOffer :one
UPDATE offers
SET
    book_id = $2,
    discount_percentage = $3,
    start_date = $4,
    end_date = $5
WHERE id = $1
RETURNING id;

-- name: DeleteOffer :one
DELETE FROM offers WHERE id = $1
RETURNING id;
