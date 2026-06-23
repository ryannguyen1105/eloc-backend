-- name: CreateRole :one
INSERT INTO roles (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: GetRole :one
SELECT * FROM roles
WHERE id = $1 LIMIT 1;

-- name: ListRoles :many
SELECT * FROM roles
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: UpdateRole :one
UPDATE roles
set name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = $1;