-- name: CreateCategory :one
INSERT INTO categories (
  name , slug
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetCategoryByID :one
SELECT * FROM categories
WHERE id = $1 LIMIT 1;

-- name: GetCategoryBySlug :one
SELECT * FROM categories
WHERE slug = $1 LIMIT 1;

-- name: ListCategories :many
SELECT * FROM categories
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: UpdateCategory :exec
UPDATE categories
  set name = $2,
  slug = $3
WHERE id = $1;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;