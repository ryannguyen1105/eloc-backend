-- name: CreateRole :one
INSERT INTO roles (
  id,
  description
) VALUES (
  $1, $2
)
ON CONFLICT (id)
DO UPDATE SET id = EXCLUDED.id
RETURNING *;

-- name: GetRole :one
SELECT * FROM roles
WHERE id = $1 LIMIT 1;

-- name: ListRoles :many
SELECT * FROM roles
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateRole :one
UPDATE roles
set description = $2
WHERE id = $1
RETURNING *;

-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = $1;