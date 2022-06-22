-- name: CreateUser :one
INSERT INTO users (
    full_name, country_code
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :exec
UPDATE users SET full_name = $2, country_code = $3 WHERE id = $1;