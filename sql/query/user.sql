-- name: CreateUser :one
INSERT INTO users (
    id,
    user_name,
    password_hash
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE user_name = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at
LIMIT $1
OFFSET $2;
