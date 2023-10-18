-- name: CreateBanner :one
INSERT INTO banners (
    name,
    image,
    start_date,
    end_date,
    offer_id
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING id;

-- name: GetBanner :one
SELECT * FROM banners WHERE id = $1;

-- name: GetAllBanners :many
SELECT * FROM banners;

-- name: UpdateBanner :one
UPDATE banners
SET
    name = $2,
    image = $3,
    start_date = $4,
    end_date = $5,
    offer_id = $6
WHERE id = $1
RETURNING id;

-- name: DeleteBanner :one
DELETE FROM banners WHERE id = $1
RETURNING id;
