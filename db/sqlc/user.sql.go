// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    full_name, country_code
) VALUES (
    $1, $2
) RETURNING id, full_name, created_at, country_code
`

type CreateUserParams struct {
	FullName    string        `json:"fullName"`
	CountryCode sql.NullInt32 `json:"countryCode"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, arg.FullName, arg.CountryCode)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.CreatedAt,
		&i.CountryCode,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, full_name, created_at, country_code FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.CreatedAt,
		&i.CountryCode,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, full_name, created_at, country_code FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.CreatedAt,
			&i.CountryCode,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET full_name = $2, country_code = $3 WHERE id = $1
`

type UpdateUserParams struct {
	ID          int32         `json:"id"`
	FullName    string        `json:"fullName"`
	CountryCode sql.NullInt32 `json:"countryCode"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.exec(ctx, q.updateUserStmt, updateUser, arg.ID, arg.FullName, arg.CountryCode)
	return err
}
