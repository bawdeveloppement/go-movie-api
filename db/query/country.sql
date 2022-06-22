-- name: CreateCountry :one
INSERT INTO countries (
    code, country_name, continent_name
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetCountry :one
SELECT * FROM countries
WHERE code = $1 LIMIT 1;

-- name: ListCountries :many
SELECT * FROM countries
ORDER BY country_name
LIMIT $1
OFFSET $2;

-- name: UpdateCountry :exec
UPDATE countries SET country_name = $2, continent_name = $3 WHERE code = $1;