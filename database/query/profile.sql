-- name: CreateProfile :one
INSERT INTO profiles (
  id,
  user_id,
  first_name,
  middle_name,
  last_name,
  full_name,
  email,
  contact_number
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetProfile :one
SELECT * FROM profiles
WHERE id = $1 LIMIT 1;

-- name: GetProfileUserID :one
SELECT * FROM profiles
WHERE user_id = $1 LIMIT 1;

-- name: ListProfiles :many
SELECT * FROM profiles
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: UpdateProfile :one
UPDATE profiles
  set updated_at = $2,
  first_name = $3,
  middle_name = $4,
  last_name = $5,
  full_name = $6,
  email = $7,
  contact_number = $8
WHERE id = $1
RETURNING *;

-- name: UpdateProfileUserID :one
UPDATE profiles
  set updated_at = $2,
  first_name = $3,
  middle_name = $4,
  last_name = $5,
  full_name = $6,
  email = $7,
  contact_number = $8
WHERE user_id = $1
RETURNING *;

-- name: DeleteProfile :exec
DELETE FROM profiles
WHERE id = $1;