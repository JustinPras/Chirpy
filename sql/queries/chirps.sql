-- name: CreateChirp :one
INSERT INTO chirps(id, created_at, updated_at, body, user_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: GetChirpByID :one
SELECT * FROM chirps
WHERE id = $1;

-- name: GetChirpsOrderByCreatedAtAsc :many
SELECT * FROM chirps
ORDER BY created_at ASC;

-- name: GetChirpsForUserOrderByCreatedAtAsc :many
SELECT * FROM chirps
WHERE user_id = $1
ORDER BY created_at ASC;

-- name: GetChirpsOrderByCreatedAtDesc :many
SELECT * FROM chirps
ORDER BY created_at DESC;

-- name: GetChirpsForUserOrderByCreatedAtDesc :many
SELECT * FROM chirps
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: DeleteChirp :exec
DELETE FROM chirps
WHERE id = $1;