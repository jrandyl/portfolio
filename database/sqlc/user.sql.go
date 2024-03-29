// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user.sql

package sqlc

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  id,
  username,
  password
) VALUES (
  $1, $2, $3
)
RETURNING id, username, password, password_updated_at, last_login, status, created_at
`

type CreateUserParams struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password []byte `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.ID, arg.Username, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.PasswordUpdatedAt,
		&i.LastLogin,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username, password, password_updated_at, last_login, status, created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.PasswordUpdatedAt,
		&i.LastLogin,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getUserUsername = `-- name: GetUserUsername :one
SELECT id, username, password, password_updated_at, last_login, status, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.PasswordUpdatedAt,
		&i.LastLogin,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, password, password_updated_at, last_login, status, created_at FROM users
ORDER BY created_at DESC
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Password,
			&i.PasswordUpdatedAt,
			&i.LastLogin,
			&i.Status,
			&i.CreatedAt,
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

const updateUserPassword = `-- name: UpdateUserPassword :one
UPDATE users
  set password_updated_at = $2,
  password = $3
WHERE id = $1
RETURNING id, username, password, password_updated_at, last_login, status, created_at
`

type UpdateUserPasswordParams struct {
	ID                string `json:"id"`
	PasswordUpdatedAt string `json:"password_updated_at"`
	Password          []byte `json:"password"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserPassword, arg.ID, arg.PasswordUpdatedAt, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.PasswordUpdatedAt,
		&i.LastLogin,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserStatus = `-- name: UpdateUserStatus :one
UPDATE users
  set status = $2,
  last_login = $3
WHERE id = $1
RETURNING id, username, password, password_updated_at, last_login, status, created_at
`

type UpdateUserStatusParams struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	LastLogin string `json:"last_login"`
}

func (q *Queries) UpdateUserStatus(ctx context.Context, arg UpdateUserStatusParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserStatus, arg.ID, arg.Status, arg.LastLogin)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.PasswordUpdatedAt,
		&i.LastLogin,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}
