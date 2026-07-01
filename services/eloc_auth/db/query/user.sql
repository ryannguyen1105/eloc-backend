-- name: CreateUser :one
INSERT INTO users (
  email, password_hash, fullname, role_id, is_active, is_verified
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT id, email, password_hash, fullname, role_id, is_active, is_verified, created_at, updated_at
FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT id, email, fullname, role_id, is_active, is_verified, created_at, updated_at 
FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUserDetail :one
UPDATE users
SET 
    fullname = $2,
    role_id = $3,
    updated_at = now()
WHERE id = $1
RETURNING id, email, fullname, role_id, is_active, is_verified, created_at, updated_at;

-- name: UpdateUserStatus :exec
UPDATE users
SET 
    is_active = $2,
    is_verified = $3,
    updated_at = now()
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE email = $1;