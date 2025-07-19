-- name: GetPocketByID :one
SELECT * FROM pockets WHERE id = $1;

-- name: GetPocketByUserID :many
SELECT * FROM pockets where user_id = $1;