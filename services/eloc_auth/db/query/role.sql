-- name: CreateRole :one
INSERT INTO roles (
  id,
  name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetRole :one
SELECT * FROM roles
WHERE id = $1 LIMIT 1;

-- name: ListRoles :many
SELECT * FROM roles
ORDER BY name;

-- name: UpdateRole :exec
UPDATE roles
set name = $2
WHERE id = $1;

-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = $1;