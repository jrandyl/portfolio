-- name: CreateClient :one
INSERT INTO clients (
  id,
  first_name,
  middle_name,
  last_name,
  full_name,
  email,
  message
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetClient :one
SELECT * FROM clients;

-- name: ListClients :many
SELECT * FROM clients
ORDER BY created_at DESC;

-- name: DeleteClient :exec
DELETE FROM clients
WHERE id = $1;