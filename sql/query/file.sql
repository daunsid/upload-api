-- name: CreateFile :one
INSERT INTO files (
    user_id,
    file_id,
    file_name
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetFile :one
SELECT * FROM files
WHERE file_id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM files
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;


