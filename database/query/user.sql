-- name: CreateUser :one
INSERT INTO users (
  id,
  username,
  password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: UpdateUserPassword :one
UPDATE users
  set password_updated_at = $2,
  password = $3
WHERE id = $1
RETURNING *;

-- name: UpdateUserStatus :one
UPDATE users
  set status = $2,
  last_login = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;