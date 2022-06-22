-- name: CreateMovie :one
INSERT INTO movies (
    title, director_id, admin_id, release_year, production_country_code
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;


-- name: GetMovieById :one
SELECT * FROM movies
WHERE id = $1 LIMIT 1;

-- name: GetMovieByTitle :one
SELECT * FROM movies
WHERE title = $1 LIMIT 1;

-- name: ListMoviesById :many
SELECT * FROM movies
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListMoviesByTitle :many
SELECT * FROM movies
ORDER BY title
LIMIT $1
OFFSET $2;

-- name: UpdateMovie :exec
UPDATE movies SET title = $2, director_id = $3, release_year = $4, production_country_code = $5 WHERE id = $1;