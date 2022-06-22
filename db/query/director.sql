-- name: CreateDirector :one
INSERT INTO directors (
    firstname, lastname
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetDirector :one
SELECT * FROM directors
WHERE id = $1 LIMIT 1;

-- name: ListDirectors :many
SELECT * FROM directors
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateDirector :exec
UPDATE directors SET firstname = $2, lastname = $3 WHERE id = $1;